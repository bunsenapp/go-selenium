package goselenium

import (
	"encoding/json"
	"fmt"
)

type findElementResponse struct {
	E element `json:"value"`
}

type element struct {
	ID string `json:"element"`
}

func (s *seleniumWebDriver) FindElement(by By) (Element, error) {
	if by.Type() == "index" {
		return nil, newInvalidArgumentError("Cannot find by index", "by", "index")
	}
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("FindElement")
	}

	var response findElementResponse
	var err error

	url := fmt.Sprintf("%s/session/%s/element", s.seleniumURL, s.sessionID)

	resp, err := s.elementRequest(&elRequest{
		url:           url,
		by:            by,
		method:        "POST",
		callingMethod: "FindElement",
	})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, "FindElement", string(resp))
	}

	el := newSeleniumElement(response.E.ID, s)
	return el, nil
}
