package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func dealRespBody(url string, fn func([]byte) error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err = fn(buf); err != nil {
		panic(err)
	}
}

func dealRespBodyJsonStruct(url string, fn func(i interface{}) error) {
	dealRespBody(url, func(data []byte) error {
		var i interface{}
		if err := json.Unmarshal(data, &i); err != nil {
			return err
		}
		if err := fn(i); err != nil {
			return err
		}
		return nil
	})
}
