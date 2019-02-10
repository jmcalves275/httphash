package client

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

	url = resolveURL(url)

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
	hash := md5Hash(bodystring)

	return hash, nil
}

func md5Hash(text string) string {
	data := []byte(text)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func resolveURL(url string) string {
	if strings.Contains(url, "http://") {
		return url
	}
	return fmt.Sprintf("%s%s", "http://", url)
}
