package model

import (
	"github.com/jmjoy/http-api-tester/base"
	"github.com/jmjoy/http-api-tester/bean"
)

type BookmarkModel struct {
}

func NewBookmarkModel() *BookmarkModel {
	return new(BookmarkModel)
}

func (this *BookmarkModel) GetCurrent() (bean.Data, error) {
	name, err := base.Db.Get("bookmark", "selected")
	if err != nil {
		return this.handleGetError(err)
	}

	data, err := NewBookmarksModel().Get(string(name))
	if err != nil {
		return this.handleGetError(err)
	}

	return data, nil
}

func (this *BookmarkModel) SetCurrent(name string) error {
	bookmarksModel := NewBookmarksModel()
	if err := bookmarksModel.validateBookmarkName(name); err != nil {
		return err
	}
	if _, err := bookmarksModel.Get(name); err != nil {
		return err
	}
	if err := base.Db.Put("bookmark", "selected", []byte(name)); err != nil {
		return err
	}
	return nil
}

func (this *BookmarkModel) handleGetError(err error) (bean.Data, error) {
	if err == base.ErrorBucketNotFound || err == base.ErrorBookmarkNotFound {
		return bean.DataDefault(), nil
	}
	return bean.Data{}, err
}
