package integrationtests

import "testing"

func Test_CookieDeleteCookie_CanDeleteSpecifiedCookie(t *testing.T) {
	setUp()
	defer tearDown()

	tests := []struct {
		url        string
		cookieName string
	}{
		{
			url:        "https://news.ycombinator.com",
			cookieName: "__cfduid",
		},
		{
			url:        "https://www.google.com",
			cookieName: "CONSENT",
		},
	}
	for _, te := range tests {
		driver := createDriver(t)
		_, err := driver.CreateSession()
		if err != nil {
			t.Errorf("Session creation failed")
		}

		_, err = driver.Go(te.url)
		if err != nil {
			t.Errorf("Navigation failed")
		}

		resp, err := driver.DeleteCookie(te.cookieName)
		if err != nil || resp.State != "success" {
			t.Errorf("Error whilst deleting cookie or was not a success.")
		}

		cookie, err := driver.Cookie(te.cookieName)
		if err != nil || cookie.Cookie.Name == te.cookieName {
			t.Errorf("Cookie still exists or an error occurred.")
		}

		driver.DeleteSession()

		printObjectResult(resp)
	}
}

func Test_CookieDeleteCookie_CanDeleteAllCookies(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		t.Errorf("Session creation failed")
	}

	_, err = driver.Go("https://www.google.com")
	if err != nil {
		t.Errorf("Navigation failed")
	}

	resp, err := driver.DeleteCookie("")
	if err != nil || resp.State != "success" {
		t.Errorf("Error whilst deleting cookie or was not a success.")
	}

	cookies, err := driver.AllCookies()
	if err != nil || len(cookies.Cookies) != 0 {
		t.Errorf("Cookies still exist or error was returned.")
	}

	printObjectResult(resp)
}
