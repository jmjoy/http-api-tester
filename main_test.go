package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func dealRespBody(method, urlStr string, bodyData interface{}, fn func([]byte) error) error {
	var body io.Reader
	if bodyData != nil {
		buf, err := json.Marshal(bodyData)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = fn(buf); err != nil {
		return err
	}
	return nil
}
