package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/playground/httphash/common"
)

type HTTPHash struct {
	URLs       []string
	Parallel   int
	httpClient http.Client
}

func New(parallel int, urls []string) (*HTTPHash, error) {
	if len(urls) == 0 {
		return nil, EmptyURL{}
	}

	if parallel <= 0 {
		return nil, InvalidFlag{}
	}

	return &HTTPHash{
		Parallel: parallel,
		URLs:     urls,
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil

}

// Makes parallel requests and returns the number of requests made.
func (h *HTTPHash) Process() int {
	buffer := make(chan string, h.Parallel)
	response := make(chan string)
	var slc []string

	for _, url := range h.URLs {
		go h.DoRequest(url, buffer, response)
	}

	for {
		select {
		case rsp := <-response:
			<-buffer
			slc = append(slc, rsp)
			fmt.Println(rsp)
			if len(slc) == len(h.URLs) {
				close(buffer)
				close(response)
				return len(slc)
			}
		}
	}

}

// Runs assynchronously. Must be called 'go' keyword
// Executes if the buffer is not full. Buffer size
// depends on the number of parallel requests.
//
// If the buffer is not full, the method makes
// an http request and writes the answer in the response channel.
func (h *HTTPHash) DoRequest(url string, buffer chan string, response chan string) {
	buffer <- url

	url = common.ResolveURL(url)

	hash, err := h.request(url)
	if err != nil {
		response <- fmt.Sprintf("%s %s", url, "error")
		return
	}

	response <- fmt.Sprintf("%s %s", url, hash)
	return

}

// Performs http GET request for a given URL.
func (h *HTTPHash) request(url string) (string, error) {

	buffer := new(bytes.Buffer)

	req, err := http.NewRequest(http.MethodGet, url, buffer)

	resp, err := h.httpClient.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp == nil {
		return "", nil
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodystring := string(bodyBytes)
	hash := common.MD5Hash(bodystring)

	return hash, nil
}
