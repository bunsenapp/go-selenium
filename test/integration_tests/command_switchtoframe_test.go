package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_CommandSwitchToFrame_CorrectResponseIsReturnedByIndex(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	_, err = driver.Go("https://bunsenapp.github.io/go-selenium/helpers/iframe.html")
	if err != nil {
		errorAndWrap(t, "Error was thrown or result was not a success.", err)
	}

	resp, err := driver.SwitchToFrame(goselenium.ByIndex(0))
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error was thrown or result was not a success", err)
	}

	printObjectResult(resp)
}

func Test_CommandSwitchToFrame_InvalidByResultsInAnError(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	_, err = driver.Go("https://bunsenapp.github.io/go-selenium/helpers/iframe.html")
	if err != nil {
		errorAndWrap(t, "Error was thrown or result was not a success.", err)
	}

	resp, err := driver.SwitchToFrame(goselenium.ByCSSSelector("iframe"))
	if err == nil {
		errorAndWrap(t, "Error was not thrown or was not the expected type.", err)
	}

	printObjectResult(resp)
}
