package goselenium

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type apiService interface {
	performRequest(string, string, io.Reader) ([]byte, error)
}

type seleniumAPIService struct{}

func (a *seleniumAPIService) performRequest(url string, method string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code %v returned", resp.StatusCode))
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	r := buf.Bytes()

	return r, nil
}
