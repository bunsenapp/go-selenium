package goselenium

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// DismissAlertResponse is the response returned from calling the DismissAlert
// method.
type DismissAlertResponse struct {
	State string
}

// AcceptAlertResponse is the response returned from calling the AcceptAlert
// method.
type AcceptAlertResponse struct {
	State string
}

// AlertTextResponse is the response returned from calling the AlertText
// method.
type AlertTextResponse struct {
	State string
	Text  string
}

// SendAlertTextResponse is the response returned from calling the
// SendAlertText method.
type SendAlertTextResponse struct {
	State string
}

func (s *seleniumWebDriver) DismissAlert() (*DismissAlertResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("DismissAlert")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/alert/dismiss", s.seleniumURL, s.sessionID)

	resp, err := s.stateRequest(&request{
		url:           url,
		method:        "POST",
		body:          nil,
		callingMethod: "DismissAlert",
	})
	if err != nil {
		return nil, err
	}

	return &DismissAlertResponse{State: resp.State}, nil
}

func (s *seleniumWebDriver) AcceptAlert() (*AcceptAlertResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("AcceptAlert")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/alert/accept", s.seleniumURL, s.sessionID)

	resp, err := s.stateRequest(&request{
		url:           url,
		method:        "POST",
		body:          nil,
		callingMethod: "AcceptAlert",
	})
	if err != nil {
		return nil, err
	}

	return &AcceptAlertResponse{State: resp.State}, nil
}

func (s *seleniumWebDriver) AlertText() (*AlertTextResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("AlertTextResponse")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/alert/text", s.seleniumURL, s.sessionID)

	resp, err := s.valueRequest(&request{
		url:           url,
		method:        "GET",
		body:          nil,
		callingMethod: "AlertText",
	})
	if err != nil {
		return nil, err
	}

	return &AlertTextResponse{State: resp.State, Text: resp.Value}, nil
}

func (s *seleniumWebDriver) SendAlertText(text string) (*SendAlertTextResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("SendAlertText")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/alert/text", s.seleniumURL, s.sessionID)

	b := map[string]string{
		"text": text,
	}
	json, err := json.Marshal(b)
	if err != nil {
		return nil, newMarshallingError(err, "SendAlertText", b)
	}

	body := bytes.NewReader(json)
	resp, err := s.valueRequest(&request{
		url:           url,
		method:        "POST",
		body:          body,
		callingMethod: "SendAlertText",
	})
	if err != nil {
		return nil, err
	}

	return &SendAlertTextResponse{State: resp.State}, nil
}
