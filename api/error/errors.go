package errors

import "fmt"

type Kind int

const (
	ErrUnknown Kind = iota
	ErrNotFound
	ErrValidation
	ErrConflict
)

type APIError struct {
	Kind    Kind
	Message string
}

func (e *APIError) Error() string {
	return e.Message
}

func New(kind Kind, msg string, args ...interface{}) *APIError {
	return &APIError{
		Kind:    kind,
		Message: fmt.Sprintf(msg, args...),
	}
}
