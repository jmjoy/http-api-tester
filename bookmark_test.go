package main

import (
	"encoding/json"
	"testing"
)

var defaultData = `{"Method":"GET","Url":"","Args":[],"Bm":{"Switch":false,"N":0,"C":0},"Plugin":{"Key":"","Data":{}}}`

var insertData = `{
	"Method":"POST",
	"Url":"http://www.baidu.com",
	"Args":[
		{"Key": "k1", "Value": "v1", "Method": "GET"},
		{"Key": "k2", "Value": "v2", "Method": "POST"}
	],
	"Bm":{"Switch":true,"N":99,"C":9},
	"Plugin":{"Key":"","Data":{}}
}`

var updateData = `{
	"Method":"GET",
	"Url":"http://www.google.com",
	"Args":[
		{"Key": "k1", "Value": "v1", "Method": "GET"},
		{"Key": "k2", "Value": "v2", "Method": "POST"},
		{"Key": "k3", "Value": "v3", "Method": "POST"}
	],
	"Bm":{"Switch":false,"N":999,"C":99},
	"Plugin":{"Key":"","Data":{}}
}`

var insertBookmark = getBookmark("test", insertData)

var updateBookmark = getBookmark("test", updateData)

var successJsonStr = `{"Data":null,"Message":"","Status":200}`

var hasInsertJsonStr = `{"Data":null,"Message":"该书签名字已经存在了","Status":4000}`

func TestBookmarkCRUD(t *testing.T) {
	// get default data
	dealRespBody("GET", "http://localhost:8080/bookmark", nil, func(jsonStr string) error {
		t.Log(jsonStr)
		dataStr, err := getJsonStrOfSuccessBody(jsonStr)
		if err != nil {
			return err
		}
		if !isJsonStrEqual(dataStr, defaultData) {
			t.Fatal("default data not match")
		}
		return nil
	}, t)

	// insert bookmark
	dealRespBody("POST", "http://localhost:8080/bookmark", insertBookmark, func(jsonStr string) error {
		t.Log(jsonStr)
		if !isJsonStrEqual(jsonStr, successJsonStr) {
			t.Fatal("insert bookmark not success")
		}
		return nil
	}, t)

	// get inserted bookmark data
	dealRespBody("GET", "http://localhost:8080/bookmark?name=test", nil, func(jsonStr string) error {
		t.Log(jsonStr)
		dataStr, err := getJsonStrOfSuccessBody(jsonStr)
		if err != nil {
			return err
		}
		if !isJsonStrEqual(dataStr, insertData) {
			t.Fatal("get test data not correct")
		}
		return nil
	}, t)

	// check duplicate insert
	dealRespBody("POST", "http://localhost:8080/bookmark", insertBookmark, func(jsonStr string) error {
		t.Log(jsonStr)
		if !isJsonStrEqual(jsonStr, hasInsertJsonStr) {
			t.Fatal("can insert duplicate bookmark")
		}
		return nil
	}, t)

	// update bookmark
	dealRespBody("PUT", "http://localhost:8080/bookmark", updateBookmark, func(jsonStr string) error {
		t.Log(jsonStr)
		if !isJsonStrEqual(jsonStr, successJsonStr) {
			t.Fatal("insert bookmark not success")
		}
		return nil
	}, t)

	// get updated bookmark data
	dealRespBody("GET", "http://localhost:8080/bookmark?name=test", nil, func(jsonStr string) error {
		t.Log(jsonStr)
		dataStr, err := getJsonStrOfSuccessBody(jsonStr)
		if err != nil {
			return err
		}
		if !isJsonStrEqual(dataStr, updateData) {
			t.Fatal("get test data not correct")
		}
		return nil
	}, t)

	// get updated bookmark data
	dealRespBody("DELETE", "http://localhost:8080/bookmark?name=test", nil, func(jsonStr string) error {
		t.Log(jsonStr)
		if !isJsonStrEqual(jsonStr, successJsonStr) {
			t.Fatal("delete bookmark not success")
		}
		return nil
	}, t)

	// get deleted bookmark data
	dealRespBody("GET", "http://localhost:8080/bookmark?name=test", nil, func(jsonStr string) error {
		t.Log(jsonStr)
		dataStr, err := getJsonStrOfSuccessBody(jsonStr)
		if err != nil {
			return err
		}
		if !isJsonStrEqual(dataStr, defaultData) {
			t.Fatal("default data not match")
		}
		return nil
	}, t)
}

func getBookmark(name string, dataStr string) string {
	var data interface{}
	err := json.Unmarshal([]byte(dataStr), &data)
	if err != nil {
		panic(err)
	}

	buf, err := json.Marshal(map[string]interface{}{
		"Name": name,
		"Data": data,
	})
	if err != nil {
		panic(err)
	}

	return string(buf)
}
