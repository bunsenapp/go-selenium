package main

import (
	"fmt"

	goselenium "github.com/bunsenapp/go-selenium"
)

func main() {
	// Create the capabilities.
	capabilities := goselenium.Capabilities{}
	capabilities.SetBrowser(goselenium.FirefoxBrowser())

	// Create the driver.
	driver, err := goselenium.NewSeleniumWebDriver("http://localhost:4444/wd/hub/", capabilities)
	if err != nil {
		fmt.Println("Error creating web driver.")
		return
	}

	// Create the session.
	_, err = driver.CreateSession()
	if err != nil {
		fmt.Println("Error creating session.")
		return
	}

	// Navigate to Google.
	_, err = driver.Go("https://www.google.com")
	if err != nil {
		fmt.Println("An error occurred whilst visiting URL.")
		return
	}

	// Find a non existent element for it to error.
	_, err = driver.FindElement(goselenium.ByCSSSelector("mynonexistentelement"))
	if err != nil {
		// Switch the different types of errors. You do not need to do this in
		// every call and can simply abstract it behind a function. If you
		// don't want to handle the custom errors, they all implement the
		// Error interface meaning it'll work anywhere your normal errors do.
		switch err.(type) {
		case goselenium.CommunicationError:
			e := err.(goselenium.CommunicationError)
			// Switch the different states that we want to handle.
			switch e.Response.State {
			case goselenium.UnknownError:
				fmt.Println("An unknown error occurred.")
			case goselenium.SessionNotCreated:
				fmt.Println("The session was not created.")
			case goselenium.NoSuchElement:
				fmt.Println("Failed to find element. Example passed!")
			}
		case goselenium.UnmarshallingError:
			fmt.Println("An unmarshalling error occurred :<")
		}
	}

	// Delete the session.
	driver.DeleteSession()
}
