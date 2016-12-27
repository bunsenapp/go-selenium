package integrationtests

import "testing"

func Test_ScreenshotScreenshot_ScreenshotCanBeTakenSuccessfully(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	goResp, err := driver.Go("https://google.com")
	if err != nil || goResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst navigating or result was not a success.", err)
	}

	scs, err := driver.Screenshot()
	if err != nil || scs.State != "success" || len(scs.EncodedImage) == 0 {
		errorAndWrap(t, "Error thrown whilst taking screenshot.", err)
	}

	dec, err := scs.ImageBytes()
	if err != nil || len(dec) == 0 {
		errorAndWrap(t, "Error thrown whilst getting image bytes of screenshot.", err)
	}

	printObjectResult(dec)
}
