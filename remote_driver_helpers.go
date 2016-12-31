package goselenium

import "time"

// Until represents a function that will be continuously repeated until it
// succeeds or a timeout is reached.
type Until func(w WebDriver) bool

// UntilElementPresent attempts to locate an element on the page. It is
// determined as existing if the state is 'Success' and the error is nil.
func UntilElementPresent(by By) Until {
	return func(w WebDriver) bool {
		_, err := w.FindElement(by)
		return err == nil
	}
}

// UntilURLIs checks whether or not the page's URL has changed.
func UntilURLIs(url string) Until {
	return func(w WebDriver) bool {
		resp, err := w.CurrentURL()
		return err == nil && resp.URL == url
	}
}

func (s *seleniumWebDriver) Wait(u Until, timeout time.Duration, sleep time.Duration) bool {
	response := make(chan bool, 1)
	quit := make(chan bool, 1)

	go func() {
	outer:
		for {
			select {
			case <-quit:
				break outer
			default:
				e := u(s)
				if e {
					response <- true
					break outer
				}
			}

			time.Sleep(sleep)
		}
	}()

	select {
	case r := <-response:
		return r
	case <-time.After(timeout):
		close(quit)
		return false
	}
}
