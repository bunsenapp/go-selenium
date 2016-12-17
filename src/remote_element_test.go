package goselenium

import (
	"errors"
	"fmt"
	"testing"
)

func Test_RemoteElement_IDCanBeRetrieved(t *testing.T) {
	el := newSeleniumElement("test", nil)
	if el.ID() != "test" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	SELECTED TESTS
*/
func Test_ElementSelected_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.Selected()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_ElementSelected_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.Selected()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_ElementSelected_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": true
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	resp, err := el.Selected()
	if err != nil || resp.State != "success" || resp.Selected != true {
		fmt.Println(err)
		fmt.Println(resp)
		t.Errorf(correctResponseErrorText)
	}
}
