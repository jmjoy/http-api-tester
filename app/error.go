package app

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type IStatusError interface {
	error
	IamStatusError()
}

type IApiStatusError interface {
	error
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

func errorGetStatus(err error) int {
	field, has := reflect.TypeOf(err).FieldByName("Status")
	if !has {
		return 0
	}
	status, _ := strconv.Atoi(field.Tag.Get("apiError"))
	return status
}

func errorGetMessage(err error) string {
	name := "Message"

	v := reflect.ValueOf(err)
	valueField := v.FieldByName(name)
	if i := valueField.Interface(); i != nil && i != "" {
		return i.(string)
	}
	field, has := v.Type().FieldByName(name)
	if !has {
		return ""
	}
	return field.Tag.Get("apiError")
}
