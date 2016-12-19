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

// ElementTagNameResponse is the response returned from calling the TagName method.
type ElementTagNameResponse struct {
	State string
	Tag   string
}

// ElementRectangleResponse is the response returned from calling the Rectangle
// method.
type ElementRectangleResponse struct {
	State     string
	Rectangle Rectangle `json:"value"`
}

// Rectangle repsents an elements size and position on the page.
type Rectangle struct {
	Dimensions

	X int `json:"x"`
	Y int `json:"y"`
}

// ElementEnabledResponse is the response returned from calling the Enabled method.
type ElementEnabledResponse struct {
	State   string `json:"state"`
	Enabled bool   `json:"value"`
}

// ElementClickResponse is the response returned from calling the Click method.
type ElementClickResponse struct {
	State string
}

// ElementClearResponse is the response returned from calling the Clear method.
type ElementClearResponse struct {
	State string
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

func (s *seleniumElement) TagName() (*ElementTagNameResponse, error) {
	var err error

	url := fmt.Sprintf("%s/session/%s/element/%s/name", s.wd.seleniumURL, s.wd.sessionID, s.ID())

	resp, err := s.wd.valueRequest(&request{
		url:           url,
		method:        "GET",
		body:          nil,
		callingMethod: "TagName",
	})
	if err != nil {
		return nil, err
	}

	return &ElementTagNameResponse{State: resp.State, Tag: resp.Value}, nil
}

func (s *seleniumElement) Rectangle() (*ElementRectangleResponse, error) {
	var response ElementRectangleResponse
	var err error

	url := fmt.Sprintf("%s/session/%s/element/%s/rect", s.wd.seleniumURL, s.wd.sessionID, s.ID())

	resp, err := s.wd.apiService.performRequest(url, "GET", nil)
	if err != nil {
		return nil, newCommunicationError(err, "Rectangle", url, nil)
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, "Rectangle", string(resp))
	}

	return &response, nil
}

func (s *seleniumElement) Enabled() (*ElementEnabledResponse, error) {
	var response ElementEnabledResponse
	var err error

	url := fmt.Sprintf("%s/session/%s/element/%s/enabled", s.wd.seleniumURL, s.wd.sessionID, s.ID())

	resp, err := s.wd.apiService.performRequest(url, "GET", nil)
	if err != nil {
		return nil, newCommunicationError(err, "Enabled", url, nil)
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, "Enabled", string(resp))
	}

	return &response, nil
}

func (s *seleniumElement) Click() (*ElementClickResponse, error) {
	var err error

	url := fmt.Sprintf("%s/session/%s/element/%s/click", s.wd.seleniumURL, s.wd.sessionID, s.ID())

	resp, err := s.wd.stateRequest(&request{
		url:           url,
		method:        "POST",
		body:          nil,
		callingMethod: "Click",
	})
	if err != nil {
		return nil, err
	}

	return &ElementClickResponse{State: resp.State}, nil
}

func (s *seleniumElement) Clear() (*ElementClearResponse, error) {
	var err error

	url := fmt.Sprintf("%s/session/%s/element/%s/clear", s.wd.seleniumURL, s.wd.sessionID, s.ID())

	resp, err := s.wd.stateRequest(&request{
		url:           url,
		method:        "POST",
		body:          nil,
		callingMethod: "Clear",
	})
	if err != nil {
		return nil, err
	}

	return &ElementClearResponse{State: resp.State}, nil
}
