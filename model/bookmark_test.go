package model

import (
	"reflect"
	"testing"

	"github.com/jmjoy/http-api-tester/errors"
)

func TestBookmarkCRUD(t *testing.T) {
	testkey := "testBookmark"

	data, err := BookmarkModel.GetCurrent()
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(data, DataDefault()) {
		t.Fatal("not euqal?")
	}

	var defaultBookmark = Bookmark{
		Name: testkey,
		Data: DataDefault(),
	}
	err = BookmarksModel.Upsert(defaultBookmark, UPSERT_ADD)
	if err != nil {
		panic(err)
	}

	err = BookmarkModel.SetCurrent("not-existed")
	if err != errors.ErrBookmarkNotFound {
		t.Fatal("found?")
	}
	err = BookmarkModel.SetCurrent(testkey)
	if err != nil {
		panic(err)
	}

	data, err = BookmarkModel.GetCurrent()
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(data, DataDefault()) {
		t.Fatal("not equal?")
	}
}
