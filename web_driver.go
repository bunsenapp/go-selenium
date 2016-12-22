package goselenium

const (
	UnidentifiedKey   = string('\uE000')
	CancelKey         = string('\uE001')
	HelpKey           = string('\uE002')
	BackspaceKey      = string('\uE003')
	TabKey            = string('\uE004')
	ClearKey          = string('\uE005')
	ReturnKey         = string('\uE006')
	EnterKey          = string('\uE007')
	ShiftKey          = string('\uE008')
	ControlKey        = string('\uE009')
	AltKey            = string('\uE00A')
	PauseKey          = string('\uE00B')
	EscapeKey         = string('\uE00C')
	SpaceKey          = string('\uE00D')
	PageUpKey         = string('\uE00E')
	PageDownKey       = string('\uE00F')
	EndKey            = string('\uE010')
	HomeKey           = string('\uE011')
	ArrowLeftKey      = string('\uE012')
	ArrowUpKey        = string('\uE013')
	ArrowRightKey     = string('\uE014')
	ArrowDownKey      = string('\uE015')
	InsertKey         = string('\uE016')
	DeleteKey         = string('\uE017')
	SemiColonKey      = string('\uE018')
	EqualsKey         = string('\uE019')
	AsteriskKey       = string('\uE024')
	PlusKey           = string('\uE025')
	CommaKey          = string('\uE026')
	MinusKey          = string('\uE027')
	PeriodKey         = string('\uE028')
	ForwardSlashKey   = string('\uE029')
	F1Key             = string('\uE031')
	F2Key             = string('\uE032')
	F3Key             = string('\uE033')
	F4Key             = string('\uE034')
	F5Key             = string('\uE035')
	F6Key             = string('\uE036')
	F7Key             = string('\uE037')
	F8Key             = string('\uE038')
	F9Key             = string('\uE039')
	F10Key            = string('\uE03A')
	F11Key            = string('\uE03B')
	F12Key            = string('\uE03C')
	MetaKey           = string('\uE03D')
	ZenkakuHankakuKey = string('\uE040')
)

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

	// FindElements works the same way as FindElement but can return more than
	// one result.
	FindElements(by By) ([]Element, error)

	/*
		DOCUMENT HANDLING METHODS
	*/

	// PageSource retrieves the outerHTML value of the current URL.
	PageSource() (*PageSourceResponse, error)

	// ExecuteScript executes a Javascript script on the currently active
	// page.
	ExecuteScript(script string) (*ExecuteScriptResponse, error)

	// ExecuteScriptAsync executes a Javascript script asynchronously on the
	// currently active page. If you do not have experience with this call,
	// there is an example below.
	//
	// The async handler runs on the concept of a callback; meaning it will run
	// your code asynchronously and if it completes, will call the callback.
	//
	// Selenium helpfully provides a callback function which is passed in
	// to the 'arguments' array that you can access within your script. The
	// callback function is always the LAST element of the array. You can
	// access it like the below:
	//		var callback = arguments[arguments.length - 1];
	// The callback function also accepts one argument as a parameter, this
	// can be anything and will be assigned to the Response property of
	// ExecuteScriptResponse.
	//
	// An example:
	//		var callback = arguments[arguments.length - 1];
	//		doLongWindedTask();
	//		callback();
	ExecuteScriptAsync(script string) (*ExecuteScriptResponse, error)

	/*
		COOKIE METHODS
	*/

	// AllCookies returns all cookies associated with the active URL of the
	// current browsing context.
	AllCookies() (*AllCookiesResponse, error)

	// Cookie gets a single named cookie associated with the active URL of the
	// current browsing context.
	Cookie(name string) (*CookieResponse, error)
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

	// Enabled gets whether or not the current selected elemented is enabled.
	Enabled() (*ElementEnabledResponse, error)

	// Click clicks the currently selected element. Please note, you may have to
	// implement your own wait to ensure the page actually navigates. This is due to
	// Selenium having no idea whether or not your click will be interrupted by JS.
	// Alternatively, you can use the WaitUntil(TitleEquals("title"), 20) to
	// automatically wait until the page title has changed.
	Click() (*ElementClickResponse, error)

	// Clear clears the currently selected element according to the specification.
	Clear() (*ElementClearResponse, error)

	// SendKeys sends a set of keystrokes to the currently selected element.
	SendKeys(keys string) (*ElementSendKeysResponse, error)
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
