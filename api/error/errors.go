package errors

import (
	"fmt"
	"net/http"
)

type Kind int

const (
	ErrUnknown Kind = iota
	ErrNotFound
	ErrValidation
	ErrConflict
)

type APIError struct {
	Kind    Kind              `json:"-"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
	Err     error             `json:"-"`
}

func (e *APIError) Error() string { return e.Message }
func (e *APIError) Unwrap() error { return e.Err }
func (e *APIError) StatusCode() int {
	switch e.Kind {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrValidation:
		return http.StatusBadRequest
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func newError(kind Kind, msg string, args ...interface{}) *APIError {
	return &APIError{
		Kind:    kind,
		Message: fmt.Sprintf(msg, args...),
	}
}
func Wrap(kind Kind, cause error, msg string, args ...interface{}) *APIError {
	return &APIError{
		Kind:    kind,
		Message: fmt.Sprintf(msg, args...),
		Err:     cause,
	}
}

func NewUnknown(msg string, args ...interface{}) *APIError {
	return newError(ErrUnknown, msg, args...)
}

func NewNotFound(msg string, args ...interface{}) *APIError {
	return newError(ErrNotFound, msg, args...)
}

func NewValidation(msg string, args ...interface{}) *APIError {
	return newError(ErrValidation, msg, args...)
}

func NewConflict(msg string, args ...interface{}) *APIError {
	return newError(ErrConflict, msg, args...)
}
