package integrationtests

import (
	"strings"
	"testing"
)

func Test_NavigateForward_NavigateFowardWorksCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	goResp, err := driver.Go("https://google.co.uk")
	if err != nil || goResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst navigating or result was not a success.", err)
	}

	goResp, err = driver.Go("https://bbc.co.uk")
	if err != nil || goResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst navigating or result was not a success.", err)
	}

	backResp, err := driver.Back()
	if err != nil || backResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst navigating backwards or results was not a success.", err)
	}

	forwardResp, err := driver.Forward()
	if err != nil || forwardResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst navigating forwards or result was not a success.", err)
	}

	currentURLResp, err := driver.CurrentURL()
	if err != nil || !strings.HasPrefix(currentURLResp.URL, "https://www.bbc.co.uk") {
		errorAndWrap(t, "Error was thrown or URL was not what it should have been.", err)
	}

	printObjectResult(currentURLResp)
}
