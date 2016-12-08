package goselenium

import (
	"errors"
	"testing"
)

/*
	CREATE SESSION TESTS
*/
func Test_CreateSession_FailedAPIRequestResultsInAnErrorBeingReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: "",
		errorToReturn: &communicationError{
			url: "hello",
			err: errors.New("Call failed! :<"),
		},
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.CreateSession()
	if !IsCommunicationError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_CreateSession_ResultGetsUnmarshalledCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
            "sessionId": "a45a54d3-5413-425c-84ef-d1190cc0521c"
        }`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	resp, err := d.CreateSession()
	if err != nil || resp.SessionID == "" {
		t.Errorf(correctResponseErrorText)
	}
}

func Test_CreateSession_ResultIsAssignedToWebDriver(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
            "sessionId": "a45a54d3-5413-425c-84ef-d1190cc0521c"
        }`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.CreateSession()
	if err != nil || d.sessionID != "a45a54d3-5413-425c-84ef-d1190cc0521c" {
		t.Errorf(correctResponseErrorText)
	}
}

func Test_CreateSession_UnmarshallingErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.CreateSession()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

/*
	DELETE SESSION TESTS
*/
func Test_DeleteSession_WhenSessionIDIsNotSetAnErrorIsThrown(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.DeleteSession()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_DeleteSession_ApiFailureIsHandled(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("This is an error"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.DeleteSession()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_DeleteSession_ResponseIsUnmarshalledCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"sessionId": "3cebaef3-4fd0-464f-bd24-0a7170074ad4"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	resp, err := d.DeleteSession()
	if err != nil || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}
}

func Test_DeleteSession_UnmarshallingFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.DeleteSession()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

/*
	Session Status Test
*/
func Test_SessionStatus_ApiFailureIsHandled(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("This is an error"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.SessionStatus()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_SessionStatusResponse_IsUnmarshalledCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success"	
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	resp, err := d.SessionStatus()
	if err != nil || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}
}

func Test_SessionStatusResponse_UnmarshallingFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.SessionStatus()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

/*
	Session Set Timeout Test
*/
func Test_SetSessionTimeout_ErrorIsThrownIfSessionIdNotSet(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.SetSessionTimeout(nil)
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_SetSessionTimeout_ApiCommunicationErrorIsHandled(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error!"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "1"

	var timeouts = []Timeout{
		SessionScriptTimeout(25000),
		SessionPageLoadTimeout(25000),
		SessionImplicitWaitTimeout(25000),
	}

	for _, i := range timeouts {
		_, err := d.SetSessionTimeout(i)
		if err == nil || !IsCommunicationError(err) {
			t.Errorf(apiCommunicationErrorText)
		}
	}
}

func Test_SetSessionTimeout_ResponseIsUnmarshalledCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "1"

	var timeouts = []Timeout{
		SessionScriptTimeout(25000),
		SessionPageLoadTimeout(25000),
		SessionImplicitWaitTimeout(25000),
	}

	for _, i := range timeouts {
		resp, err := d.SetSessionTimeout(i)
		if err != nil || resp.State != "success" {
			t.Errorf(correctResponseErrorText)
		}
	}
}

func Test_SetSessionTimeout_UnmarshallingFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "1"

	var timeouts = []Timeout{
		SessionScriptTimeout(25000),
		SessionPageLoadTimeout(25000),
		SessionImplicitWaitTimeout(25000),
	}

	for _, i := range timeouts {
		_, err := d.SetSessionTimeout(i)
		if err == nil || !IsUnmarshallingError(err) {
			t.Errorf(unmarshallingErrorText)
		}
	}
}
