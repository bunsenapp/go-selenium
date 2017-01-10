package goselenium

import (
	"encoding/base64"
	"fmt"
)

// ScreenshotResponse is the response returned from the Screenshot and
// ScreenshotElement methods.
type ScreenshotResponse struct {
	State        string
	EncodedImage string
}

// ImageBytes is a helpful function for decoding the base64 encoded image URL.
// The image returned is a PNG image and as such can be manipulated by the
// image/png package. Trying to save this as any other image type will
// result in it failing to open.
func (s *ScreenshotResponse) ImageBytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(s.EncodedImage)
}

func (s *seleniumWebDriver) Screenshot() (*ScreenshotResponse, error) {
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("Screenshot")
	}

	var err error

	url := fmt.Sprintf("%s/session/%s/screenshot", s.seleniumURL, s.sessionID)

	resp, err := s.valueRequest(&request{
		url:           url,
		method:        "GET",
		body:          nil,
		callingMethod: "Screenshot",
	})
	if err != nil {
		return nil, err
	}

	return &ScreenshotResponse{State: resp.State, EncodedImage: resp.Value}, nil
}
