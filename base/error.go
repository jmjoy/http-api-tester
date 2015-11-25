package base

import (
	"fmt"
	"net/http"
)

var _ error = &statusError{}

type statusError struct {
	status  int
	message string
}

func NewStatusError(status int, message string) *statusError {
	return &statusError{
		status:  status,
		message: message,
	}
}

func (this *statusError) Error() string {
	return fmt.Sprintf("[%d] %s", this.status, this.message)
}

var _ error = &apiStatusError{}

type apiStatusError struct {
	*statusError
}

func NewApiStatusError(status int, message string) *apiStatusError {
	return &apiStatusError{
		statusError: NewStatusError(status, message),
	}
}

// System StatusError
var (
	ErrorMethodNotAllowed = NewStatusError(http.StatusMethodNotAllowed, "405 method not allowed")
)
