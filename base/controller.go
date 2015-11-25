package base

import (
	"encoding/json"
	"net/http"
)

var _ Restful = &controller{}

type controller struct {
	w http.ResponseWriter
	r *http.Request
}

func NewController(w http.ResponseWriter, r *http.Request) *controller {
	return &controller{
		w: w,
		r: r,
	}
}

func (this *controller) RenderJson(code int, msg string, data interface{}) {
	out := map[string]interface{}{
		"status":  code,
		"message": msg,
		"data":    data,
	}
	buf, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}

	this.w.Write(buf)
}

func (this *controller) MethodNotAllowed() *statusError {
	return ErrorMethodNotAllowed
}

func (this *controller) Get() *statusError {
	return this.MethodNotAllowed()
}

func (this *controller) Post() *statusError {
	return this.MethodNotAllowed()
}

func (this *controller) Put() *statusError {
	return this.MethodNotAllowed()
}

func (this *controller) Delete() *statusError {
	return this.MethodNotAllowed()
}
