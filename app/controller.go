package app

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"encoding/json"
)

// IController is a interface named IController
type IController interface {
	// reset rw of one context
	Reset(http.ResponseWriter, *http.Request)

	// RESTful
	Get() error
	Post() error
	Put() error
	Delete() error
}

func ResetController(c IController, w http.ResponseWriter, r *http.Request) {
	field := reflect.ValueOf(c).FieldByName("Controller")

	panic(field.CanSet())
}

var _ IController = new(Controller)

// Controller is a struct named Controller
type Controller struct {
	W http.ResponseWriter
	R *http.Request

	query url.Values // get params
}

// Reset reset controller for handle one request
func (this *Controller) Reset(w http.ResponseWriter, r *http.Request) {
	if this == nil {
		this = new(Controller)
	}

	this.W = w
	this.R = r
}

func (this *Controller) MethodNotAllowed() error {
	return ErrMethodNotAllowed
}

func (this *Controller) Get() error {
	return this.MethodNotAllowed()
}

func (this *Controller) Post() error {
	return this.MethodNotAllowed()
}

func (this *Controller) Put() error {
	return this.MethodNotAllowed()
}

func (this *Controller) Delete() error {
	return this.MethodNotAllowed()
}

func (this *Controller) JsonRender(status int, message string, data interface{}) error {
	return jsonRender(this.W, status, message, data)
}

func (this *Controller) JsonSuccess(data interface{}) error {
	return this.JsonRender(200, "", data)
}

func (this *Controller) QueryGet(key string) string {
	if this.query == nil {
		this.query = this.R.URL.Query()
	}
	return this.query.Get(key)
}

func (this *Controller) ParseJsonBody(data interface{}) (err error) {
	// Get Body
	buf, err := ioutil.ReadAll(this.R.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(buf, data); err != nil {
		return ErrBadRequest.NewMessage(err)
	}

	return
}
