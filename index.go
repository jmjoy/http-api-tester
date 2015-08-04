package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type jsonConfig struct {
	Selected           string
	Bookmarks, Plugins map[string]map[string]interface{}
}

var gJsonConfig = new(jsonConfig)

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

	err = json.Unmarshal(src, gJsonConfig)
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

func saveConfigJson() error {
	buf, err := json.Marshal(gJsonConfig)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	err = json.Indent(buffer, buf, "", "    ")
	if err != nil {
		return err
	}

	fw, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer fw.Close()

	_, err = buffer.WriteTo(fw)
	if err != nil {
		return err
	}

	return nil
}
