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
		// get default bookmark
		if err == errors.ErrBookmarkEditDefault {
			data = DataDefault()
			err = nil
		}
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

func (this *bookmarksModel) GetAllNames() (names []string, err error) {
	otherNames, err := this.Model.Keys()
	if err != nil {
		return
	}
	names = []string{BOOKMARK_DEFAULT_NAME}
	if len(otherNames) != 0 {
		names = append(names, otherNames...)
	}
	return
}

func (this *bookmarksModel) Upsert(bookmark Bookmark, typ UpsertType) (err error) {
	if err = this.validateBookmarkName(bookmark.Name); err != nil {
		return
	}

	var data Data
	has, err := this.Model.Get(bookmark.Name, data)
	if err != nil {
		return
	}

	// check is exists or not
	if typ == UPSERT_ADD {
		if has {
			return errors.ErrBookmarkExisted
		}
	} else {
		if !has {
			return errors.ErrBookmarkNotFound
		}
	}

	if err = this.Put(bookmark.Name, bookmark.Data); err != nil {
		return
	}

	// set upserted bookmark to current
	_, err = BookmarkModel.SetCurrent(bookmark.Name)
	return
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

	if name == BOOKMARK_DEFAULT_NAME {
		return errors.ErrBookmarkEditDefault
	}

	// TODO 暂时允许所有名字
	return nil
}
