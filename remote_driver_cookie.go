package goselenium

import (
	"encoding/json"
	"fmt"
)

// AllCookiesResponse is the response returned from the AllCookies method.
type AllCookiesResponse struct {
	State   string   `json:"state"`
	Cookies []Cookie `json:"value"`
}

// CookieResponse is the response returned from the Cookie method.
type CookieResponse struct {
	State  string `json:"state"`
	Cookie Cookie `json:"value"`
}

// Cookie represents a browser cookie.
type Cookie struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	Path       string `json:"path"`
	Domain     string `json:"domain"`
	SecureOnly bool   `json:"secure"`
	HTTPOnly   bool   `json:"httpOnly"`
}

func (s *seleniumWebDriver) AllCookies() (*AllCookiesResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("AllCookies")
	}

	var response AllCookiesResponse
	var err error

	url := fmt.Sprintf("%s/session/%s/cookie", s.seleniumURL, s.sessionID)

	resp, err := s.apiService.performRequest(url, "GET", nil)
	if err != nil {
		return nil, newCommunicationError(err, "AllCookies", url, nil)
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, "AllCookies", string(resp))
	}

	return &response, nil
}

func (s *seleniumWebDriver) Cookie(name string) (*CookieResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("Cookie")
	}

	var response CookieResponse
	var err error

	url := fmt.Sprintf("%s/session/%s/cookie/%s", s.seleniumURL, s.sessionID, name)

	resp, err := s.apiService.performRequest(url, "GET", nil)
	if err != nil {
		return nil, newCommunicationError(err, "Cookie", url, nil)
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, "Cookie", string(resp))
	}

	return &response, nil
}
