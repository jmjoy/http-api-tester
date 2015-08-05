package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var restfulPools = make(map[string]*sync.Pool)

type Restful interface {
	SetWR(http.ResponseWriter, *http.Request)
	Get()
	Post()
	Put()
	Delete()
}

func HandleRestful(pattern string, fn func() interface{}) {
	restfulPools[pattern] = &sync.Pool{New: fn}

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		pool, ok := restfulPools[pattern]
		if !ok {
			panic(fmt.Sprintf("HandleRestful: pool for %q not found", pattern))
		}
		rf := pool.Get().(Restful)
		defer pool.Put(rf)
		rf.SetWR(w, r)

		switch r.Method {
		case "GET":
			rf.Get()

		case "POST":
			rf.Post()

		case "PUT":
			rf.Put()

		case "DELETE":
			rf.Delete()

		default:
			w.Write([]byte("I don't know this method, get out please!"))
		}
	})
}

type Controller struct {
	w http.ResponseWriter
	r *http.Request
}

func (this *Controller) SetWR(w http.ResponseWriter, r *http.Request) {
	this.w = w
	this.r = r
}

func (this *Controller) RenderJson(code int, msg string, data interface{}) {
	out := map[string]interface{}{
		"status": code,
		"msg":    msg,
		"data":   data,
	}
	buf, err := json.Marshal(out)
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
