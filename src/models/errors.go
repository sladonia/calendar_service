package models

import (
	"errors"
	"fmt"
)

var (
	BasicModelError = errors.New("model error")
	EmptyIdError = NewModeError("unable to delete user. id should be provided")
)

type ModelError struct {
	Msg string
	Err error
}

func NewModeError(msg string) *ModelError {
	return &ModelError{
		Msg: msg,
		Err: BasicModelError,
	}
}

func (e *ModelError) Error() string {
	if e.Err == nil {
		return e.Msg
	}
	return fmt.Sprintf("%s: %s", e.Err.Error(), e.Msg)
}

func (e *ModelError) Unwrap() error {
	return e.Err
}
