package integrationtests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bunsenapp/go-selenium/src"
)

func Test_ElementCSSValue_CanGetCorrectCSSValue(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url     string
		by      goselenium.By
		cssName string
		cssVal  string
	}{
		{
			url:     "https://www.google.com",
			by:      goselenium.ByCSSSelector("input[name='btnK']"),
			cssName: "font-family",
			cssVal:  "arial,sans-serif",
		},
		{
			url:     "https://news.ycombinator.com",
			by:      goselenium.ByCSSSelector("a[href='news']"),
			cssName: "color",
			cssVal:  "rgb(0, 0, 0)",
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

		att, err := el.CSSValue(te.cssName)
		if err != nil || !strings.Contains(att.Value, te.cssVal) {
			fmt.Println(att)
			errorAndWrap(t, "Error whilst retrieving CSS class or value was not correct", err)
		}

		driver.DeleteSession()

		printObjectResult(att)
	}
}

func Test_ElementCSSValue_CSSValueThatDoesNotExistDoesNotCauseAnError(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url     string
		by      goselenium.By
		cssName string
	}{
		{
			url:     "https://www.google.com",
			by:      goselenium.ByCSSSelector("input[name='btnK']"),
			cssName: "background",
		},
		{
			url:     "https://news.ycombinator.com",
			by:      goselenium.ByCSSSelector("a[href='news']"),
			cssName: "border-radius",
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

		att, err := el.CSSValue(te.cssName)
		if err != nil {
			fmt.Println(att)
			errorAndWrap(t, "Error whilst retrieving CSS class", err)
		}

		driver.DeleteSession()

		printObjectResult(att)
	}
}
