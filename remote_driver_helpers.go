package goselenium

import "time"

type waitResponse struct {
	s bool
	e error
}

// Until represents a function that will be continuously repeated until it
// succeeds or a timeout is reached.
type Until func(w WebDriver) (bool, error)

// UntilElementPresent attempts to locate an element on the page. It is
// determined as existing if the state is 'Success' and the error is nil.
func UntilElementPresent(by By, sleep time.Duration) Until {
	return func(w WebDriver) (bool, error) {
		time.Sleep(sleep)

		el, err := w.FindElement(by)
		return err == nil && el != nil, err
	}
}

func (s *seleniumWebDriver) Wait(u Until, timeout time.Duration) (bool, error) {
	response := make(chan *waitResponse, 1)
	quit := make(chan bool, 1)

	go func() {
	outer:
		for {
			select {
			case <-quit:
				break outer
			default:
				s, e := u(s)
				if e == nil && s {
					response <- &waitResponse{s: s, e: e}
					break outer
				}
			}
		}
	}()

	select {
	case r := <-response:
		return r.s, r.e
	case <-time.After(timeout):
		close(quit)
		return false, WaitTimeoutError
	}
}
