package goselenium

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type apiServicer interface {
	performRequest(string, string, io.Reader) ([]byte, error)
}

type requestError struct {
	State string            `json:"state"`
	Value requestErrorValue `json:"value"`
}

func (r *requestError) Error() string {
	return fmt.Sprintf("Invalid status code returned, message: %v, information: %v", r.State, r.Value.Message)
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
		return nil, errors.Wrap(err, "An unexpected error occurred")
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)
	r := buf.Bytes()

	if resp.StatusCode != 200 {
		var reqErr requestError
		var errStr string

		err := json.Unmarshal(r, &reqErr)
		if err == nil {
			return nil, &reqErr
		} else {
			errStr = fmt.Sprintf("Status code %v returned with no body", resp.StatusCode)
			return nil, errors.New(errStr)
		}
	}

	return r, nil
}
