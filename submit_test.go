package main

import "testing"

var submitDataNoMethod = `{
	"Method":"",
	"Url":"http://www.baidu.com",
	"Args":[
		{"Key": "k1", "Value": "v1", "Method": "GET"},
		{"Key": "k2", "Value": "v2", "Method": "POST"}
	],
	"Bm":{"Switch":true,"N":99,"C":9},
	"Plugin":{"Key":"","Data":{}}
}`

var submitDataErrorUrl = `{
	"Method":"POST",
	"Url":"sjdkflajskldf",
	"Args":[
		{"Key": "k1", "Value": "v1", "Method": "GET"},
		{"Key": "k2", "Value": "v2", "Method": "POST"}
	],
	"Bm":{"Switch":true,"N":99,"C":9},
	"Plugin":{"Key":"","Data":{}}
}`

var submitData = `{
	"Method":"GET",
	"Url":"http://localhost:8080/bookmarks",
	"Args":[
		{"Key": "k1", "Value": "v1", "Method": "GET"},
		{"Key": "k2", "Value": "v2", "Method": "POST"}
	],
	"Bm":{"Switch":true,"N":99,"C":9},
	"Plugin":{"Key":"","Data":{}}
}`

func TestSubmit(t *testing.T) {
	// submit without plugn
	dealRespBody("GET", "http://localhost:8080?act=initData", nil, func(jsonStr string) error {
		t.Log(jsonStr)
		return nil
	}, t)
}
