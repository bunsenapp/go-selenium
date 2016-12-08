package integrationtests

import "testing"

func Test_CommandWindowHandle_CorrectResponseIsReturned(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	resp, err := driver.WindowHandle()
	if err != nil || resp.State != "success" || resp.Handle == "" {
		errorAndWrap(t, "Error was returned or response was not correct", err)
	}

	printObjectResult(resp)
}
