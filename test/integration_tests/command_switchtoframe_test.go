package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium/src"
)

func Test_CommandSwitchToFrame_CorrectResponseIsReturnedByIndex(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	_, err = driver.Go("https://heraclmene.github.io/helpers/goselenium/iframe.html")
	if err != nil {
		errorAndWrap(t, "Error was thrown or result was not a success.", err)
	}

	idx, _ := goselenium.ByIndex(0)
	resp, err := driver.SwitchToFrame(idx)
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

	_, err = driver.Go("https://heraclmene.github.io/helpers/goselenium/iframe.html")
	if err != nil {
		errorAndWrap(t, "Error was thrown or result was not a success.", err)
	}

	selector, _ := goselenium.ByCSSSelector("iframe")
	resp, err := driver.SwitchToFrame(selector)
	if err == nil || !goselenium.IsInvalidArgumentError(err) {
		errorAndWrap(t, "Error was not thrown or was not the expected type.", err)
	}

	printObjectResult(resp)
}
