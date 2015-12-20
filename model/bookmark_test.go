package model

import (
	"reflect"
	"testing"

	"github.com/jmjoy/http-api-tester/errors"
)

func TestBookmarkCRUD(t *testing.T) {
	testkey := "testBookmark"

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

	err = BookmarkModel.SetCurrent("not-existed")
	if err != errors.ErrBookmarkNotFound {
		t.Fatal("found?")
	}
	err = BookmarkModel.SetCurrent(testkey)
	if err != nil {
		panic(err)
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
