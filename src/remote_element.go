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
