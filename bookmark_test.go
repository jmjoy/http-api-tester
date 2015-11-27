package main

import "testing"

var defaultBookmarkResult = `{"data":{"Method":"GET","Url":"","Args":[],"Bm":{"Switch":false,"N":0,"C":0},"Plugin":{"Key":"","Data":{}}},"message":"","status":200}`

var insertBookmark = `{
	"name": "test",
	"data":{
		"Method":"POST",
		"Url":"http://www.baidu.com",
		"Args":[
		{"Key": "k1", "Value": v1, "Method": "GET"},
		{"Key": "k2", "Value": v2, "Method": "POST"},
		],
		"Bm":{"Switch":true,"N":99,"C":9},
		"Plugin":{"Key":"","Data":null}}
}`

func TestBookmarkCRUD(t *testing.T) {
	//failWhenErrorNotNil(dealRespBody("GET", "http://localhost:8080/bookmark", nil, func(buf []byte) error {
	//    jsonStr := string(buf)
	//    t.Log(jsonStr)
	//    if jsonStr != defaultBookmarkResult {
	//        t.Fatal("default data not match")
	//    }
	//    return nil
	//}), t)

	//failWhenErrorNotNil(dealRespBody("POST", "http://www.baidu.com", nil, func(buf []byte) error {
	//    fmt.Println(string(buf))
	//    return nil
	//}), t)

	failWhenErrorNotNil(dealRespBody("POST", "http://localhost:8080/bookmark", nil, func(buf []byte) error {
		jsonStr := string(buf)
		t.Log(jsonStr)
		return nil
	}), t)
}

func failWhenErrorNotNil(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
