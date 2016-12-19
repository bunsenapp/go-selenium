package goselenium

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CreateSessionResponse is the response returned from the API when the
// CreateSession() method does not throw an error.
type CreateSessionResponse struct {
	Capabilities CreateSessionCapabilities `json:"value"`
	SessionID    string                    `json:"sessionId"`
}

// CreateSessionCapabilities is a summarisation of the capabilities returned
// from the CreateSession method.
type CreateSessionCapabilities struct {
	AcceptInsecureCerts bool   `json:"acceptSslCerts"`
	BrowserName         string `json:"browserName"`
	BrowserVersion      string `json:"browserVersion"`
	PlatformName        string `json:"platformVersion"`
}

// DeleteSessionResponse is the response returned from the API when the
// DeleteSession() method does not thrown an error.
type DeleteSessionResponse struct {
	State     string `json:"state"`
	SessionID string `json:"sessionId"`
}

// SessionStatusResponse is the response returned from the API when the
// SessionStatus() method is called.
type SessionStatusResponse struct {
	State string
}

// SetSessionTimeoutResponse is the response returned from the API when the
// SetSessionTimeoutResponse() method is called.
type SetSessionTimeoutResponse struct {
	State string
}

func (s *seleniumWebDriver) CreateSession() (*CreateSessionResponse, error) {
	var response CreateSessionResponse
	var err error

	url := fmt.Sprintf("%s/session", s.seleniumURL)

	capabilitiesJSON, err := s.capabilities.toJSON()
	if err != nil {
		return nil, newMarshallingError(err, "CreateSession", s.capabilities)
	}

	body := bytes.NewReader([]byte(capabilitiesJSON))
	resp, err := s.apiService.performRequest(url, "POST", body)
	if err != nil {
		return nil, newCommunicationError(err, "CreateSession", url, resp)
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, "CreateSession", string(resp))
	}

	s.sessionID = response.SessionID
	return &response, nil
}

func (s *seleniumWebDriver) DeleteSession() (*DeleteSessionResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("DeleteSession")
	}

	var response DeleteSessionResponse
	var err error

	url := fmt.Sprintf("%s/session/%s", s.seleniumURL, s.sessionID)

	resp, err := s.apiService.performRequest(url, "DELETE", nil)
	if err != nil {
		return nil, newCommunicationError(err, "DeleteSession", url, resp)
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, "DeleteSession", string(resp))
	}

	return &response, nil
}

func (s *seleniumWebDriver) SessionStatus() (*SessionStatusResponse, error) {
	var err error

	url := fmt.Sprintf("%s/status", s.seleniumURL)

	resp, err := s.stateRequest(&request{
		url:           url,
		method:        "GET",
		body:          nil,
		callingMethod: "SessionStatus",
	})
	if err != nil {
		return nil, err
	}

	return &SessionStatusResponse{State: resp.State}, nil
}

func (s *seleniumWebDriver) SetSessionTimeout(to Timeout) (*SetSessionTimeoutResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("SetSessionTimeout")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/timeouts", s.seleniumURL, s.sessionID)

	params := map[string]interface{}{
		"type": to.Type(),
		"ms":   to.Timeout(),
	}
	marshalledJSON, err := json.Marshal(params)
	if err != nil {
		return nil, newMarshallingError(err, "SetSessionTimeout", params)
	}

	bodyReader := bytes.NewReader([]byte(marshalledJSON))
	resp, err := s.stateRequest(&request{
		url:           url,
		method:        "POST",
		body:          bodyReader,
		callingMethod: "SetSessionTimeout",
	})
	if err != nil {
		return nil, err
	}

	return &SetSessionTimeoutResponse{State: resp.State}, nil
}
