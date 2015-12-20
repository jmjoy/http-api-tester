package model

import (
	"github.com/jmjoy/http-api-tester/app"
)

var BookmarkModel = &bookmarkModel{
	Model:    app.NewModel("bookmark"),
	selected: "selected",
}

type bookmarkModel struct {
	*app.Model
	selected string
}

func (this *bookmarkModel) GetCurrent() (data Data, err error) {
	var name string
	has, err := this.Get(this.selected, &name)
	if err != nil {
		return
	}
	if !has {
		return DataDefault(), nil
	}

	return BookmarksModel.Get(name)
}

func (this *bookmarkModel) SetCurrent(name string) (err error) {
	if err = BookmarksModel.validateBookmarkName(name); err != nil {
		return
	}

	_, err = BookmarksModel.Get(name)
	if err != nil {
		return
	}

	return this.Put(this.selected, name)
}
