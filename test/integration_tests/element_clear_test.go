package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_ElementClear_ElementsAreClearedSuccessfully(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
	}{
		{
			url: "https://heraclmene.github.io/helpers/goselenium/clear.html",
			by:  goselenium.ByCSSSelector("input#should-clear"),
		},
	}
	for _, te := range tests {
		driver := createDriver(t)
		_, err := driver.CreateSession()
		if err != nil {
			errorAndWrap(t, "Create session failed", err)
		}

		_, err = driver.Go(te.url)
		if err != nil {
			errorAndWrap(t, "Navigating to URL failed", err)
		}

		el, err := driver.FindElement(te.by)
		if err != nil {
			errorAndWrap(t, "Retrieving element failed", err)
		}

		_, err = el.Clear()
		if err != nil {
			errorAndWrap(t, "Clearing element failed", err)
		}

		resp, err := el.Text()
		if err != nil || len(resp.Text) > 0 {
			errorAndWrap(t, "Retrieving text failed or text was not cleared", err)
		}

		driver.DeleteSession()

		printObjectResult(resp)
	}
}
