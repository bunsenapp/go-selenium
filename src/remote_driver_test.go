package goselenium

import (
	"io"
	"strings"
	"testing"
)

const (
	apiCommunicationErrorText = "An error was not returned or was not of the API communication type"
	sessionIDErrorText        = "An error was not returned or was not of the SessionIDError type"
	correctResponseErrorText  = "An error was returned or the result was not what was expected"
	argumentErrorText         = "An error was not returned or was not of the ArgumentError type"
	unmarshallingErrorText    = "An error was not returned or was not of the UnmarshallingError type"
)

func setUpDefaultCaps() *Capabilities {
	caps := Capabilities{}
	caps.SetBrowser(FirefoxBrowser())
	return &caps
}

func setUpDriver(caps *Capabilities, api apiService) *seleniumWebDriver {
	return &seleniumWebDriver{
		seleniumURL:  "http://localhost:4444/wd/hub/",
		capabilities: caps,
		apiService:   api,
	}
}

type testableAPIService struct {
	jsonToReturn  string
	errorToReturn error
	bodyNilError  error
}

func (t *testableAPIService) performRequest(url string, method string, body io.Reader) ([]byte, error) {
	json := []byte(t.jsonToReturn)
	return json, t.errorToReturn
}

func Test_NewSelenium_WebDriverCreatesErrorIfSeleniumURLIsInvalid(t *testing.T) {
	invalidSeleniumUrls := []string{
		"",
		" ",
		"myRequirementWithoutProtocol",
	}
	for _, i := range invalidSeleniumUrls {
		caps := setUpDefaultCaps()
		_, err := NewSeleniumWebDriver(i, *caps)
		if err == nil {
			t.Errorf("Passing an invalid remote Selenium URL did not cause an error")
		}
	}
}

func Test_NewSelenium_WebDriverCreatesSuccessfullyIfSeleniumURLIsValid(t *testing.T) {
	validSeleniumUrls := []string{
		"http://google.com",
		"https://google.com",
	}
	for _, i := range validSeleniumUrls {
		caps := setUpDefaultCaps()
		w, err := NewSeleniumWebDriver(i, *caps)
		if w == nil || err != nil {
			t.Errorf("Passing a valid remote Selenium URL caused an error or did not return a valid driver.")
		}
	}
}

func Test_NewSelenium_WebDriverCreatesErrorIfCapabilitiesAreEmpty(t *testing.T) {
	_, err := NewSeleniumWebDriver("http://google.com", Capabilities{})
	if err == nil {
		t.Errorf("Passing an empty capabilities object did not cause an error.")
	}
}

func Test_NewSelenium_TrailingSlashIsRemovedIfTheUserDoesNotSpecifyOne(t *testing.T) {
	invalidUrls := []string{
		"http://localhost/",
		"http://localhost:444/",
	}
	for _, i := range invalidUrls {
		caps := setUpDefaultCaps()
		d, err := NewSeleniumWebDriver(i, *caps)
		if err != nil || strings.HasSuffix(d.DriverURL(), "/") {
			t.Errorf("Trailing slash was not removed from URL or an error was returned.")
		}
	}
}

/*
	By tests
*/
func Test_ByByIndex_OutOfRangeIndexReturnsError(t *testing.T) {
	outOfRangeIndexes := []uint{
		65536,
		234234,
		45694569,
	}
	for _, i := range outOfRangeIndexes {
		_, err := ByIndex(i)
		if err == nil || !IsInvalidArgumentError(err) {
			t.Errorf(argumentErrorText)
		}
	}
}

func Test_ByByIndex_CorrectIndexReturnsAsExpected(t *testing.T) {
	correctIndexes := []uint{
		1,
		58,
		65535,
	}
	for _, i := range correctIndexes {
		r, err := ByIndex(i)
		if err != nil || r.Type() != "index" || r.Value().(uint) != i {
			t.Errorf(correctResponseErrorText)
		}
	}
}

func Test_ByByCSSSelector_EmptyClassReturnsAnError(t *testing.T) {
	_, err := ByCSSSelector("")
	if err == nil || !IsInvalidArgumentError(err) {
		t.Errorf(argumentErrorText)
	}
}
