package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiErrorInterface interface {
	error
	GetMessage() string
	GetStatusCode() int
}

type ApiError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Err        string `json:"error"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("message: %s. status: %d. error: %s", e.Message, e.StatusCode, e.Err)
}

func (e ApiError) GetMessage() string {
	return e.Message
}

func (e ApiError) GetStatusCode() int {
	return e.StatusCode
}

func NewApiError(message, err string, statusCode int) ApiErrorInterface {
	return ApiError{
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}

func NewNotImplementedApiError(message string) ApiErrorInterface {
	return ApiError{
		Message:    message,
		StatusCode: http.StatusNotImplemented,
		Err:        "method not implemented",
	}
}

func NewBadRequestApiError(message string) ApiErrorInterface {
	return ApiError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Err:        "bad request",
	}
}

func NewNotFoundApiError(message string) ApiErrorInterface {
	return ApiError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Err:        "resource not found",
	}
}

func NewApiErrorFromBytes(bytes []byte) (ApiErrorInterface, error) {
	var myError ApiError
	err := json.Unmarshal(bytes, &myError)
	if err != nil {
		return nil, err
	}
	return &myError, nil
}
