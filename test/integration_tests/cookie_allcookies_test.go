package integrationtests

import "testing"

func Test_CookieAllCookies_CanRetrieveAllCookiesFromWebPage(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	_, err = driver.Go("https://news.ycombinator.com")
	if err != nil {
		errorAndWrap(t, "Error was thrown whilst navigating or result was not a success.", err)
	}

	resp, err := driver.AllCookies()
	if err != nil || resp.State != "success" || resp.Cookies[0].Name == "" {
		errorAndWrap(t, "Error was thrown whilst retrieving cookies", err)
	}

	printObjectResult(resp)
}
