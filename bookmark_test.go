package main

import "testing"

func TestBookmarkGetCurrent(t *testing.T) {
	dealRespBodyJsonStruct("http://localhost:8080/bookmark", func(i interface{}) error {
		return nil
	})
}
