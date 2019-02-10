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
	urls       []string
	parallel   int
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
		parallel: parallel,
		urls:     urls,
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil

}

func (h *HTTPHash) Process() {
	buffer := make(chan string, h.parallel)
	response := make(chan string)

	var slc []string

	for _, url := range h.urls {
		go h.doRequest(url, buffer, response)
	}

	for {
		select {
		case rsp := <-response:
			<-buffer
			slc = append(slc, rsp)
			fmt.Println(rsp)
			if len(slc) == len(h.urls) {
				close(buffer)
				close(response)
				return
			}
		}
	}

}

func (h *HTTPHash) doRequest(url string, buffer chan string, response chan string) {
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
