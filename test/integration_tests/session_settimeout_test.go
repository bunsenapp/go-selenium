package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_SessionSetTimeout_CallingSetTimeoutWithoutSessionCausesAnError(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.SetSessionTimeout(goselenium.SessionPageLoadTimeout(20000))
	if err == nil || !goselenium.IsSessionIDError(err) {
		errorAndWrap(t, "Session error was not returned.", err)
	}
}

func Test_SessionSetTimeout_CanSetScriptTimeout(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error returned whilst creating session", err)
	}

	resp, err := driver.SetSessionTimeout(goselenium.SessionScriptTimeout(20000))
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error returned or state was not success whilst setting script timeout.", err)
	}

	printObjectResult(resp)
}

func Test_SessionSetTimeout_CanSetPageLoadTimeout(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error returned whilst creating session", err)
	}

	resp, err := driver.SetSessionTimeout(goselenium.SessionPageLoadTimeout(20000))
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error returned or state was not success whilst setting page load timeout.", err)
	}

	printObjectResult(resp)
}

func Test_SessionSetTimeout_CanSetImplicitWaitTimeout(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error returned whilst creating session", err)
	}

	resp, err := driver.SetSessionTimeout(goselenium.SessionImplicitWaitTimeout(20000))
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error returned or state was not success whilst setting implicit wait timeout.", err)
	}

	printObjectResult(resp)
}
