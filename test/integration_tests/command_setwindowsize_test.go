package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_CommandSetWindowSize_CorrectResponseIsReturned(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	dimensions := &goselenium.Dimensions{
		Width:  600,
		Height: 400,
	}
	resp, err := driver.SetWindowSize(dimensions)
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error was returned or response was not correct", err)
	}

	printObjectResult(resp)
}
