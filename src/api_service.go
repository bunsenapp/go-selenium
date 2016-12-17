package goselenium

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type apiService interface {
	performRequest(string, string, io.Reader) ([]byte, error)
}

type requestError struct {
	State string            `json:"state"`
	Value requestErrorValue `json:"value"`
}

type requestErrorValue struct {
	Message string `json:"localizedMessage"`
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

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	r := buf.Bytes()

	if resp.StatusCode != 200 {
		var reqErr requestError
		var errStr string

		err := json.Unmarshal(r, &reqErr)
		if err == nil {
			errStr = fmt.Sprintf("Status code %v returned, message: %v, information: %v", resp.StatusCode, reqErr.State, reqErr.Value.Message)
		} else {
			errStr = fmt.Sprintf("Status code %v returned with no body", resp.StatusCode)
		}
		return nil, errors.New(errStr)
	}

	return r, nil
}
