package integrationtests

import (
	"testing"
	"time"

	goselenium "github.com/bunsenapp/go-selenium"
)

func Test_WaitUntilURLIs_WorksCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error whilst creating session", err)
	}

	_, err = driver.Go("https://bunsenapp.github.io/go-selenium/helpers/url-change.html")
	if err != nil {
		errorAndWrap(t, "Error whilst navigating", err)
	}

	ok := driver.Wait(goselenium.UntilURLIs("https://bunsenapp.github.io/test"), 30*time.Second, 0)
	if !ok {
		errorAndWrap(t, "Timeout was exceeded", nil)
	}
}
