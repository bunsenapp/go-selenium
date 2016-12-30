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
		errorAndWrap(t, "Error creating session", err)
	}

	_, err = driver.Go("https://bunsenapp.github.io/go-selenium/helpers/prompt.html")
	if err != nil {
		errorAndWrap(t, "Error visiting URL", err)
	}

	resp, err := driver.SendAlertText("test")
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error sending alert text", err)
	}

	_, err = driver.AcceptAlert()
	if err != nil {
		errorAndWrap(t, "Error accepting alert", err)
	}

	_, err = driver.AlertText()
	if err != nil {
		comErr := err.(goselenium.CommunicationError)
		if comErr.Response.State != goselenium.NoSuchAlert {
			errorAndWrap(t, "Error returned was not correct", err)
		}
	} else {
		errorAndWrap(t, "Error returned was not correct", err)
	}

	printObjectResult(resp)
}
