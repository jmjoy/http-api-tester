package main

import "testing"

var defaultBookmarkResult = `{"data":{"Method":"GET","Url":"","Args":[],"Bm":{"Switch":false,"N":0,"C":0},"Plugin":{"Key":"","Data":null}},"message":"","status":200}`

func TestBookmarkCRUD(t *testing.T) {
	err := dealRespBody("GET", "http://localhost:8080/bookmark", nil, func(buf []byte) error {
		jsonStr := string(buf)
		t.Log(jsonStr)
		if jsonStr != defaultBookmarkResult {
			t.Fatal("default data not match")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

}
