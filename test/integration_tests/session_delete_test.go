package integrationtests

import (
	"testing"

	goselenium "github.com/bunsenapp/go-selenium/src"
)

func Test_SessionDelete_CallingDeleteSessionMethodWithoutASessionIdResultsInAnError(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.DeleteSession()
	if err == nil || !goselenium.IsSessionIDError(err) {
		errorAndWrap(t, "Session error was not returned when it should have been.", err)
	}
}

func Test_SessionDelete_DeleteSessionMethodWorksCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error returned whilst creating session", err)
	}

	resp, err := driver.DeleteSession()
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error returned or state was not success whilst deleting session.", err)
	}

	printObjectResult(resp)
}
