package app

import (
	"errors"
	"fmt"
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

func (this *StatusError) NewMessage(message string) *StatusError {
	return NewStatusError(this.status, message)
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
	ErrorNotFound         = NewStatusError(400, "not found")
	ErrorMethodNotAllowed = NewStatusError(405, "method not allowed")

	ErrorBucketNotFound = errors.New("Bucket not found")
)
