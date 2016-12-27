package integrationtests

import (
	"testing"

	goselenium "github.com/bunsenapp/go-selenium"
)

func Test_AlertSendAlertText_CanSendAlertTextCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		t.Errorf("Create session error")
	}

	_, err = driver.Go("https://heraclmene.github.io/helpers/goselenium/prompt.html")
	if err != nil {
		t.Errorf("Error visiting URL")
	}

	resp, err := driver.SendAlertText("test")
	if err != nil || resp.State != "success" {
		t.Errorf("Error returned or sending alert text was not a success")
	}

	_, err = driver.AcceptAlert()
	if err != nil {
		t.Errorf("Error was returned when accepting alert.")
	}

	_, err = driver.AlertText()
	if err != nil {
		comErr := err.(goselenium.CommunicationError)
		if comErr.Response.State != goselenium.NoSuchAlert {
			t.Errorf("Error returned was not correct.")
		}
	} else {
		t.Errorf("Error returned was not correct.")
	}

	printObjectResult(resp)
}
