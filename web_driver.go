package goselenium

// WebDriver is an interface which adheres to the W3C specification
// for WebDrivers (https://w3c.github.io/webdriver/webdriver-spec.html).
type WebDriver interface {
	/*
		PROPERTY ACCESS METHODS
	*/

	// DriverURL returns the URL where the W3C compliant web driver is hosted.
	DriverURL() string

	/*
		SESSION METHODS
	*/

	// CreateSession creates a session in the remote driver with the
	// desired capabilities.
	CreateSession() (*CreateSessionResponse, error)

	// DeleteSession deletes the current session associated with the web driver.
	DeleteSession() (*DeleteSessionResponse, error)

	// SessionStatus gets the status about whether a remove end is in a state
	// which it can create new sessions.
	SessionStatus() (*SessionStatusResponse, error)

	// SetSessionTimeout sets a timeout for one of the 3 options.
	// Call SessionScriptTimeout() to generate a script timeout.
	// Call SessionPageLoadTimeout() to generate a page load timeout.
	// Call SessionImplicitWaitTimeout() to generate an implicit wait timeout.
	SetSessionTimeout(to Timeout) (*SetSessionTimeoutResponse, error)

	/*
		NAVIGATION METHODS
	*/

	// Go forces the browser to perform a GET request on a URL.
	Go(url string) (*GoResponse, error)

	// CurrentURL returns the current URL of the top level browsing context.
	CurrentURL() (*CurrentURLResponse, error)

	// Back instructs the web driver to go one step back in the page history.
	Back() (*BackResponse, error)

	// Forward instructs the web driver to go one step forward in the page history.
	Forward() (*ForwardResponse, error)

	// Refresh instructs the web driver to refresh the page that it is currently on.
	Refresh() (*RefreshResponse, error)

	// Title gets the title of the current page of the web driver.
	Title() (*TitleResponse, error)

	/*
		COMMAND METHODS
	*/

	// WindowHandle retrieves the current active browsing string for the current session.
	WindowHandle() (*WindowHandleResponse, error)

	// CloseWindow closes the current active window (see WindowHandle() for what
	// window that will be).
	CloseWindow() (*CloseWindowResponse, error)

	// SwitchToWindow switches the current browsing context to a specified window
	// handle.
	SwitchToWindow(handle string) (*SwitchToWindowResponse, error)

	// WindowHandles gets all of the window handles for the current session.
	// To retrieve the currently active window handle, see WindowHandle().
	WindowHandles() (*WindowHandlesResponse, error)

	// SwitchToFrame switches to a frame determined by the "by" parameter.
	// You can use ByIndex to find the frame to switch to. Any other
	// By implementation will yield an InvalidByParameter error.
	SwitchToFrame(by By) (*SwitchToFrameResponse, error)

	// SwitchToParentFrame switches to the parent of the current top level
	// browsing context.
	SwitchToParentFrame() (*SwitchToParentFrameResponse, error)

	// WindowSize retrieves the current browser window size for the
	// active session.
	WindowSize() (*WindowSizeResponse, error)

	// SetWindowSize sets the current browser window size for the active
	// session.
	SetWindowSize(dimensions *Dimensions) (*SetWindowSizeResponse, error)

	// Maximize increases the current browser window to its maximum size.
	MaximizeWindow() (*MaximizeWindowResponse, error)

	/*
		ELEMENT METHODS
	*/

	// FindElement finds an element via a By implementation (i.e. CSS selector,
	// x-path). Attempting to find via index will result in an argument error
	// being thrown.
	FindElement(by By) (Element, error)

	// FindElements works the sameway as FindElement but can return more than
	// one result.
	FindElements(by By) ([]Element, error)
}

// Element is an interface which specifies what all WebDriver elements
// must do. Any requests which involve this element (i.e. FindElements that
// are children of the current element) will be specified. Anything not related
// will exist within the WebDriver interface/implementation.
type Element interface {
	// ID is the assigned ID for the element returned from the Selenium driver.
	ID() string

	// Selected gets whether or not the current element is selected. This only
	// makes sense for inputs such as radio buttons and checkboxes.
	Selected() (*ElementSelectedResponse, error)

	// Attribute retrieves an attribute (i.e. href, class) of the current
	// active element.
	Attribute(att string) (*ElementAttributeResponse, error)

	// CSSValue retrieves a CSS property associated with the current element. As an example, this could be the 'background' or 'font-family' properties.
	CSSValue(prop string) (*ElementCSSValueResponse, error)

	// Text gets the value of element.innerText for the current element.
	Text() (*ElementTextResponse, error)

	// TagName gets the HTML element name (i.e. p, div) of the currently selected
	// element.
	TagName() (*ElementTagNameResponse, error)

	// Rectangle gets the dimensions and co-ordinates of the currently selected
	// element.
	Rectangle() (*ElementRectangleResponse, error)
}

// Timeout is an interface which specifies what all timeout requests must follow.
type Timeout interface {
	// Type is the type of the timeout that is being used.
	Type() string

	// Timeout is the timeout in milliseconds.
	Timeout() int
}

// By is an interface that defines what all 'ByX' methods must return.
type By interface {
	// Type is the type of by (i.e. id, xpath, class, name, index).
	Type() string

	// Value is the actual value to retrieve something by (i.e. #test, 1).
	Value() interface{}
}
