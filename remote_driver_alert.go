package goselenium

import "fmt"

type DismissAlertResponse struct {
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
