package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jmjoy/http-api-tester/base"
	"github.com/jmjoy/http-api-tester/model"
)

type BookmarksController struct {
	*base.Controller

	model *model.BookmarksModel
}

func NewBookmarksController(w http.ResponseWriter, r *http.Request) base.Restful {
	return &BookmarksController{
		Controller: base.NewController(w, r),
		model:      model.NewBookmarksModel(),
	}
}

// Get: get bookmark config by name or current
func (this *BookmarksController) Get() error {
	name := this.R().URL.Query().Get("name")

	data, err := this.model.Get(name)
	if err != nil {
		return base.NewApiStatusError(4000, err)
	}

	return this.RenderJson(data)
}

// Post: add bookmark config
func (this *BookmarksController) Post() error {
	bookmark, err := this.parseBookmarkFromBody()
	if err != nil {
		return err
	}

	// 添加书签
	if err = this.model.Upsert(bookmark, model.UPSERT_ADD); err != nil {
		return base.NewApiStatusError(4000, err)
	}

	return this.RenderJson(nil)
}

// Put: update bookmark config
func (this *BookmarksController) Put() error {
	bookmark, err := this.parseBookmarkFromBody()
	if err != nil {
		return base.NewApiStatusError(4000, err)
	}

	// 修改书签
	if err = this.model.Upsert(bookmark, model.UPSERT_UPDATE); err != nil {
		return base.NewApiStatusError(4000, err)
	}

	return this.RenderJson(nil)
}

// Delete: delete bookmark
func (this *BookmarksController) Delete() error {
	name := this.R().URL.Query().Get("name")
	if err := this.model.Delete(name); err != nil {
		return base.NewApiStatusError(4000, err)
	}
	return this.RenderJson(nil)
}

// for Post and Put: upsert data
func (this *BookmarksController) parseBookmarkFromBody() (model.Bookmark, error) {
	// Get Body
	buf, err := ioutil.ReadAll(this.R().Body)
	if err != nil {
		return model.Bookmark{}, base.NewApiStatusError(4000, fmt.Errorf("Read body error: %s", err))
	}

	// 解析输入JSON
	var bookmark model.Bookmark
	if err = json.Unmarshal(buf, &bookmark); err != nil {
		return model.Bookmark{}, base.NewApiStatusError(4000, fmt.Errorf("Unmarshal body error: %s", err))
	}

	return bookmark, nil
}
