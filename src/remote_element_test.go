package goselenium

import "testing"

func Test_RemoteElement_IDCanBeRetrieved(t *testing.T) {
	el := newSeleniumElement("test", nil)
	if el.ID() != "test" {
		t.Errorf(correctResponseErrorText)
	}
}
