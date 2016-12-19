package goselenium

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/pkg/errors"
)

// NewSeleniumWebDriver creates a new instance of a Selenium web driver with a
// service URL (usually http://domain:port/wd/hub) and a Capabilities object.
// This method will return validation errors if the Selenium URL is invalid or
// the required capabilities (BrowserName) are not set.
func NewSeleniumWebDriver(serviceURL string, capabilities Capabilities) (WebDriver, error) {
	if serviceURL == "" {
		return nil, errors.New("Provided Selenium URL is invalid")
	}

	urlValid := strings.HasPrefix(serviceURL, "http://") || strings.HasPrefix(serviceURL, "https://")
	if !urlValid {
		return nil, errors.New("Provided Selenium URL is invalid.")
	}

	browser := capabilities.Browser()
	hasBrowserCapability := browser.BrowserName() != ""
	if !hasBrowserCapability {
		return nil, errors.New("An invalid capabilities object was provided.")
	}

	if strings.HasSuffix(serviceURL, "/") {
		serviceURL = strings.TrimSuffix(serviceURL, "/")
	}

	driver := &seleniumWebDriver{
		seleniumURL:  serviceURL,
		capabilities: &capabilities,
		apiService:   &seleniumAPIService{},
	}

	return driver, nil
}

// SessionScriptTimeout creates an appropriate Timeout implementation for the
// script timeout.
func SessionScriptTimeout(to int) Timeout {
	return &timeout{
		timeoutType: "script",
		timeout:     to,
	}
}

// SessionPageLoadTimeout creates an appropriate Timeout implementation for the
// page load timeout.
func SessionPageLoadTimeout(to int) Timeout {
	return &timeout{
		timeoutType: "page load",
		timeout:     to,
	}
}

// SessionImplicitWaitTimeout creates an appropriate timeout implementation for the
// session implicit wait timeout.
func SessionImplicitWaitTimeout(to int) Timeout {
	return &timeout{
		timeoutType: "implicit",
		timeout:     to,
	}
}

// ByIndex accepts an integer that represents what the index of an element is
// and returns the appropriate By implementation.
func ByIndex(index uint) By {
	return &by{
		t:     "index",
		value: index,
	}
}

// ByCSSSelector accepts a CSS selector (i.e. ul#id > a) for use in the
// FindElement(s) functions.
func ByCSSSelector(selector string) By {
	return &by{
		t:     "css selector",
		value: selector,
	}
}

// ByLinkText is used to find an anchor element by its innerText.
func ByLinkText(text string) By {
	return &by{
		t:     "link text",
		value: text,
	}
}

// ByPartialLinkText works the same way as ByLinkText but performs a search
// where the link text contains the string passed in instead of a full match.
func ByPartialLinkText(text string) By {
	return &by{
		t:     "partial link text",
		value: text,
	}
}

// ByXPath utilises the xpath to find elements (see http://www.guru99.com/xpath-selenium.html).
func ByXPath(path string) By {
	return &by{
		t:     "xpath",
		value: path,
	}
}

type seleniumWebDriver struct {
	seleniumURL  string
	sessionID    string
	capabilities *Capabilities
	apiService   apiServicer
}

func (s *seleniumWebDriver) DriverURL() string {
	return s.seleniumURL
}

func (s *seleniumWebDriver) stateRequest(req *request) (*stateResponse, error) {
	var response stateResponse
	var err error

	resp, err := s.apiService.performRequest(req.url, req.method, req.body)
	if err != nil {
		return nil, newCommunicationError(err, req.callingMethod, req.url, resp)
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, req.callingMethod, string(resp))
	}

	return &response, nil
}

func (s *seleniumWebDriver) valueRequest(req *request) (*valueResponse, error) {
	var response valueResponse
	var err error

	resp, err := s.apiService.performRequest(req.url, req.method, req.body)
	if err != nil {
		return nil, newCommunicationError(err, req.callingMethod, req.url, resp)
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, req.callingMethod, string(resp))
	}

	return &response, nil
}

func (s *seleniumWebDriver) elementRequest(req *elRequest) ([]byte, error) {
	b := map[string]interface{}{
		"using": req.by.Type(),
		"value": req.by.Value(),
	}
	bJSON, err := json.Marshal(b)
	if err != nil {
		return nil, newMarshallingError(err, req.callingMethod, bJSON)
	}

	body := bytes.NewReader(bJSON)
	resp, err := s.apiService.performRequest(req.url, req.method, body)
	if err != nil {
		return nil, newCommunicationError(err, req.callingMethod, req.url, resp)
	}

	return resp, nil
}

type timeout struct {
	timeoutType string
	timeout     int
}

func (t *timeout) Type() string {
	return t.timeoutType
}

func (t *timeout) Timeout() int {
	return t.timeout
}

type request struct {
	url           string
	method        string
	body          io.Reader
	callingMethod string
}

type elRequest struct {
	url           string
	by            By
	method        string
	callingMethod string
}

type stateResponse struct {
	State string `json:"state"`
}

type valueResponse struct {
	State string `json:"state"`
	Value string `json:"value"`
}

type by struct {
	t     string
	value interface{}
}

func (b *by) Type() string {
	return b.t
}

func (b *by) Value() interface{} {
	return b.value
}
