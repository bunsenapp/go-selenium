package goselenium

import (
	"errors"
	"testing"
)

func Test_ScreenshotScreenshot_NoSessionIdCausesError(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)

	_, err := d.Screenshot()
	if err == nil || !IsSessionIDError(err) {
		t.Errorf(sessionIDErrorText)
	}
}

func Test_ScreenshotScreenshot_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error! :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.Screenshot()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_ScreenshotScreenshot_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	_, err := d.Screenshot()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_ScreenshotScreen_CorrectResponseCanBeReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": "dGVzdA=="
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	resp, err := d.Screenshot()
	if err != nil || resp.State != "success" || resp.EncodedImage != "dGVzdA==" {
		t.Errorf(correctResponseErrorText)
	}
}

func Test_ScreenshotScreenshot_Base64StringCanBeDecoded(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": "dGVzdA=="
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	resp, err := d.Screenshot()
	if err != nil || resp.State != "success" {
		t.Errorf(correctResponseErrorText)
	}

	bytes, err := resp.ImageBytes()
	if err != nil || len(bytes) == 0 {
		t.Errorf(correctResponseErrorText)
	}
}
