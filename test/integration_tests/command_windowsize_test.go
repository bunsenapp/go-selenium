package integrationtests

import "testing"

func Test_CommandWindowSize_CorrectResultIsReturned(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	resp, err := driver.WindowSize()
	if err != nil || resp.State != "success" || resp.Dimensions.Width == 0 || resp.Dimensions.Height == 0 {
		errorAndWrap(t, "Error was returned or response was not correct", err)
	}

	printObjectResult(resp)
}
