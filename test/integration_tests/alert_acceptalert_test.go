package integrationtests

import (
	"testing"
	"time"
)

func Test_AlertAcceptAlert_CanAcceptAnAlertCorrectly(t *testing.T) {
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

	resp, err := driver.AcceptAlert()
	if err != nil || resp.State != "success" {
		t.Errorf("Error returned or accepting an alert was not a success")
	}

	printObjectResult(resp)
}
