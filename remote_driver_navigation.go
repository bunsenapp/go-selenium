package goselenium

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// GoResponse is the response returned from the selenium web driver when calling
// the Go() call. Unfortunately, the W3C specification defines that the response
// should only be whether the call succeeded or not. Should there be any redirects
// they will not be catered for in this response. Should you expect any redirects
// to happen, call the CurrentURL() method.
type GoResponse struct {
	State string
}

// CurrentURLResponse is the response returned from the GET Url call.
type CurrentURLResponse struct {
	State string
	URL   string
}

// BackResponse is the response returned from the Back call.
type BackResponse struct {
	State string
}

// ForwardResponse is the response returned from the Forward call.
type ForwardResponse struct {
	State string
}

// RefreshResponse is the response returned from the Refresh call.
type RefreshResponse struct {
	State string
}

// TitleResponse is the response returned from the Title call.
type TitleResponse struct {
	State string
	Title string
}

func (s *seleniumWebDriver) Go(goURL string) (*GoResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("Go")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/url", s.seleniumURL, s.sessionID)

	invalidURL := goURL == ""
	validProtocol := strings.HasPrefix(goURL, "https://") || strings.HasPrefix(goURL, "http://")
	if invalidURL || !validProtocol {
		return nil, newInvalidURLError(errors.New("Invalid URL in Go"), goURL)
	}

	params := map[string]string{
		"url": goURL,
	}
	marshalledJSON, err := json.Marshal(params)
	if err != nil {
		return nil, newMarshallingError(err, "Go", params)
	}

	bodyReader := bytes.NewReader([]byte(marshalledJSON))
	resp, err := s.stateRequest(&request{
		url:           url,
		method:        "POST",
		body:          bodyReader,
		callingMethod: "Go",
	})
	if err != nil {
		return nil, err
	}

	return &GoResponse{State: resp.State}, nil
}

func (s *seleniumWebDriver) CurrentURL() (*CurrentURLResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("CurrentURL")
	}

	var response CurrentURLResponse
	var err error

	url := fmt.Sprintf("%s/session/%s/url", s.seleniumURL, s.sessionID)

	resp, err := s.valueRequest(&request{
		url:           url,
		method:        "GET",
		body:          nil,
		callingMethod: "CurrentURL",
	})
	if err != nil {
		return nil, err
	}

	response = CurrentURLResponse{
		State: resp.State,
		URL:   resp.Value,
	}
	return &response, nil
}

func (s *seleniumWebDriver) Back() (*BackResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("Back")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/back", s.seleniumURL, s.sessionID)

	resp, err := s.stateRequest(&request{
		url:           url,
		method:        "POST",
		body:          nil,
		callingMethod: "Back",
	})
	if err != nil {
		return nil, err
	}

	return &BackResponse{State: resp.State}, nil
}

func (s *seleniumWebDriver) Forward() (*ForwardResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("Forward")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/forward", s.seleniumURL, s.sessionID)

	resp, err := s.stateRequest(&request{
		url:           url,
		method:        "POST",
		body:          nil,
		callingMethod: "Forward",
	})
	if err != nil {
		return nil, err
	}

	return &ForwardResponse{State: resp.State}, nil
}

func (s *seleniumWebDriver) Refresh() (*RefreshResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("Refresh")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/refresh", s.seleniumURL, s.sessionID)

	resp, err := s.stateRequest(&request{
		url:           url,
		method:        "POST",
		body:          nil,
		callingMethod: "Refresh",
	})
	if err != nil {
		return nil, err
	}

	return &RefreshResponse{State: resp.State}, nil
}

func (s *seleniumWebDriver) Title() (*TitleResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("Title")
	}

	var response TitleResponse
	var err error

	url := fmt.Sprintf("%s/session/%s/title", s.seleniumURL, s.sessionID)

	resp, err := s.valueRequest(&request{
		url:           url,
		method:        "GET",
		body:          nil,
		callingMethod: "Title",
	})
	if err != nil {
		return nil, err
	}

	response = TitleResponse{
		State: resp.State,
		Title: resp.Value,
	}
	return &response, nil
}
