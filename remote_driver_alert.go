package goselenium

import "fmt"

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
