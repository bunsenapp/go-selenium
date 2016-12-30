package integrationtests

import (
	"testing"

	goselenium "github.com/bunsenapp/go-selenium"
)

func Test_AlertDismissAlert_CanDismissAnAlertCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error creating session", err)
	}

	_, err = driver.Go("https://bunsenapp.github.io/go-selenium/helpers/alert.html")
	if err != nil {
		errorAndWrap(t, "Error visiting URL", err)
	}

	resp, err := driver.DismissAlert()
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Alert response was not correct", err)
	}

	printObjectResult(resp)
}

func Test_AlertDismissAlert_DismissingAnInvalidAlertResultsInAnError(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error creating session", err)
	}

	_, err = driver.Go("https://google.com")
	if err != nil {
		errorAndWrap(t, "Error visiting URL", err)
	}

	resp, err := driver.DismissAlert()
	if err != nil {
		comErr := err.(goselenium.CommunicationError)
		if comErr.Response.State != goselenium.NoSuchAlert {
			errorAndWrap(t, "Incorrect result returned", err)
		}
	} else {
		errorAndWrap(t, "Incorrect result returned", err)
	}

	printObjectResult(resp)
}
