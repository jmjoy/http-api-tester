package base

import (
	"net/http"
)

var _ Restful = &Controller{}

type Controller struct {
	w http.ResponseWriter
	r *http.Request
}

func NewController(w http.ResponseWriter, r *http.Request) *Controller {
	return &Controller{
		w: w,
		r: r,
	}
}

func (this *Controller) W() http.ResponseWriter {
	return this.w
}

func (this *Controller) R() *http.Request {
	return this.r
}

func (this *Controller) RenderJson(data interface{}) error {
	RenderJson(this.w, 200, "", data)
	return nil
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
