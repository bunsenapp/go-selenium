package integrationtests

import "testing"

func Test_CookieCookie_CanRetrieveCookieFromWebPage(t *testing.T) {
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

	resp, err := driver.Cookie("__cfduid")
	if err != nil || resp.State != "success" || resp.Cookie.Name == "" {
		errorAndWrap(t, "Error was thrown whilst retrieving cookie", err)
	}

	printObjectResult(resp)
}
