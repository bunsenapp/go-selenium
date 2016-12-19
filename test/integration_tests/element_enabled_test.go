package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_ElementEnabled_EnabledElementIsReturnedCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	_, err = driver.Go("https://google.com")
	if err != nil {
		errorAndWrap(t, "Error thrown whilst visiting url.", err)
	}

	el, err := driver.FindElement(goselenium.ByCSSSelector("input#lst-ib"))
	if err != nil || el == nil {
		errorAndWrap(t, "Error whilst finding element or element was not found", err)
	}

	resp, err := el.Enabled()
	if err != nil || !resp.Enabled {
		errorAndWrap(t, "Error whilst retrieving response or element was not enabled", err)
	}

	printObjectResult(resp)
}

func Test_ElementEnabled_DisabledElementIsReturnedCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
		exp bool
	}{
		{
			url: "https://news.ycombinator.com",
			by:  goselenium.ByCSSSelector("a[href='news']"),
			exp: true,
		},
		{
			url: "https://heraclmene.github.io/helpers/goselenium/disabled.html",
			by:  goselenium.ByCSSSelector("input"),
			exp: false,
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

		resp, err := el.Enabled()
		if err != nil || resp.Enabled != te.exp {
			errorAndWrap(t, "Error whilst retrieving response or the element's enabled property was not correct", err)
		}

		driver.DeleteSession()

		printObjectResult(resp)
	}
}
