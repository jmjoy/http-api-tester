package app

import (
	"errors"
	"fmt"
)

type IStatusError interface {
	IamStatusError()
}

type IApiStatusError interface {
	IamApiStatusError()
}

var _ error = new(StatusError)
var _ IStatusError = new(StatusError)

type StatusError struct {
	Status  int
	Message string
}

func (this *StatusError) Error() string {
	return fmt.Sprintf("[%d] %s", this.Status, this.Message)
}

func (_ *StatusError) IamStatusError() {}

var _ error = new(ApiStatusError)
var _ IApiStatusError = new(ApiStatusError)

type ApiStatusError struct {
	*StatusError
}

func (_ *ApiStatusError) IamApiStatusError() {}

type ErrorNotFound struct {
	*StatusError
	Status  int    `appError:404`
	Message string `appError:"not found"`
}

type ErrorMethodNotAllowed struct {
	*StatusError
	Status  int    `appError:405`
	Message string `appError:"method not allowed"`
}

var (
	ErrorBucketNotFound   = errors.New("Bucket not found")
	ErrorBookmarkNotFound = errors.New("该书签不存在")
)
