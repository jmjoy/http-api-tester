package model

import "testing"

func TestBookmarkGetCurrent(t *testing.T) {
	_, err := NewBookmarkModel().GetCurrent()
	t.Fatal(err)
}
