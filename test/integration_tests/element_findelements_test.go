package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_ElementFindElements_CanFindElementsByCSSSelector(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
	}{
		{"https://www.google.com", goselenium.ByCSSSelector("input")},
		{"https://www.reddit.com", goselenium.ByCSSSelector("a")},
		{"https://news.ycombinator.com", goselenium.ByCSSSelector(".storylink")},
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

		el, err := driver.FindElements(te.by)
		if err != nil || el == nil || len(el) <= 1 {
			errorAndWrap(t, "Error whilst finding elements or element was not found", err)
		}

		printObjectResult(el)

		driver.DeleteSession()
	}
}

func Test_ElementFindElements_CanFindElementsByLinkText(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
	}{
		{"https://www.reddit.com", goselenium.ByLinkText("share")},
		{"https://news.ycombinator.com", goselenium.ByLinkText("hide")},
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

		el, err := driver.FindElements(te.by)
		if err != nil || el == nil || len(el) <= 1 {
			errorAndWrap(t, "Error whilst finding elements or elements were not found", err)
		}

		printObjectResult(el)

		driver.DeleteSession()
	}
}

func Test_ElementFindElements_CanFindElementsByPartialLinkText(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
	}{
		{"https://www.google.com", goselenium.ByPartialLinkText("a")},
		{"https://www.reddit.com", goselenium.ByPartialLinkText("hi")},
		{"https://news.ycombinator.com", goselenium.ByPartialLinkText("comment")},
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

		el, err := driver.FindElements(te.by)
		if err != nil || el == nil || len(el) <= 1 {
			errorAndWrap(t, "Error whilst finding elements or elements were not found", err)
		}

		printObjectResult(el)

		driver.DeleteSession()
	}
}

func Test_ElementFindElements_CanFindElementsByXPath(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
	}{
		{"https://www.google.com", goselenium.ByXPath("//input")},
		{"https://news.ycombinator.com", goselenium.ByXPath("//a")},
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

		el, err := driver.FindElements(te.by)
		if err != nil || el == nil {
			errorAndWrap(t, "Error whilst finding elements or element were not found", err)
		}

		printObjectResult(el)

		driver.DeleteSession()
	}
}
