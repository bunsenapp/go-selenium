package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_ElementFindElement_CanFindElementByCSSSelector(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
	}{
		{"https://www.google.com", goselenium.ByCSSSelector("input#lst-ib")},
		{"https://www.reddit.com", goselenium.ByCSSSelector("input[name=q]")},
		{"https://news.ycombinator.com", goselenium.ByCSSSelector("#hnmain")},
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

		printObjectResult(el)

		driver.DeleteSession()
	}
}

func Test_ElementFindElement_CanFindElementByLinkText(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
	}{
		{"https://www.google.com", goselenium.ByLinkText("Gmail")},
		{"https://www.reddit.com", goselenium.ByLinkText("new")},
		{"https://news.ycombinator.com", goselenium.ByLinkText("submit")},
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

		printObjectResult(el)

		driver.DeleteSession()
	}
}

func Test_ElementFindElement_CanFindElementByPartialLinkText(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
	}{
		{"https://www.google.com", goselenium.ByPartialLinkText("Gmai")},
		{"https://www.reddit.com", goselenium.ByPartialLinkText("ew")},
		{"https://news.ycombinator.com", goselenium.ByPartialLinkText("ubmit")},
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

		printObjectResult(el)

		driver.DeleteSession()
	}
}

func Test_ElementFindElement_CanFindElementByXPath(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url string
		by  goselenium.By
	}{
		{"https://www.google.com", goselenium.ByXPath("//input[@id='lst-ib']")},
		{"https://news.ycombinator.com", goselenium.ByXPath("//b[@class='hnname']/a[@href='news']")},
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

		printObjectResult(el)

		driver.DeleteSession()
	}
}
