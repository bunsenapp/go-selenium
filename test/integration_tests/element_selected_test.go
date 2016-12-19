package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_ElementSelected_CheckedElementReturnsCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []goselenium.By{
		goselenium.ByCSSSelector("input[name='i_am_checked']"),
		goselenium.ByCSSSelector("input[name='i_am_selected']"),
	}
	for _, te := range tests {
		driver := createDriver(t)
		_, err := driver.CreateSession()
		if err != nil {
			errorAndWrap(t, "Error thrown whilst creating session.", err)
		}

		_, err = driver.Go("https://heraclmene.github.io/helpers/goselenium/selected.html")
		if err != nil {
			errorAndWrap(t, "Error thrown whilst visiting url.", err)
		}

		el, err := driver.FindElement(te)
		if err != nil || el == nil {
			errorAndWrap(t, "Error whilst finding element or element was not found", err)
		}

		resp, err := el.Selected()
		if err != nil || !resp.Selected {
			errorAndWrap(t, "Error was returned or element was not selected.", err)
		}

		driver.DeleteSession()

		printObjectResult(resp)
	}
}

func Test_ElementSelected_UncheckedElementReturnsCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []goselenium.By{
		goselenium.ByCSSSelector("input[name='i_am_not_checked']"),
		goselenium.ByCSSSelector("input[name='i_am_not_selected']"),
	}
	for _, te := range tests {
		driver := createDriver(t)
		_, err := driver.CreateSession()
		if err != nil {
			errorAndWrap(t, "Error thrown whilst creating session.", err)
		}

		_, err = driver.Go("https://heraclmene.github.io/helpers/goselenium/selected.html")
		if err != nil {
			errorAndWrap(t, "Error thrown whilst visiting url.", err)
		}

		el, err := driver.FindElement(te)
		if err != nil || el == nil {
			errorAndWrap(t, "Error whilst finding element or element was not found", err)
		}

		resp, err := el.Selected()
		if err != nil || resp.Selected {
			errorAndWrap(t, "Error was returned or element was selected.", err)
		}

		driver.DeleteSession()

		printObjectResult(resp)
	}
}

func Test_ElementSelected_RandomElementsDoNotError(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []goselenium.By{
		goselenium.ByCSSSelector("a.storylink"),
		goselenium.ByCSSSelector(".title"),
	}
	for _, te := range tests {
		driver := createDriver(t)
		_, err := driver.CreateSession()
		if err != nil {
			errorAndWrap(t, "Error thrown whilst creating session.", err)
		}

		_, err = driver.Go("https://news.ycombinator.com")
		if err != nil {
			errorAndWrap(t, "Error thrown whilst visiting url.", err)
		}

		el, err := driver.FindElement(te)
		if err != nil || el == nil {
			errorAndWrap(t, "Error whilst finding element or element was not found", err)
		}

		resp, err := el.Selected()
		if err != nil || resp.Selected {
			errorAndWrap(t, "Error was returned or element was selected.", err)
		}

		driver.DeleteSession()

		printObjectResult(resp)
	}
}
