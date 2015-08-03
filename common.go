package main

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	W http.ResponseWriter
	R *http.Request
}

func (this *Controller) Do(w http.ResponseWriter, r *http.Request) {
	this.W = w
	this.R = r

	switch this.R.Method {
	case "GET":
		this.Get()

	case "POST":
		this.Post()

	case "PUT":
		this.Put()

	case "DELETE":
		this.Delete()

	default:
		this.W.Write("I don't know this method, get out please!")
	}
}

func (this *Controller) RenderJson(code int, msg string, data interface{}) {
	buf, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Write(buf)
}

func (this *Controller) Get() {
}

func (this *Controller) Post() {
}

func (this *Controller) Put() {
}

func (this *Controller) Delete() {
}
