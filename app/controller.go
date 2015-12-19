package app

import "net/http"

// IController is a interface named IController
type IController interface {
	// for handle one request
	SetR(*http.Request)
	SetW(http.ResponseWriter)

	// restful
	Get() error
	Post() error
	Put() error
	Delete() error
}

var _ IController = new(Controller)

// Controller is a struct named Controller
type Controller struct {
	W http.ResponseWriter
	R *http.Request
}

func (this *Controller) SetR(r *http.Request) {
	this.R = r
}

func (this *Controller) SetW(w http.ResponseWriter) {
	this.W = w
}

func (this *Controller) JsonRender(status int, message string, data interface{}) error {
	return jsonRender(this.W, status, message, data)
}

func (this *Controller) JsonSuccess(data interface{}) error {
	return this.JsonRender(200, "", data)
}

func (this *Controller) MethodNotAllowed() error {
	return ErrorMethodNotAllowed
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
