package goselenium

import "encoding/json"

// Browser defines a supported selenium enabled browser.
type Browser interface {
	BrowserName() string
}

// Browser represents a browser to run within Selenium.
type browser struct {
	browserName string
}

// BrowserName returns the browser name assigned to the current browser object.
func (b browser) BrowserName() string {
	return b.browserName
}

// FirefoxBrowser returns a Firefox browser object.
func FirefoxBrowser() Browser {
	return browser{"firefox"}
}

// ChromeBrowser returns a Chrome browser object.
func ChromeBrowser() Browser {
	return browser{"chrome"}
}

// AndroidBrowser returns an Android browser object.
func AndroidBrowser() Browser {
	return browser{"android"}
}

// HTMLUnitBrowser returns a HTMLUnit browser object.
func HTMLUnitBrowser() Browser {
	return browser{"htmlunit"}
}

// InternetExplorerBrowser returns an IE browser object.
func InternetExplorerBrowser() Browser {
	return browser{"internetexplorer"}
}

// IPhoneBrowser returns an IPhone browser object.
func IPhoneBrowser() Browser {
	return browser{"iphone"}
}

// IPadBrowser returns an IPad browser object.
func IPadBrowser() Browser {
	return browser{"ipad"}
}

// OperaBrowser returns an Opera browser object.
func OperaBrowser() Browser {
	return browser{"opera"}
}

// SafariBrowser returns a Safari browser object.
func SafariBrowser() Browser {
	return browser{"safari"}
}

// Capabilities represents the capabilities defined in the W3C specification.
// The main capability is the browser, which can be set by calling one of the
// \wBrowser\(\) methods.
type Capabilities struct {
	browser Browser
}

// Browser yields the browser capability assigned to the current Capabilities
// object..
func (c *Capabilities) Browser() Browser {
	if c.browser != nil {
		return c.browser
	}

	return browser{}
}

// SetBrowser sets the browser capability to be one of the allowed browsers.
func (c *Capabilities) SetBrowser(b Browser) {
	c.browser = b
}

func (c *Capabilities) toJSON() (string, error) {
	capabilities := map[string]map[string]interface{}{
		"desiredCapabilities": {
			"browserName": c.browser.BrowserName(),
		},
	}

	capabilitiesJSON, err := json.Marshal(capabilities)
	if err != nil {
		return "", err
	}

	return string(capabilitiesJSON), nil
}
