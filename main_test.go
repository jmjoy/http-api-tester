package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/bitly/go-simplejson"
)

func TestMain(m *testing.M) {
	os.Remove("http-api-tester.db")
	m.Run()
}

func dealRespBody(method, urlStr string, bodyData interface{}, fn func(string) error, t *testing.T) {
	err := func() error {
		var body io.Reader
		if bodyData != nil {
			if str, ok := bodyData.(string); ok {
				body = strings.NewReader(str)

			} else {
				buf, err := json.Marshal(bodyData)
				if err != nil {
					return err
				}
				body = bytes.NewBuffer(buf)
			}
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

		if err = fn(string(buf)); err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		t.Fatal(err)
	}
}

func isJsonStrEqual(strs ...string) bool {
	formatStrs := make([]string, len(strs))
	for i, str := range strs {
		var data interface{}
		err := json.Unmarshal([]byte(str), &data)
		if err != nil {
			panic(err)
		}
		buf, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		formatStrs[i] = string(buf)
	}
	for i := 0; i < len(formatStrs)-1; i++ {
		if formatStrs[i] != formatStrs[i+1] {
			return false
		}
	}
	return true
}

func getJsonStrOfSuccessBody(body string) (string, error) {
	json, err := simplejson.NewJson([]byte(body))
	if err != nil {
		return "", err
	}
	status, err := json.Get("Status").Int()
	if err != nil {
		return "", err
	}
	if status != 200 {
		return "", fmt.Errorf("status isn't 200")
	}
	data, err := json.Get("Data").Encode()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
