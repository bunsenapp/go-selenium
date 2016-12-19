package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_ElementTagName_CanRetrieveCorrectTagName(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
		exp string
	}{
		{
			url: "https://news.ycombinator.com",
			by:  goselenium.ByCSSSelector("a[href='submit']"),
			exp: "a",
		},
		{
			url: "https://google.com",
			by:  goselenium.ByCSSSelector("input#lst-ib"),
			exp: "input",
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

		resp, err := el.TagName()
		if err != nil || resp.Tag != te.exp {
			errorAndWrap(t, "Error was returned or tag name was incorrect.", err)
		}

		driver.DeleteSession()

		printObjectResult(resp)
	}
}
