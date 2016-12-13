package integrationtests

import "testing"

func Test_CommandMaximizeWindow_CorrectResultIsReturned(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	resp, err := driver.MaximizeWindow()
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error was returned or response was not correct", err)
	}

	printObjectResult(resp)
}
