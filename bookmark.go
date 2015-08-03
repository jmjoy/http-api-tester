package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type BookmarkController struct {
	*Controller
}

func (this *BookmarkController) Post(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RenderJson(w, 400, "传入参数[JSON]解析出错", nil)
		return
	}

	fmt.Println(string(buf))

	RenderJson(w, 400, "", nil)
}
