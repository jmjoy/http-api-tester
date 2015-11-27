package base

import (
	"errors"
	"fmt"
	"net/http"
)

var _ error = &statusError{}

type statusError struct {
	status  int
	message string
}

func NewStatusError(status int, i interface{}) *statusError {
	return &statusError{
		status:  status,
		message: fmt.Sprint(i),
	}
}

func (this *statusError) Error() string {
	return fmt.Sprintf("[%d] %s", this.status, this.message)
}

var _ error = &apiStatusError{}

type apiStatusError struct {
	*statusError
}

func NewApiStatusError(status int, i interface{}) *apiStatusError {
	return &apiStatusError{
		statusError: NewStatusError(status, i),
	}
}

// System StatusError
var (
	ErrorNotFound         = NewStatusError(http.StatusNotFound, "404 not found")
	ErrorMethodNotAllowed = NewStatusError(http.StatusMethodNotAllowed, "405 method not allowed")
)

var (
	ErrorBucketNotFound = errors.New("Bucket not found")
)
