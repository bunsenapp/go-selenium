package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_NavigateGo_CanNavigateSuccessfully(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	resp, err := driver.Go("https://www.google.com")
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error was thrown or result was not a success.", err)
	}
}

func Test_NavigateGo_InvalidURLIsReturned(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	resp, err := driver.Go("www.google.com")
	if err != nil {
		if !goselenium.IsInvalidURLError(err) {
			errorAndWrap(t, "Error was thrown or result was not a success.", err)
		}
	}

	printObjectResult(resp)
}

func Test_NavigateGo_CanGetCurrentURL(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	_, err = driver.Go("https://github.com/")
	if err != nil {
		errorAndWrap(t, "Error was thrown when it shouldn't have been.", err)
	}

	resp, err := driver.CurrentURL()
	if err != nil || resp.URL != "https://github.com/" {
		errorAndWrap(t, "Error was thrown or URL was not what was sent.", err)
	}

	printObjectResult(resp)
}
