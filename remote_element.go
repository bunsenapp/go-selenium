package goselenium

import (
	"encoding/json"
	"fmt"
)

func newSeleniumElement(i string, w *seleniumWebDriver) *seleniumElement {
	return &seleniumElement{
		id: i,
		wd: w,
	}
}

// ElementSelectedResponse is the response returned from the Selected() call.
// The result /should/ always be successfully returned unless there is a
// server error.
type ElementSelectedResponse struct {
	State    string `json:"state"`
	Selected bool   `json:"value"`
}

// ElementAttributeResponse is the response returned from the Attribute call.
type ElementAttributeResponse struct {
	State string
	Value string
}

// ElementCSSValueResponse is the response returned when the CSSValue method
// is called on an Element implementation.
type ElementCSSValueResponse struct {
	State string
	Value string
}

// ElementTextResponse is the response returned from calling the Text method.
type ElementTextResponse struct {
	State string
	Text  string
}

type seleniumElement struct {
	id string
	wd *seleniumWebDriver
}

func (s *seleniumElement) ID() string {
	return s.id
}

func (s *seleniumElement) Selected() (*ElementSelectedResponse, error) {
	var el ElementSelectedResponse
	var err error

	url := fmt.Sprintf("%s/session/%s/element/%s/selected", s.wd.seleniumURL, s.wd.sessionID, s.ID())

	resp, err := s.wd.apiService.performRequest(url, "GET", nil)
	if err != nil {
		return nil, newCommunicationError(err, "Selected", url, nil)
	}

	err = json.Unmarshal(resp, &el)
	if err != nil {
		return nil, newUnmarshallingError(err, "Selected", string(resp))
	}

	return &el, nil
}

func (s *seleniumElement) Attribute(att string) (*ElementAttributeResponse, error) {
	var err error

	url := fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", s.wd.seleniumURL, s.wd.sessionID, s.ID(), att)

	resp, err := s.wd.valueRequest(&request{
		url:           url,
		method:        "GET",
		body:          nil,
		callingMethod: "Attribute",
	})
	if err != nil {
		return nil, err
	}

	return &ElementAttributeResponse{State: resp.State, Value: resp.Value}, nil
}

func (s *seleniumElement) CSSValue(prop string) (*ElementCSSValueResponse, error) {
	var err error

	url := fmt.Sprintf("%s/session/%s/element/%s/css/%s", s.wd.seleniumURL, s.wd.sessionID, s.ID(), prop)

	resp, err := s.wd.valueRequest(&request{
		url:           url,
		method:        "GET",
		body:          nil,
		callingMethod: "CSSValue",
	})
	if err != nil {
		return nil, err
	}

	return &ElementCSSValueResponse{State: resp.State, Value: resp.Value}, nil
}

func (s *seleniumElement) Text() (*ElementTextResponse, error) {
	var err error

	url := fmt.Sprintf("%s/session/%s/element/%s/text", s.wd.seleniumURL, s.wd.sessionID, s.ID())

	resp, err := s.wd.valueRequest(&request{
		url:           url,
		method:        "GET",
		body:          nil,
		callingMethod: "Text",
	})
	if err != nil {
		return nil, err
	}

	return &ElementTextResponse{State: resp.State, Text: resp.Value}, nil
}
