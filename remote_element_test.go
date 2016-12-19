package goselenium

import (
	"errors"
	"testing"
)

func Test_RemoteElement_IDCanBeRetrieved(t *testing.T) {
	el := newSeleniumElement("test", nil)
	if el.ID() != "test" {
		t.Errorf(correctResponseErrorText)
	}
}

/* SELECTED TESTS
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
		t.Errorf(correctResponseErrorText)
	}
}

/*
	ATTRIBUTE TESTS
*/
func Test_ElementAttribute_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.Attribute("test")
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_ElementAttribute_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.Attribute("test")
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_ElementAttribute_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": "test value"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	resp, err := el.Attribute("test")
	if err != nil || resp.State != "success" || resp.Value != "test value" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	CSSVALUE TESTS
*/
func Test_ElementCSSValue_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.CSSValue("test")
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_ElementCSSValue_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.CSSValue("test")
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_ElementCSSValue_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": "test value"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	resp, err := el.CSSValue("test")
	if err != nil || resp.State != "success" || resp.Value != "test value" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	TEXT TESTS
*/
func Test_ElementText_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.Text()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_ElementText_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.Text()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_ElementText_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": "test value"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	resp, err := el.Text()
	if err != nil || resp.State != "success" || resp.Text != "test value" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	TAG NAME TESTS
*/
func Test_ElementTagName_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.TagName()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_ElementTagName_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.TagName()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_ElementTagName_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": "test value"
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	resp, err := el.TagName()
	if err != nil || resp.State != "success" || resp.Tag != "test value" {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	RECTANGLE TESTS
*/
func Test_ElementRectangle_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.Rectangle()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_ElementRectangle_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.Rectangle()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_ElementRectangle_CorrectResponseIsReturned(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn: `{
			"state": "success",
			"value": {
				"x": 100,
				"y": 200,
				"width": 50,
				"height": 50
			}
		}`,
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	resp, err := el.Rectangle()
	if err != nil || resp.State != "success" || resp.Rectangle.X != 100 || resp.Rectangle.Height != 50 {
		t.Errorf(correctResponseErrorText)
	}
}

/*
	ENABLED TESTS
*/
func Test_ElementEnabled_CommunicationErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "",
		errorToReturn: errors.New("An error :<"),
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.Enabled()
	if err == nil || !IsCommunicationError(err) {
		t.Errorf(apiCommunicationErrorText)
	}
}

func Test_ElementEnabled_UnmarshallingErrorIsReturnedCorrectly(t *testing.T) {
	api := &testableAPIService{
		jsonToReturn:  "Invalid JSON!",
		errorToReturn: nil,
	}

	d := setUpDriver(setUpDefaultCaps(), api)
	d.sessionID = "12345"

	el := newSeleniumElement("0", d)
	_, err := el.Enabled()
	if err == nil || !IsUnmarshallingError(err) {
		t.Errorf(unmarshallingErrorText)
	}
}

func Test_ElementEnabled_CorrectResponseIsReturned(t *testing.T) {
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
	resp, err := el.Enabled()
	if err != nil || resp.State != "success" || !resp.Enabled {
		t.Errorf(correctResponseErrorText)
	}
}
