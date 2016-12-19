package integrationtests

import (
	"testing"
	"time"

	"github.com/bunsenapp/go-selenium"
)

func Test_ElementClick_ClickSuccessfullyNavigates(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url      string
		by       goselenium.By
		expTitle string
	}{
		{
			url:      "https://www.google.com",
			by:       goselenium.ByLinkText("Privacy"),
			expTitle: "Privacy Policy – Privacy & Terms – Google",
		},
		{
			url:      "https://news.ycombinator.com",
			by:       goselenium.ByLinkText("new"),
			expTitle: "New Links | Hacker News",
		},
	}
	for _, te := range tests {
		driver := createDriver(t)
		_, err := driver.CreateSession()
		if err != nil {
			errorAndWrap(t, "Error creating session.", err)
		}

		_, err = driver.Go(te.url)
		if err != nil {
			errorAndWrap(t, "Error visiting URL.", err)
		}

		el, err := driver.FindElement(te.by)
		if err != nil {
			errorAndWrap(t, "Error finding element", err)
		}

		_, err = el.Click()
		if err != nil {
			errorAndWrap(t, "Error clicking element", err)
		}

		// TODO: Unfortunate Selenium flaw - will need the helpers to get around
		// an explicit wait.
		time.Sleep(1 * time.Second)

		resp, err := driver.Title()
		if err != nil || resp.Title != te.expTitle {
			errorAndWrap(t, "Error retrieving title or title was not correct", err)
		}

		driver.DeleteSession()

		printObjectResult(resp)
	}
}
