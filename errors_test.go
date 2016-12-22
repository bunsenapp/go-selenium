package goselenium

import (
	"errors"
	"testing"
)

func communicationError() error {
	return newCommunicationError(errors.New(":<"), "Test", "", nil)
}

func sessionError() error {
	return newSessionIDError("Test")
}

func unmarshallingError() error {
	return newUnmarshallingError(errors.New(":<"), "Test", "")
}

func marshallingError() error {
	return newMarshallingError(errors.New(":<"), "Test", "test")
}

func Test_Errors_CommunicationErrorCanBeCastSuccessfully(t *testing.T) {
	e := communicationError()

	back, ok := e.(CommunicationError)
	if !ok || back.method != "Test" {
		t.Errorf("Could not assert error")
	}
}

func Test_Errors_SessionErrorCanBeCastSuccessfully(t *testing.T) {
	e := sessionError()

	_, ok := e.(SessionIDError)
	if !ok {
		t.Errorf("Could not assert error")
	}
}

func Test_Errors_UnmarshallingErrorCanBeCastSuccessfully(t *testing.T) {
	e := unmarshallingError()

	body, ok := e.(UnmarshallingError)
	if !ok || body.method != "Test" {
		t.Errorf("Could not assert error")
	}
}

func Test_Errors_MarshallingErrorCanBeCastSuccessfully(t *testing.T) {
	e := marshallingError()

	body, ok := e.(MarshallingError)
	if !ok || body.method != "Test" {
		t.Errorf("Could not assert error")
	}
}
