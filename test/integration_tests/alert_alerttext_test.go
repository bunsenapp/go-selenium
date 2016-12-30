package integrationtests

import "testing"

func Test_AlertAlertText_CanGetTheAlertText(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error creating session.", err)
	}

	_, err = driver.Go("https://bunsenapp.github.io/go-selenium/helpers/alert.html")
	if err != nil {
		errorAndWrap(t, "Error visiting URL.", err)
	}

	resp, err := driver.AlertText()
	if err != nil || resp.State != "success" || resp.Text != "this is an alert" {
		errorAndWrap(t, "Error getting alert text.", err)
	}

	printObjectResult(resp)
}
