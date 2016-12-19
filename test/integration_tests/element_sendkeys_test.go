package integrationtests

import (
	"strings"
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_ElementSendKeys_CanSendKeysToInputField(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error creating session", err)
	}

	_, err = driver.Go("https://www.google.com")
	if err != nil {
		errorAndWrap(t, "Error navigating to URL", err)
	}

	el, err := driver.FindElement(goselenium.ByCSSSelector("input#lst-ib"))
	if err != nil {
		errorAndWrap(t, "Error finding element", err)
	}

	_, err = el.SendKeys("test")
	if err != nil {
		errorAndWrap(t, "Error sending keys to element", err)
	}

	_, err = el.SendKeys(goselenium.EnterKey)
	if err != nil {
		errorAndWrap(t, "Error sending enter key to element", err)
	}

	url, err := driver.CurrentURL()
	if err != nil || !strings.Contains(url.URL, "test") {
		errorAndWrap(t, "Error retrieving current URL or it did not contain the correct value", err)
	}

	printObjectResult(url)
}
