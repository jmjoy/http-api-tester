package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/jmjoy/http-api-tester/base"
	"github.com/jmjoy/http-api-tester/model"
)

type BookmarkController struct {
	*base.Controller

	model *model.BookmarkModel
}

func NewBookmarkController(w http.ResponseWriter, r *http.Request) base.Restful {
	return &BookmarkController{
		Controller: base.NewController(w, r),
		model:      model.NewBookmarkModel(),
	}
}

// Get: Get current bookmark
func (this *BookmarkController) Get() error {
	data, err := this.model.GetCurrent()
	if err != nil {
		return base.NewApiStatusError(4000, err)
	}

	return this.RenderJson(data)
}

// Post: Set current bookmark
func (this *BookmarkController) Post() error {
	fmt.Println("Setting current bookmark...")

	var name string

	// Get Body
	buf, err := ioutil.ReadAll(this.R().Body)
	if err != nil {
		return base.NewApiStatusError(4000, fmt.Errorf("Read body error: %s", err))
	}

	if len(buf) != 0 {
		json, err := simplejson.NewJson(buf)
		if err != nil {
			return base.NewApiStatusError(4000, fmt.Errorf("Parse body error: %s", err))
		}
		name, _ = json.Get("name").String()
	}

	if err = this.model.SetCurrent(name); err != nil {
		return base.NewApiStatusError(4000, err)
	}

	return this.RenderJson(nil)
}
