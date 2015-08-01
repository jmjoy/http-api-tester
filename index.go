package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.html").ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, map[string]string{
		"Config": getConfigJson().String(),
	})
}

func getConfigJson() *bytes.Buffer {
	src, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	dst := new(bytes.Buffer)
	err = json.Compact(dst, src)
	if err != nil {
		panic(err)
	}

	return dst
}
