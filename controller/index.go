package controller

import (
	"io"

	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/model"
	"github.com/jmjoy/http-api-tester/text"
)

type IndexController struct {
	*app.Controller
}

// Get:
func (this *IndexController) Get() error {
	switch this.QueryGet("act") {
	case "initData": // api for get init data
		return this.initData()

	default: // index page
		return this.indexPage()
	}
}

func (this *IndexController) indexPage() (err error) {
	_, err = io.WriteString(this.W, text.ProvideString("view/index.html"))
	return
}

func (this *IndexController) initData() (err error) {
	bookmark, err := model.BookmarkModel.GetCurrent()
	if err != nil {
		return
	}

	return this.JsonSuccess(map[string]interface{}{
		"Bookmark": bookmark,
		"Plugins":  model.PluginPool(),
	})
}

// Post: Submit
func (this *IndexController) Post() (err error) {
	var data model.Data
	if err = this.ParseJsonBody(&data); err != nil {
		return
	}

	resp, err := model.SubmitModel.Submit(data)
	if err != nil {
		return
	}
	return this.JsonSuccess(resp)
}
