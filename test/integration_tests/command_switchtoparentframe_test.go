package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_CommandSwitchToParentFrame_CorrectResponseCanBeReturned(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	_, err = driver.Go("https://bunsenapp.github.io/go-selenium/helpers/iframe.html")
	if err != nil {
		errorAndWrap(t, "Error was thrown whilst navigating.", err)
	}

	_, err = driver.SwitchToFrame(goselenium.ByIndex(0))
	if err != nil {
		errorAndWrap(t, "Error was thrown whilst switching to frame 0.", err)
	}

	resp, err := driver.SwitchToParentFrame()
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error was thrown or response was not a success.", err)
	}

	printObjectResult(resp)
}
