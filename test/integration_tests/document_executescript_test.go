package integrationtests

import "testing"

func Test_DocumentExecuteScript_CanExecuteScriptsSuccessfully(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	_, err = driver.Go("https://news.ycombinator.com")
	if err != nil {
		errorAndWrap(t, "Error was thrown whilst navigating or result was not a success.", err)
	}

	resp, err := driver.ExecuteScript("return \"Test\";")
	if err != nil || resp.Response != "Test" {
		errorAndWrap(t, "Error was thrown whilst executing script or response was not correct", err)
	}

	printObjectResult(resp)
}
