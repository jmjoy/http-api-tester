package main

import (
	"fmt"
	"io/ioutil"
)

type BookmarkController struct {
	*Controller
}

func NewBookmarkController() *BookmarkController {
	return &BookmarkController{
		&Controller{},
	}
}

func (this *BookmarkController) Post() {
	buf, err := ioutil.ReadAll(this.r.Body)
	if err != nil {
		this.RenderJson(400, "传入参数[JSON]解析出错", nil)
		return
	}

	// TODO Add bookmark
	fmt.Println(string(buf))

	this.RenderJson(400, "", nil)
}
