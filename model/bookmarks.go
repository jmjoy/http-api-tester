package model

import (
	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/errors"
)

var BookmarksModel = &bookmarksModel{
	Model: app.NewModel("bookmarks"),
}

type bookmarksModel struct {
	*app.Model
}

func (this *bookmarksModel) Get(name string) (data Data, err error) {
	if err = this.validateBookmarkName(name); err != nil {
		return
	}

	has, err := this.Model.Get(name, &data)
	if err != nil {
		return
	}

	if !has {
		err = errors.ErrBookmarkNotFound
		return
	}

	return
}

func (this *bookmarksModel) Upsert(bookmark Bookmark, typ UpsertType) (err error) {
	if err = this.validateBookmarkName(bookmark.Name); err != nil {
		return
	}

	// check is exists when add
	if typ == UPSERT_ADD {
		var data Data
		var has bool

		has, err = this.Model.Get(bookmark.Name, &data)
		if err != nil {
			return
		}
		if has {
			return errors.ErrBookmarkExisted
		}
	}

	return this.Put(bookmark.Name, bookmark.Data)
}

func (this *bookmarksModel) Delete(name string) (err error) {
	if err = this.validateBookmarkName(name); err != nil {
		return
	}

	return this.Model.Delete(name)
}

func (this *bookmarksModel) validateBookmarkName(name string) error {
	if name == "" {
		return errors.ErrBookmarkNameEmpty
	}

	// TODO 暂时允许所有名字
	return nil
}
