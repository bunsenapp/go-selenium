package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium/src"
)

func Test_ElementText_CorrectElementTextGetsReturned(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
		exp string
	}{
		{
			url: "https://google.com",
			by:  goselenium.ByCSSSelector("input#lst-ib"),
			exp: "",
		},
		{
			url: "https://news.ycombinator.com",
			by:  goselenium.ByCSSSelector("a[href='submit']"),
			exp: "submit",
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

		txt, err := el.Text()
		if err != nil || txt.Text != te.exp {
			errorAndWrap(t, "Error whilst getting element text or text was not correct", err)
		}

		driver.DeleteSession()

		printObjectResult(txt)
	}
}
