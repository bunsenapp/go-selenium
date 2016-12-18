package integrationtests

import (
	"strings"
	"testing"
)

func Test_NavigateRefresh_RefreshWorksCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	goResp, err := driver.Go("https://news.ycombinator.com")
	if err != nil || goResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst navigating or result was not a success.", err)
	}

	refreshResp, err := driver.Refresh()
	if err != nil || refreshResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst refreshing or result was not a success.", err)
	}

	currentURLResp, err := driver.CurrentURL()
	if err != nil || !strings.HasPrefix(currentURLResp.URL, "https://news.ycombinator.com") {
		errorAndWrap(t, "Error was thrown or URL was not what it should have been.", err)
	}

	printObjectResult(currentURLResp)
}
