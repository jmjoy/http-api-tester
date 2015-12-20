package model

import (
	"testing"

	"reflect"

	"github.com/jmjoy/http-api-tester/errors"
)

var testData = Data{}

func TestCRUD(t *testing.T) {
	testkey := "test"

	_, err := BookmarksModel.Get(testkey)
	if err != errors.ErrBookmarkNotFound {
		t.Fatal("bookmark existd?")
	}

	var defaultBookmark = Bookmark{
		Name: testkey,
		Data: DataDefault(),
	}

	err = BookmarksModel.Upsert(defaultBookmark, UPSERT_ADD)
	if err != nil {
		panic(err)
	}

	err = BookmarksModel.Upsert(defaultBookmark, UPSERT_ADD)
	if err != errors.ErrBookmarkExisted {
		t.Fatal("Bookmark not existed?")
	}

	data, err := BookmarksModel.Get(testkey)
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(data, defaultBookmark.Data) {
		t.Fatal("Data not equal!")
	}

	// new bookmark
	newBookmark := defaultBookmark
	defaultBookmark.Data.Method = "POST"
	defaultBookmark.Data.Url = "http://www.baidu.com"

	err = BookmarksModel.Upsert(newBookmark, UPSERT_UPDATE)
	if err != nil {
		panic(err)
	}

	data, err = BookmarksModel.Get(testkey)
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(data, newBookmark.Data) {
		t.Fatal("Data not equal!")
	}

	err = BookmarksModel.Delete(testkey)
	if err != nil {
		panic(err)
	}

	_, err = BookmarksModel.Get(testkey)
	if err != errors.ErrBookmarkNotFound {
		t.Fatal("bookmark existd?")
	}
}
