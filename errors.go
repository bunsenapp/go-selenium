package goselenium

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// CommunicationError represents an error that is returned whilst communicating
// with the web driver service.
type CommunicationError interface {
	error

	URL() string
	Response() *ErrorResponse
}

// ErrorResponse is what is returned from the Selenium API when an error
// occurs.
type ErrorResponse struct {
	Message string `json:"state"`
	Status  int    `json:"status"`
}

// UnmarshallingError represents an error that is returned whilst unmarshalling
// a JSON string into an object.
type UnmarshallingError interface {
	error

	JSON() string
}

// MarshallingError is an error that is returned when a json.Marshal error occurs.
type MarshallingError interface {
	error

	Object() interface{}
}

// SessionIDError is an error that is returned when the WebDriver instance
// does not have a session id set.
type SessionIDError interface {
	error
}

// InvalidURLError is an error that is returned whenever a URL is not correctly
// formatted.
type InvalidURLError interface {
	error

	URL() string
}

// InvalidArgumentError is an error that is returned whenever a function
// argument is invalid.
type InvalidArgumentError interface {
	error

	ArgName() string
	ArgValue() string
}

// IsCommunicationError checks whether an error is a selenium communication
// error.
func IsCommunicationError(err error) bool {
	_, ok := err.(CommunicationError)
	return ok
}

// IsUnmarshallingError checks whether an error is a selenium unmarshalling
// error.
func IsUnmarshallingError(err error) bool {
	_, ok := err.(UnmarshallingError)
	return ok
}

// IsMarshallingError checks whether an error is a marshalling error.
func IsMarshallingError(err error) bool {
	_, ok := err.(MarshallingError)
	return ok
}

// IsSessionIDError checks whether an error is due to a session ID not being
// set.
func IsSessionIDError(err error) bool {
	_, ok := err.(*sessionIDError)
	return ok
}

// IsInvalidURLError checks whether an error is due to the URL being incorrectly
// formatted.
func IsInvalidURLError(err error) bool {
	_, ok := err.(InvalidURLError)
	return ok
}

// IsInvalidArgumentError checks whether an error is due to a function argument
// being incorrect.
func IsInvalidArgumentError(err error) bool {
	_, ok := err.(InvalidArgumentError)
	return ok
}

// CommunicationError
func newCommunicationError(err error, method string, url string, resp []byte) *communicationError {
	var convertedResponse ErrorResponse
	json.Unmarshal(resp, &convertedResponse)

	wrappedErr := errors.Wrap(err, fmt.Sprintf("An API error occurred in %s", method))
	return &communicationError{
		err:  wrappedErr,
		url:  url,
		resp: &convertedResponse,
	}
}

type communicationError struct {
	err  error
	url  string
	resp *ErrorResponse
}

func (c *communicationError) Error() string {
	return c.err.Error()
}

func (c *communicationError) URL() string {
	return c.url
}

func (c *communicationError) Response() *ErrorResponse {
	return c.resp
}

// UnmarshallingError
func newUnmarshallingError(err error, method string, json string) *unmarshallingError {
	return &unmarshallingError{
		err:  errors.Wrap(err, fmt.Sprintf("An unmarshalling error occurred in %s", method)),
		json: json,
	}
}

type unmarshallingError struct {
	err  error
	json string
}

func (u *unmarshallingError) Error() string {
	return u.err.Error()
}

func (u *unmarshallingError) JSON() string {
	return u.json
}

// SessionIDError
func newSessionIDError(method string) *sessionIDError {
	return &sessionIDError{
		err: errors.New(fmt.Sprintf("Session id not set in %s", method)),
	}
}

type sessionIDError struct {
	err error
}

func (s *sessionIDError) Error() string {
	return s.err.Error()
}

// MarshallingError
func newMarshallingError(err error, method string, obj interface{}) *marshallingError {
	return &marshallingError{
		err: errors.Wrap(err, fmt.Sprintf("A marshalling error occurred in %s", method)),
		obj: obj,
	}
}

type marshallingError struct {
	err error
	obj interface{}
}

func (m *marshallingError) Error() string {
	return m.err.Error()
}

func (m *marshallingError) Object() interface{} {
	return m.obj
}

// InvalidURLError
func newInvalidURLError(err error, url string) *invalidURLError {
	return &invalidURLError{
		err: err,
		url: url,
	}
}

type invalidURLError struct {
	err error
	url string
}

func (i *invalidURLError) Error() string {
	return i.err.Error()
}

func (i *invalidURLError) URL() string {
	return i.url
}

// ArgumentError
func newInvalidArgumentError(err string, argName string, argValue string) *invalidArgumentError {
	return &invalidArgumentError{
		err:      errors.New(err),
		argName:  argName,
		argValue: argValue,
	}
}

type invalidArgumentError struct {
	err      error
	argName  string
	argValue string
}

func (i *invalidArgumentError) Error() string {
	return i.err.Error()
}

func (i *invalidArgumentError) ArgName() string {
	return i.argName
}

func (i *invalidArgumentError) ArgValue() string {
	return i.argValue
}
