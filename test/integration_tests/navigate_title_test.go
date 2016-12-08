package integrationtests

import "testing"

func Test_NavigateTitle_TitleCanBeRetrievedSuccessfully(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	goResp, err := driver.Go("https://bbc.co.uk")
	if err != nil || goResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst navigating or result was not a success.", err)
	}

	titleResponse, err := driver.Title()
	if err != nil || titleResponse.Title != "BBC - Home" {
		errorAndWrap(t, "Error was thrown or URL was not what it should have been.", err)
	}

	printObjectResult(titleResponse)
}
