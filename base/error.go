package base

import (
	"fmt"
	"net/http"
)

var _ error = &statusError{}

type statusError struct {
	code    int
	message string
}

func NewStatusError(code int, message string) *statusError {
	return &statusError{
		code:    code,
		message: message,
	}
}

func (this *statusError) Error() string {
	return fmt.Sprintf("[%d] %s", this.code, this.message)
}

var _ error = &apiStatusError{}

type apiStatusError struct {
	*statusError
}

func NewApiStatusError(code int, message string) *apiStatusError {
	return &apiStatusError{
		statusError: NewStatusError(code, message),
	}
}

// System StatusError
var (
	ErrorMethodNotAllowed = NewStatusError(http.StatusMethodNotAllowed, "405 method not allowed")
)
