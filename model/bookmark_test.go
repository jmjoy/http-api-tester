package model

import (
	"reflect"
	"testing"

	"github.com/jmjoy/http-api-tester/errors"
)

func TestBookmarkCRUD(t *testing.T) {
	testkey := "testBookmark"

	name, has, err := BookmarkModel.GetCurrentKey()
	t.Log("name:", name)
	if err != nil {
		panic(err)
	}
	if has {
		t.Fatal("Has set current bookmark?")
	}

	bookmark, err := BookmarkModel.GetCurrent()
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(bookmark.Data, DataDefault()) ||
		bookmark.Name != BOOKMARK_DEFAULT_NAME {
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

	_, err = BookmarkModel.SetCurrent("not-existed")
	if err != errors.ErrBookmarkNotFound {
		t.Fatal("found?")
	}
	data, err := BookmarkModel.SetCurrent(testkey)
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(data, defaultBookmark.Data) {
		t.Fatal("not equal?")
	}

	bookmark, err = BookmarkModel.GetCurrent()
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(bookmark.Data, DataDefault()) || bookmark.Name != testkey {
		t.Fatal("not equal?")
	}

	if err = BookmarksModel.Delete(testkey); err != nil {
		panic(err)
	}
}
