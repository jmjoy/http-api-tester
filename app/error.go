package app

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type StatusError struct {
	status  int
	message string
}

func NewStatusError(status int, message string) *StatusError {
	return &StatusError{
		status:  status,
		message: message,
	}
}

func (this *StatusError) Error() string {
	return this.message
}

func (this *StatusError) NewMessage(message interface{}) *StatusError {
	return NewStatusError(this.status, fmt.Sprint(message))
}

func (this *StatusError) NewMessageSpf(args ...interface{}) *StatusError {
	return NewStatusError(this.status, fmt.Sprintf(this.message, args...))
}

func (this *StatusError) isNameEqual(err error) bool {
	return reflect.TypeOf(this).Name() == reflect.TypeOf(err).Name()
}

type ApiStatusError struct {
	*StatusError
}

func NewApiStatusError(status int, message string) *ApiStatusError {
	return &ApiStatusError{
		StatusError: NewStatusError(status, message),
	}
}

// definded error
var (
	ErrBadRequest          = NewStatusError(http.StatusBadRequest, "bad request")
	ErrNotFound            = NewStatusError(http.StatusNotFound, "not found")
	ErrMethodNotAllowed    = NewStatusError(http.StatusMethodNotAllowed, "method not allowed")
	ErrInternalServerError = NewStatusError(http.StatusInternalServerError, "internal server error")

	ErrBucketNotFound = errors.New("Bucket not found")
)
