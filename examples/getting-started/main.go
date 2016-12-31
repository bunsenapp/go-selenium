package main

import (
	"fmt"

	"github.com/bunsenapp/go-selenium"
)

func main() {
	// Create a capabilities object.
	capabilities := goselenium.Capabilities{}

	// Populate it with the browser you wish to use.
	capabilities.SetBrowser(goselenium.FirefoxBrowser())

	// Initialise a new web driver.
	driver, err := goselenium.NewSeleniumWebDriver("http://localhost:4444/wd/hub", capabilities)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a session.
	_, err = driver.CreateSession()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Defer the deletion of the session.
	defer driver.DeleteSession()

	// Navigate to Google.
	_, err = driver.Go("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}

	// Hooray, we navigated to Google!
	fmt.Println("Successfully navigated to Google!")
}
