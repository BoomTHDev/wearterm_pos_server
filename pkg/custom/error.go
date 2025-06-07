package custom

import (
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("AppError: %s (Code: %s) - %s", e.Message, e.Code, e.Err.Error())
	}

	return fmt.Sprintf("AppError: %s (Code: %s)", e.Message, e.Code)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewError(statusCode int, code string, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

func ErrNotFound(code, message string, err error) *AppError {
	if message == "" {
		message = "The request resource was not found."
	}
	return NewError(http.StatusNotFound, code, message, err)
}

func ErrInvalidInput(code, message string, err error) *AppError {
	if message == "" {
		message = "Invalid input provided."
	}
	return NewError(http.StatusBadRequest, code, message, err)
}

func ErrUnauthorized(code, message string, err error) *AppError {
	if message == "" {
		message = "Unauthorized access."
	}
	return NewError(http.StatusUnauthorized, code, message, err)
}

func ErrForbidden(code, message string, err error) *AppError {
	if message == "" {
		message = "Access to this resource is forbidden."
	}
	return NewError(http.StatusForbidden, code, message, err)
}

func ErrIntervalServer(code, message string, err error) *AppError {
	if message == "" {
		message = "An unexcepted internal server error occurred."
	}
	if code == "" {
		code = "INTERNAL_SERVER_ERROR"
	}
	return NewError(http.StatusInternalServerError, code, message, err)
}

func ErrConflict(code, message string, err error) *AppError {
	if message == "" {
		message = "A conflict occurred with the current state of the resource."
	}
	return NewError(http.StatusConflict, code, message, err)
}
