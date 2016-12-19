package integrationtests

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/bunsenapp/go-selenium"
	"github.com/pkg/errors"
)

func setUp() {
	command := exec.Command("docker", "run", "-d", "--name=goselenium-tests", "-p=4444:4444", "selenium/standalone-firefox")
	command.CombinedOutput()
	time.Sleep(2000 * time.Millisecond)
}

func tearDown() {
	command := exec.Command("docker", "rm", "goselenium-tests", "-f")
	command.CombinedOutput()
}

func errorAndWrap(t *testing.T, message string, oldError error) {
	if oldError == nil {
		t.Errorf(errors.New(message).Error())
	} else {
		err := errors.Wrap(oldError, message)
		t.Errorf(err.Error())
	}
}

func printObjectResult(obj interface{}) {
	envResult := os.Getenv("GOSELENIUM_TEST_DETAIL")
	shouldShowDetailedResults, err := strconv.ParseBool(envResult)
	if shouldShowDetailedResults && err == nil {
		fmt.Println(fmt.Sprintf("Object returned: %+v", obj))
	}
}

func createDriver(t *testing.T) goselenium.WebDriver {
	caps := goselenium.Capabilities{}
	caps.SetBrowser(goselenium.FirefoxBrowser())

	driver, err := goselenium.NewSeleniumWebDriver("http://localhost:4444/wd/hub/", caps)
	if err != nil {
		t.Errorf("Driver creation threw an error.")
	}

	return driver
}
