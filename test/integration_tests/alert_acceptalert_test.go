package integrationtests

import "testing"

func Test_AlertAcceptAlert_CanAcceptAnAlertCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Create session error", err)
	}

	_, err = driver.Go("https://heraclmene.github.io/helpers/goselenium/alert.html")
	if err != nil {
		errorAndWrap(t, "Error navigating to URL", err)
	}

	resp, err := driver.AcceptAlert()
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error accepting alert", err)
	}

	printObjectResult(resp)
}
