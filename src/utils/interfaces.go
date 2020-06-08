package utils

import (
	"io/ioutil"
	"net/http"
	"time"
)

type HttpRequest interface {
	GetDataApi() (string, error)
}

func (api ApiPublic) GetDataApi(_type string, contentType string) (string, error) {
	timeout := 5 * time.Second
	fecth := http.Client{Timeout: timeout}

	api.Type = _type
	clientRequets, err := ClientHttp(api)
	if err != nil {
		return "", err
	}
	clientRequets.Header.Set("Content-type", contentType)
	clientResponse, err := fecth.Do(clientRequets)
	if err != nil {
		return "", err
	}
	defer clientResponse.Body.Close()
	body, err := ioutil.ReadAll(clientResponse.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}