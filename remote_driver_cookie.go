package goselenium

import (
	"bytes"
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

// AddCookieResponse is the result returned from calling the AddCookie method.
type AddCookieResponse struct {
	State string
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

func (s *seleniumWebDriver) AddCookie(c *Cookie) (*AddCookieResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("AddCookie")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/cookie", s.seleniumURL, s.sessionID)

	j := map[string]Cookie{
		"cookie": *c,
	}
	b, err := json.Marshal(j)
	if err != nil {
		return nil, newMarshallingError(err, "AddCookie", c)
	}
	body := bytes.NewReader(b)
	resp, err := s.stateRequest(&request{
		url:           url,
		body:          body,
		method:        "POST",
		callingMethod: "AddCookie",
	})
	if err != nil {
		return nil, err
	}

	return &AddCookieResponse{State: resp.State}, nil
}
