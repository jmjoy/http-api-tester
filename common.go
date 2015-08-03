package main

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	w http.ResponseWriter
	r *http.Request
}

func (this *Controller) SetWR(w http.ResponseWriter, r *http.Request) {
	this.w = w
	this.r = r
}

func (this *Controller) RenderJson(code int, msg string, data interface{}) {
	buf, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	this.w.Write(buf)
}

func (this *Controller) Get() {
}

func (this *Controller) Post() {
}

func (this *Controller) Put() {
}

func (this *Controller) Delete() {
}
