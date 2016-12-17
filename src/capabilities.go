package goselenium

import "encoding/json"

// Browser represents a browser to run within Selenium.
type Browser struct {
	browserName string
}

// BrowserName returns the browser name assigned to the current browser object.
func (b *Browser) BrowserName() string {
	return b.browserName
}

// FirefoxBrowser returns a Firefox browser object.
func FirefoxBrowser() Browser {
	return Browser{"firefox"}
}

// ChromeBrowser returns a Chrome browser object.
func ChromeBrowser() Browser {
	return Browser{"chrome"}
}

// Capabilities represents the capabilities defined in the W3C specification.
// The main capability is the browser, which can be set by calling one of the
// \wBrowser\(\) methods.
type Capabilities struct {
	browser *Browser
}

// Browser yields the browser capability assigned to the current Capabilities
// object..
func (c *Capabilities) Browser() Browser {
	if c.browser != nil {
		return *c.browser
	}

	return Browser{}
}

// SetBrowser sets the browser capability to be one of the allowed browsers.
func (c *Capabilities) SetBrowser(b Browser) {
	c.browser = &b
}

func (c *Capabilities) toJSON() (string, error) {
	capabilities := map[string]map[string]interface{}{
		"desiredCapabilities": map[string]interface{}{
			"browserName": c.browser.browserName,
		},
	}

	capabilitiesJSON, err := json.Marshal(capabilities)
	if err != nil {
		return "", err
	}

	return string(capabilitiesJSON), nil
}
