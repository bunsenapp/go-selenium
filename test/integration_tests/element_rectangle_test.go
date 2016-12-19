package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_ElementRectangle_SizeIsReturnedCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	_, err = driver.Go("https://heraclmene.github.io/helpers/goselenium/size.html")
	if err != nil {
		errorAndWrap(t, "Error thrown whilst visiting url.", err)
	}

	el, err := driver.FindElement(goselenium.ByCSSSelector("div"))
	if err != nil || el == nil {
		errorAndWrap(t, "Error whilst finding element or element was not found", err)
	}

	resp, err := el.Rectangle()
	if err != nil || resp.Rectangle.Width != 100 || resp.Rectangle.Height != 100 {
		errorAndWrap(t, "Error was returned or element's size was not correct.", err)
	}

	printObjectResult(resp)
}

func Test_ElementRectangle_PositionIsReturnedCorrectly(t *testing.T) {
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

	resp, err := el.Rectangle()
	if err != nil || resp.Rectangle.X == 0 || resp.Rectangle.Height == 0 {
		errorAndWrap(t, "Error was returned or element's size was not correct.", err)
	}

	printObjectResult(resp)
}
