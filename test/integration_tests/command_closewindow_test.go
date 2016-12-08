package integrationtests

import "testing"

func Test_CommandCloseWindow_CanCloseTheWindow(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	resp, err := driver.CloseWindow()
	if err != nil || resp.State != "success" || len(resp.Handles) > 0 {
		errorAndWrap(t, "Error was returned or response was not correct", err)
	}

	printObjectResult(resp)
}
