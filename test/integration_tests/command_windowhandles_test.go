package integrationtests

import "testing"

func Test_CommandWindowHandles_WindowHandlesAreReturned(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	resp, err := driver.WindowHandles()
	if err != nil || resp.State != "success" || resp.Handles[0] == "" {
		errorAndWrap(t, "Error thrown or result was not what was expected.", err)
	}

	printObjectResult(resp)
}
