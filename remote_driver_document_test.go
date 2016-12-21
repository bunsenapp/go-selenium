package goselenium

import (
	"errors"
	"testing"
)

/*
	PageSource tests
*/
func Test_DocumentPageSource_InvalidSessionIDResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)

	_, err := d.PageSource()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_DocumentPageSource_CommunicationErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.PageSource()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_DocumentPageSource_UnmarshallingErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.PageSource()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_DocumentPageSource_ResultIsReturnedSuccessfully(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": "this would be HTML"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	resp, err := d.PageSource()
	if err != nil || resp.State != "success" || resp.Source != "this would be HTML" {
		t.Errorf(correctResponseErrorText)
	}
}
