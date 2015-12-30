package model

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRequestMaker(t *testing.T) {
	reqMaker, err := NewRequestMaker(testData0)
	if err != nil {
		panic(err)
	}
	req, err := reqMaker.NewRequest()
	if err != nil {
		panic(err)
	}
	t.Log("request content-type:", req.Header.Get("Content-Type"))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	t.Log(string(buf))

	if "k1=v1&k2=v2 k3=v3&k4=v4" != string(buf) {
		t.Fatal("response content not correct!")
	}
}
