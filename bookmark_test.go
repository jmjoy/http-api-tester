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
	"Plugin":{"Key":"","Data":null}
}`

var insertBookmark = getInsertBookmark("test", insertData)

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
}

func getInsertBookmark(name string, dataStr string) string {
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
