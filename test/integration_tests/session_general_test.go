package integrationtests

import "testing"

func Test_Session_CanCreateAndDeleteSession(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Creation of session yielded an error.", err)
	}

	_, err = driver.DeleteSession()
	if err != nil {
		errorAndWrap(t, "Creation of session yielded an error.", err)
	}
}

func Test_Session_CanCreateSessionAndGetStatus(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Creation of session yielded an error.", err)
	}

	resp, err := driver.SessionStatus()
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error occurred or the state was not a success", err)
	}
}
