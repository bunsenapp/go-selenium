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
	State string `json:"state"`
	Value string `json:"value"`
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
	var response ElementAttributeResponse
	var err error

	url := fmt.Sprintf("%s/session/%s/element/%s/attribute/%s", s.wd.seleniumURL, s.wd.sessionID, s.ID(), att)

	resp, err := s.wd.apiService.performRequest(url, "GET", nil)
	if err != nil {
		return nil, newCommunicationError(err, "Attribute", url, nil)
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, "Attribute", string(resp))
	}

	return &response, nil
}
