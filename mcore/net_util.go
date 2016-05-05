package mcore

import (
	"io/ioutil"
	"net/http"
)

// ReadUrl defautl read to []byte
func ReadUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return []byte(body), nil
}

func ReadUrlText(url string) (string, error) {
	b, err := ReadUrl(url)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
