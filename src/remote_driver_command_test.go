package goselenium

import (
	"errors"
	"testing"
)

/*
	WindowHandle() Tests
*/

func Test_CommandWindowHandle_InvalidSessionIdResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.WindowHandle()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_CommandWindowHandle_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.WindowHandle()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_CommandWindowHandle_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.WindowHandle()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_CommandWindowHandle_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": "8"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	resp, err := d.WindowHandle()
	if err != nil || resp.State != "success" || resp.Handle != "8" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	CloseWindow() Tests
*/
func Test_CommandCloseWindow_InvalidSessionIdResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.CloseWindow()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_CommandCloseWindow_CommunicationErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.CloseWindow()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}

}

func Test_CommandCloseWindow_UnmarshallingErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.CloseWindow()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

/*
	SwitchToWindow() Tests
*/

/*
	WindowHandles() Tests
*/
func Test_CommandWindowHandles_InvalidSessionIdResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.WindowHandles()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_CommandWindowHandles_CommunicationErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.WindowHandles()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_CommandWindowHandles_UnmarshallingErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	_, err := d.WindowHandles()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_CommandWindowHandles_SingleResultCanBeReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": [
				"8"
			]
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	resp, err := d.WindowHandles()
	if err != nil || resp.State != "success" || resp.Handles[0] != "8" {
		t.Errorf(correctResponseErrorText)
	}
}

func Test_CommandWindowHandles_MultipleResultsCanBeReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": [
				"8",
				"9"
			]
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"
	resp, err := d.WindowHandles()
	if err != nil || resp.State != "success" || resp.Handles[0] != "8" || resp.Handles[1] != "9" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	SwitchToFrame Tests
*/
func Test_CommandSwitchToFrame_InvalidSessionIdResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.SwitchToFrame(nil)
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_CommandSwitchToFrame_InvalidByResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	byCSS, _ := ByCSSSelector("test")
	invalidBys := []By{
		nil,
		byCSS,
	}

	for _, i := range invalidBys {
		_, err := d.SwitchToFrame(i)
		if err == nil || !IsInvalidArgumentError(err) {
			t.Errorf(argumentErrorText)
		}
	}
}

func Test_CommandSwitchToFrame_APICommunicationErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	idx, _ := ByIndex(32)
	_, err := d.SwitchToFrame(idx)
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_CommandSwitchToFrame_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success"	
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	idx, _ := ByIndex(32)
	resp, err := d.SwitchToFrame(idx)
	if err != nil || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	SwitchToParentFrame tests
*/
func Test_CommandSwitchToParentFrame_InvalidSessionIDResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.SwitchToFrame(nil)
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_CommandSwitchToParentFrame_ApiCommunicationErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.SwitchToParentFrame()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_CommandSwitchToParentFrame_UnmarshallingErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.SwitchToParentFrame()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_CommandSwitchToParentFrame_CorrectResponseCanBeReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	resp, err := d.SwitchToParentFrame()
	if err != nil || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	WindowSize tests
*/
func Test_CommandWindowSize_InvalidSessionIDResultsInAnError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	_, err := d.WindowSize()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_CommandWindowSize_CommunicationErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.WindowSize()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_CommandWindowSize_UnmarshallingErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.WindowSize()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_CommandWindowSize_CorrectResultIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": {
				"width": 800,
				"height": 600
			}
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	resp, err := d.WindowSize()
	if err != nil || resp.State != "success" || resp.Dimensions.Width == 0 || resp.Dimensions.Height == 0 {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	SetWindowSize tests
*/

func Test_CommandSetWindowSize_NullDimensionResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)

	_, err := d.SetWindowSize(nil)
	if err == nil || !IsInvalidArgumentError(err) {
		t.Errorf(argumentErrorText)
	}
}

func Test_CommandSetWindowSize_InvalidSessionIDResultsInError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)

	dimensions := &Dimensions{
		Width:  830,
		Height: 255,
	}

	_, err := d.SetWindowSize(dimensions)
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_CommandSetWindowSize_CommunicationErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	dimensions := &Dimensions{
		Width:  830,
		Height: 255,
	}

	_, err := d.SetWindowSize(dimensions)
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_CommandSetWindowSize_UnmarshallingErrorIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	dimensions := &Dimensions{
		Width:  830,
		Height: 255,
	}

	_, err := d.SetWindowSize(dimensions)
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_CommandSetWindowSize_ResultIsReturnedSuccessfully(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	dimensions := &Dimensions{
		Width:  830,
		Height: 255,
	}

	resp, err := d.SetWindowSize(dimensions)
	if err != nil || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}
}
