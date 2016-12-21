package integrationtests

import "testing"

func Test_DocumentPageSource_PageSourceIsCorrectlyRetrieved(t *testing.T) {
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

	sauce, err := driver.PageSource()
	if err != nil || len(sauce.Source) == 0 {
		errorAndWrap(t, "Error was thrown or page source was empty", err)
	}

	printObjectResult(sauce)
}
