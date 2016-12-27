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
		t.Errorf("Create session error")
	}

	_, err = driver.Go("https://heraclmene.github.io/helpers/goselenium/alert.html")
	if err != nil {
		t.Errorf("Error visiting URL")
	}

	resp, err := driver.DismissAlert()
	if err != nil || resp.State != "success" {
		t.Errorf("Error returned or dismissing an alert was not a success")
	}

	printObjectResult(resp)
}

func Test_AlertDismissAlert_DismissingAnInvalidAlertResultsInAnError(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		t.Errorf("Create session error")
	}

	_, err = driver.Go("https://google.com")
	if err != nil {
		t.Errorf("Error visiting URL")
	}

	resp, err := driver.DismissAlert()
	if err != nil {
		comErr := err.(goselenium.CommunicationError)
		if comErr.Response.State != goselenium.NoSuchAlert {
			t.Errorf("Incorrect error was returned.")
		}
	} else {
		t.Errorf("Error was not returned when it should have been.")
	}

	printObjectResult(resp)
}
