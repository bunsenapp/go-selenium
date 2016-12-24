package integrationtests

import (
	"testing"

	"github.com/bunsenapp/go-selenium"
)

func Test_CookieAddCookie_CanAddCookieWithCorrectFields(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst trying to create session.", err)
	}

	_, err = driver.Go("https://news.ycombinator.com")
	if err != nil {
		errorAndWrap(t, "Error was thrown whilst navigating or result was not a success.", err)
	}

	resp, err := driver.AddCookie(&goselenium.Cookie{
		Name:       "cookie",
		Value:      "cookieValue",
		Path:       "/",
		Domain:     ".ycombinator.com",
		SecureOnly: false,
		HTTPOnly:   true,
	})
	if err != nil || resp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst retrieving cookies", err)
	}

	printObjectResult(resp)

}
