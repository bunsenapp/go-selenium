package main

import (
	"fmt"
	"time"

	goselenium "github.com/bunsenapp/go-selenium"
)

func main() {
	// Create capabilities, driver etc.
	capabilities := goselenium.Capabilities{}
	capabilities.SetBrowser(goselenium.FirefoxBrowser())

	driver, err := goselenium.NewSeleniumWebDriver("http://localhost:4444/wd/hub", capabilities)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = driver.CreateSession()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Delete the session once this function is completed.
	defer driver.DeleteSession()

	// Navigate to the HackerNews website.
	_, err = driver.Go("https://news.ycombinator.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Click the 'new' link at the top
	el, err := driver.FindElement(goselenium.ByCSSSelector("a[href='newest']"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Click the link.
	_, err = el.Click()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Wait until the URL has changed with a timeout of 1 second and a check
	// interval of 10ms..
	newLink := "https://news.ycombinator.com/newest"
	ok := driver.Wait(goselenium.UntilURLIs(newLink), 1*time.Second, 10*time.Millisecond)
	if !ok {
		fmt.Println("Wait timed out :<")
		return
	}

	// Woohoo! We have successfully navigated to a page.
	fmt.Println("Successfully navigated to URL " + newLink)
}
