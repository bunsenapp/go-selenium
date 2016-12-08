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

// FirefoxBrowser returns a browser object for a Firefox browser.
func FirefoxBrowser() Browser {
	return Browser{"firefox"}
}

// ChromeBrowser returns a browser object for a Chrome browser.
func ChromeBrowser() Browser {
	return Browser{"chrome"}
}

// Capabilities represents the capabilities defined in the W3C specification.
type Capabilities struct {
	browser *Browser
}

// Browser yields the browser capability.
func (c *Capabilities) Browser() Browser {
	if c.browser != nil {
		return *c.browser
	}

	return Browser{}
}

// SetBrowser sets the browser capability.
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
