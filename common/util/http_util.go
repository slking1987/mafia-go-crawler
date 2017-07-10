package util

import (
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url string) (string, error) {
	httpClient := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}

	defer closeRespBody(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func closeRespBody(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
}
