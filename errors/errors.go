package errors

import "net/http"

type AppError struct {
	Message string `json:"message" xml:"message"`
	Code    int    `json:"code,omitempty" xml:"code"`
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
