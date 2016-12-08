package goselenium

import (
	"errors"
	"testing"
)

/*
   Navigation Go Tests
*/
func Test_NavigateGo_NoSessionIdCausesError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.Go("http://google.com")
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_NavigateGo_InvalidURLFormatResultsInError(t *testing.T) {
	invalidURLs := []string{
		"",
		"google.com",
		"htt://google.com",
		"://google.com",
		"/\\",
	}

	for _, i := range invalidURLs {
		api := &testableAPIService{
			jsonToReturn:  "",
			errorToReturn: nil,
		}

		d := setUpDriver(setUpDefaultCaps(), api)
		d.sessionID = "12345"

		_, err := d.Go(i)
		if err == nil || !IsInvalidURLError(err) {
			t.Errorf("URL error was not returned")
		}
	}
}

func Test_NavigateGo_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error! :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.Go("https://www.google.com")
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_NavigateGo_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.Go("https://www.google.com")
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_NavigateGo_ResultIsUnmarshalledSuccessfully(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
            "state": "success"
        }`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	resp, err := d.Go("https://www.google.com")
	if err != nil || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	CurrentURL tests
*/
func Test_NavigateCurrentURL_InvalidSessionIdResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.CurrentURL()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_NavigateCurrentURL_CommunicationFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("AN error :< "),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.CurrentURL()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_NavigateCurrentURL_UnmarshallingFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.CurrentURL()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_NavigateCurrentURL_SuccessfulResultGetsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": "http://google.com"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	resp, err := d.CurrentURL()
	if err != nil || resp.URL != "http://google.com" || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	Back tests
*/
func Test_NavigateBack_InvalidSessionIdResultsInAnError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.Back()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_NavigateBack_CommunicationFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("AN error :< "),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.Back()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_NavigateBack__UnmarshallingFailureResultsInAnError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.Back()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_NavigateBack_SuccessfulResultGetsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	resp, err := d.Back()
	if err != nil || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	Forward tests
*/
func Test_NavigateForward_InvalidSessionIdResultsInAnError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.Forward()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_NavigateForward_CommunicationFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("AN error :< "),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.Forward()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_NavigateForward__UnmarshallingFailureResultsInAnError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.Forward()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_NavigateForward_SuccessfulResultGetsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	resp, err := d.Forward()
	if err != nil || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	Refresh tests
*/
func Test_NavigateRefresh_InvalidSessionIdResultsInAnError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.Refresh()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_NavigateRefresh_CommunicationFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("AN error :< "),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.Refresh()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_NavigateRefresh_UnmarshallingFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.Refresh()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_NavigateRefresh_SuccessfulResultGetsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	resp, err := d.Refresh()
	if err != nil || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	Title tests
*/
func Test_NavigateTitle_InvalidSessionIdResultsInAnError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.Title()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_NavigateTitle_CommunicationFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("AN error :< "),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.Title()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_NavigateTitle_UnmarshallingFailureResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.Title()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_NavigateTitle_SuccessfulResultGetsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": "Google"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	resp, err := d.Title()
	if err != nil || resp.State != "success" || resp.Title != "Google" {
		t.Errorf(correctResponseErrorText)
	}
}
