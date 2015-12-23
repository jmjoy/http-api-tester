package controller

import (
	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/model"
)

type BookmarksController struct {
	*app.Controller
}

// Get: get bookmark config by name or current
func (this *BookmarksController) Get() error {
	name := this.QueryGet("name")

	data, err := model.BookmarksModel.Get(name)
	if err != nil {
		return err
	}

	return this.JsonSuccess(data)
}

// Post: add bookmark config
func (this *BookmarksController) Post() error {
	return this.Upsert(model.UPSERT_ADD)
}

// Put: update bookmark config
func (this *BookmarksController) Put() error {
	return this.Upsert(model.UPSERT_UPDATE)
}

func (this *BookmarksController) Upsert(typ model.UpsertType) (err error) {
	var bookmark model.Bookmark

	if err = this.ParseJsonBody(&bookmark); err != nil {
		return
	}

	if err = model.BookmarksModel.Upsert(bookmark, typ); err != nil {
		return
	}

	return this.JsonSuccess(nil)
}

// Delete: delete bookmark
func (this *BookmarksController) Delete() (err error) {
	name := this.QueryGet("name")
	if err = model.BookmarksModel.Delete(name); err != nil {
		return
	}
	return this.JsonSuccess(nil)
}
