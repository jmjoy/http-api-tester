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

func (this *bookmarkModel) GetCurrentKey() (name string, has bool, err error) {
	has, err = this.Get(this.selected, &name)
	return
}

func (this *bookmarkModel) GetCurrent() (bookmark Bookmark, err error) {
	name, has, err := this.GetCurrentKey()
	if err != nil {
		return
	}
	if !has {
		return Bookmark{
			Name: BOOKMARK_DEFAULT_NAME,
			Data: DataDefault(),
		}, nil
	}

	data, err := BookmarksModel.Get(name)
	bookmark = Bookmark{
		Name: name,
		Data: data,
	}
	return
}

func (this *bookmarkModel) SetCurrent(name string) (data Data, err error) {
	if data, err = BookmarksModel.Get(name); err != nil {
		return
	}

	if err = this.Put(this.selected, name); err != nil {
		return
	}

	return
}

func (this *bookmarkModel) DeleteCurrent() (err error) {
	return this.Delete(this.selected)
}
