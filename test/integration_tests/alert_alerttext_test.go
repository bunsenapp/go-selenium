package integrationtests

import (
	"testing"
	"time"
)

func Test_AlertAlertText_CanGetTheAlertText(t *testing.T) {
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

	time.Sleep(time.Second)

	resp, err := driver.AlertText()
	if err != nil || resp.State != "success" || resp.Text != "this is an alert" {
		t.Errorf("Error returned or alert text was not correct")
	}

	printObjectResult(resp)
}
