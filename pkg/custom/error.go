package custom

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
	return NewError(fiber.StatusNotFound, code, message, err)
}

func ErrInvalidInput(code, message string, err error) *AppError {
	if message == "" {
		message = "Invalid input provided."
	}
	return NewError(fiber.StatusBadRequest, code, message, err)
}

func ErrUnauthorized(code, message string, err error) *AppError {
	if message == "" {
		message = "Unauthorized access."
	}
	return NewError(fiber.StatusUnauthorized, code, message, err)
}

func ErrForbidden(code, message string, err error) *AppError {
	if message == "" {
		message = "Access to this resource is forbidden."
	}
	return NewError(fiber.StatusForbidden, code, message, err)
}

func ErrIntervalServer(code, message string, err error) *AppError {
	if message == "" {
		message = "An unexcepted internal server error occurred."
	}
	if code == "" {
		code = "INTERNAL_SERVER_ERROR"
	}
	return NewError(fiber.StatusInternalServerError, code, message, err)
}

func ErrConflict(code, message string, err error) *AppError {
	if message == "" {
		message = "A conflict occurred with the current state of the resource."
	}
	return NewError(fiber.StatusConflict, code, message, err)
}
