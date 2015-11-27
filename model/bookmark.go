package model

import (
	"encoding/json"

	"github.com/jmjoy/http-api-tester/base"
)

type BookmarkModel struct {
}

func NewBookmarkModel() *BookmarkModel {
	return new(BookmarkModel)
}

func (this *BookmarkModel) GetCurrent() (Data, error) {
	bookmarkName, err := base.Db.Get("bookmark", "selected")
	if err != nil {
		return this.handleGetError(err)
	}

	bookmark, err := base.Db.Get("bookmarks", string(bookmarkName))
	if err != nil {
		return this.handleGetError(err)
	}

	var data Data
	err = json.Unmarshal(bookmark, &data)
	return data, err
}

func (this *BookmarkModel) handleGetError(err error) (Data, error) {
	if err == base.ErrorBucketNotFound {
		return this.DefaultData(), nil
	}
	return Data{}, err
}

func (this *BookmarkModel) DefaultData() Data {
	return Data{
		Method: "GET",
		Args:   []Arg{},
	}
}
