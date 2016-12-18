package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium/src"
)

func Test_ElementAttribute_CanRetrieveAttributeCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url  string
		by   goselenium.By
		attr string
		val  string
	}{
		{
			url:  "https://www.google.com",
			by:   goselenium.ByCSSSelector("input#lst-ib"),
			attr: "maxlength",
			val:  "2048",
		},
	}
	for _, te := range tests {
		driver := createDriver(t)
		_, err := driver.CreateSession()
		if err != nil {
			errorAndWrap(t, "Error thrown whilst creating session.", err)
		}

		_, err = driver.Go(te.url)
		if err != nil {
			errorAndWrap(t, "Error thrown whilst visiting url.", err)
		}

		el, err := driver.FindElement(te.by)
		if err != nil || el == nil {
			errorAndWrap(t, "Error whilst finding element or element was not found", err)
		}

		att, err := el.Attribute(te.attr)
		if err != nil || att.Value != te.val {
			errorAndWrap(t, "Error whilst retrieving attribute or attribute value was not correct", err)
		}

		driver.DeleteSession()

		printObjectResult(att)
	}
}
