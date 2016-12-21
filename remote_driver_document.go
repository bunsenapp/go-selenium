package goselenium

import "fmt"

// PageSourceResponse is the response returned from calling the PageSource
// method.
type PageSourceResponse struct {
	State  string
	Source string
}

// ExecuteScriptResponse is the response returned from calling the ExecuteScript
// method.
type ExecuteScriptResponse struct {
	State    string
	Response string
}

func (s *seleniumWebDriver) PageSource() (*PageSourceResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("PageSource")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/source", s.seleniumURL, s.sessionID)

	resp, err := s.valueRequest(&request{
		url:           url,
		method:        "GET",
		body:          nil,
		callingMethod: "PageSource",
	})
	if err != nil {
		return nil, err
	}

	return &PageSourceResponse{State: resp.State, Source: resp.Value}, nil
}

func (s *seleniumWebDriver) ExecuteScript(script string) (*ExecuteScriptResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("ExecuteScript")
	}

	url := fmt.Sprintf("%s/session/%s/execute", s.seleniumURL, s.sessionID)

	return s.scriptRequest(script, url, "ExecuteScript")
}

func (s *seleniumWebDriver) ExecuteScriptAsync(script string) (*ExecuteScriptResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("ExecuteScriptAsync")
	}

	url := fmt.Sprintf("%s/session/%s/execute_async", s.seleniumURL, s.sessionID)

	return s.scriptRequest(script, url, "ExecuteScriptAsync")
}
