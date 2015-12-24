package model

import (
	"testing"

	"reflect"

	"github.com/jmjoy/http-api-tester/errors"
)

var testData = Data{}

func TestBookmarksCRUD(t *testing.T) {
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

	var inDefaultBookmark = Bookmark{
		Name: BOOKMARK_DEFAULT_NAME,
		Data: DataDefault(),
	}
	err = BookmarksModel.Upsert(inDefaultBookmark, UPSERT_ADD)
	if err != errors.ErrBookmarkEditDefault {
		t.Fatal("Can edit default?")
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

	_, has, err := BookmarkModel.GetCurrentKey()
	if has {
		t.Fatal("has selected bookmark?")
	}

	err = BookmarksModel.Delete(BOOKMARK_DEFAULT_NAME)
	if err != errors.ErrBookmarkEditDefault {
		t.Fatal("Can delete default?")
	}
}

func TestGetAllNames(t *testing.T) {
	testKey := "test0"
	defaultNames := []string{BOOKMARK_DEFAULT_NAME}
	testNames := []string{BOOKMARK_DEFAULT_NAME, testKey}

	names, err := BookmarksModel.GetAllNames()
	if err != nil {
		panic(err)
	}
	t.Log(defaultNames, names)
	if !reflect.DeepEqual(defaultNames, names) {
		t.Fatal("not equal?")
	}

	if err = BookmarksModel.Put(testKey, "hello world"); err != nil {
		panic(err)
	}

	names, err = BookmarksModel.GetAllNames()
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(testNames, names) {
		t.Fatal("not equal?")
	}
}
