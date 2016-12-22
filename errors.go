package goselenium

import (
	"encoding/json"
	"fmt"
)

// ErrorResponse is what is returned from the Selenium API when an error
// occurs.
type ErrorResponse struct {
	Message string `json:"state"`
	Status  int    `json:"status"`
}

// CommunicationError is the result of a communication failure between
// this library and the WebDriver API.
type CommunicationError struct {
	url      string
	response *ErrorResponse
	method   string
}

// Error returns a formatted communication error string.
func (c CommunicationError) Error() string {
	return fmt.Sprintf("%s: api error, url: %s, err: %+v", c.method, c.url, c.response)
}

// IsCommunicationError checks whether an error is a selenium communication
// error.
func IsCommunicationError(err error) bool {
	_, ok := err.(CommunicationError)
	return ok
}

func newCommunicationError(err error, method string, url string, resp []byte) CommunicationError {
	var convertedResponse ErrorResponse
	json.Unmarshal(resp, &convertedResponse)

	return CommunicationError{
		url:      url,
		response: &convertedResponse,
		method:   method,
	}
}

// UnmarshallingError is the result of an unmarshalling failure of a JSON
// string.
type UnmarshallingError struct {
	err    error
	json   string
	method string
}

// Error returns a formatted unmarshalling error string.
func (u UnmarshallingError) Error() string {
	return fmt.Sprintf("%s: unmarshalling error, json: %s, err: %s", u.method, u.json, u.err)
}

// IsUnmarshallingError checks whether an error is a selenium unmarshalling
// error.
func IsUnmarshallingError(err error) bool {
	_, ok := err.(UnmarshallingError)
	return ok
}

func newUnmarshallingError(err error, method string, json string) UnmarshallingError {
	return UnmarshallingError{
		err:    err,
		json:   json,
		method: method,
	}
}

// MarshallingError is an error that is returned when a json.Marshal error occurs.
type MarshallingError struct {
	err    error
	object interface{}
	method string
}

// Error returns a formatted marshalling error string.
func (m MarshallingError) Error() string {
	return fmt.Sprintf("%s: marshalling error for object %+v, err: %s", m.method, m.object, m.err.Error())
}

// IsMarshallingError checks whether an error is a marshalling error.
func IsMarshallingError(err error) bool {
	_, ok := err.(MarshallingError)
	return ok
}

func newMarshallingError(err error, method string, obj interface{}) MarshallingError {
	return MarshallingError{
		err:    err,
		object: obj,
		method: method,
	}
}

// SessionIDError is an error that is returned when the session id is
// invalid. This value will contain the method that the session error occurred
// in.
type SessionIDError string

// Error returns a formatted session error string.
func (s SessionIDError) Error() string {
	return fmt.Sprintf("%s: session id is invalid (have you created a session yet?)", string(s))
}

// IsSessionIDError checks whether an error is due to a session ID not being
// set.
func IsSessionIDError(err error) bool {
	_, ok := err.(SessionIDError)
	return ok
}

func newSessionIDError(method string) SessionIDError {
	return SessionIDError(method)
}

// InvalidURLError is an error that is returned whenever a URL is not correctly
// formatted.
type InvalidURLError string

// Error returns the formatted invalid error string.
func (i InvalidURLError) Error() string {
	return fmt.Sprintf("invalid url: %s", string(i))
}

// IsInvalidURLError checks whether an error is due to the URL being incorrectly
// formatted.
func IsInvalidURLError(err error) bool {
	_, ok := err.(InvalidURLError)
	return ok
}

// InvalidURLError
func newInvalidURLError(url string) InvalidURLError {
	return InvalidURLError(url)
}
