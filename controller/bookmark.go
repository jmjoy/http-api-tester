package controller

import (
	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/model"
)

type BookmarkController struct {
	*app.Controller
}

// Get: Get current bookmark
func (this *BookmarkController) Get() (err error) {
	data, err := model.BookmarkModel.GetCurrent()
	if err != nil {
		return
	}

	return this.JsonSuccess(data)
}

// Post: Set current bookmark
func (this *BookmarkController) Post() (err error) {
	var data map[string]string
	if err = this.ParseJsonBody(&data); err != nil {
		return
	}

	if err = model.BookmarkModel.SetCurrent(data["name"]); err != nil {
		return
	}

	return this.JsonSuccess(nil)
}
