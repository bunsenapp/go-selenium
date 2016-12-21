package goselenium

import (
	"errors"
	"testing"
)

/*
	FIND ELEMENT TESTS
*/
func Test_ElementFindElement_ByIndexResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)

	_, err := d.FindElement(ByIndex(32))
	if err == nil || !IsInvalidArgumentError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_ElementFindElement_InvalidSessionIdResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)

	_, err := d.FindElement(ByCSSSelector("test"))
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_ElementFindElement_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.FindElement(ByCSSSelector("iframe > ul"))
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_ElementFindElement_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.FindElement(ByCSSSelector("iframe > ul"))
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_ElementFindElement_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": {
				"ELEMENT": "0"
			}
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	resp, err := d.FindElement(ByCSSSelector("iframe > ul"))
	if err != nil || resp.ID() != "0" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	FIND ELEMENTS TESTS
*/
func Test_ElementFindElements_ByIndexResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)

	_, err := d.FindElements(ByIndex(32))
	if err == nil || !IsInvalidArgumentError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_ElementFindElements_InvalidSessionIdResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)

	_, err := d.FindElements(ByCSSSelector("test"))
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_ElementFindElements_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.FindElements(ByCSSSelector("iframe > ul"))
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_ElementFindElements_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.FindElements(ByCSSSelector("iframe > ul"))
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_ElementFindElements_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": [
				{"ELEMENT": "0"},
				{"ELEMENT": "1"}
			]
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	resp, err := d.FindElements(ByCSSSelector("iframe > ul"))
	if err != nil || len(resp) <= 1 || resp[1].ID() != "1" {
		t.Errorf(correctResponseErrorText)
	}
}
