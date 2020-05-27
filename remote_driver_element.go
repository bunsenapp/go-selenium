package goselenium

import (
	"encoding/json"
	"errors"
	"fmt"
)

type findElementResponse struct {
	E map[string]string `json:"value"`
}

type findElementsResponse struct {
	E map[string]string `json:"value"`
}

func (s *seleniumWebDriver) FindElement(by By) (Element, error) {
	if by.Type() == "index" {
		return nil, errors.New("findelement: invalid by argument")
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

	var values []string
	for _, value := range response.E {
		values = append(values, value)
	}
	if len(values) > 0 {
		return newSeleniumElement(values[0], s), nil
	}
	return nil, errors.New("Not Found")
}

func (s *seleniumWebDriver) FindElements(by By) ([]Element, error) {
	if by.Type() == "index" {
		return nil, errors.New("findelements: invalid by argument")
	}
	if len(s.sessionID) == 0 {
		return nil, newSessionIDError("FindElements")
	}

	var response findElementsResponse
	var err error

	url := fmt.Sprintf("%s/session/%s/elements", s.seleniumURL, s.sessionID)

	resp, err := s.elementRequest(&elRequest{
		url:           url,
		by:            by,
		method:        "POST",
		callingMethod: "FindElements",
	})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, newUnmarshallingError(err, "FindElements", string(resp))
	}

	elements := make([]Element, len(response.E))
	for _, el := range response.E {
		elements = append(elements, newSeleniumElement(el, s))
	}

	return elements, nil
}
