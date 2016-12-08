package integrationtests

import (
	"strings"
	"testing"
)

func Test_SessionCreate_ANewSessionCanBeCreated(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	session, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Creation of session yielded an error.", err)
	} else if session.SessionID == "" || session.Capabilities.BrowserName != "firefox" {
		errorAndWrap(t, "Returned object was not set correctly.", err)
	}

	printObjectResult(session)
}

func Test_SessionCreate_TrailingSlashIsAdded(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if strings.HasSuffix(driver.DriverURL(), "/") || err != nil {
		errorAndWrap(t, "Driver URL did not get a slash appended or an error occurred.", err)
	}
}
