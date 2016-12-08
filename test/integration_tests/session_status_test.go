package integrationtests

import "testing"

func Test_SessionStatus_CanRetrieveStatusOfDriverSuccessfully(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	resp, err := driver.SessionStatus()
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error occurred or state is incorrect", err)
	}

	printObjectResult(resp)
}
