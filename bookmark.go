package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

type BookmarkController struct {
	*Controller
}

func NewBookmarkController() interface{} {
	return &BookmarkController{
		&Controller{},
	}
}

func (this *BookmarkController) Get() {
	querys := this.r.URL.Query()
	key := querys.Get("key")
	if key == "" {
		this.RenderJson(400, "请输入书签的键", nil)
		return
	}

	jsonConfig := GetConfigJson()
	bookmark, has := jsonConfig.Bookmarks[key]
	if !has {
		this.RenderJson(400, "所选书签找不到", nil)
		return
	}

	jsonConfig.Selected = key
	SaveConfigJson(jsonConfig)

	this.RenderJson(200, "", bookmark)
}

func (this *BookmarkController) Post() {
	buf, err := ioutil.ReadAll(this.r.Body)
	if err != nil {
		this.RenderJson(400, "读取输入出错[罕见]", nil)
		return
	}

	// 解析输入JSON
	input := new(Bookmark)
	err = json.Unmarshal(buf, input)
	if err != nil {
		this.RenderJson(40001, "传入参数[JSON]解析出错: "+err.Error(), nil)
		return
	}

	// 检查名字是否重复
	jsonConfig := GetConfigJson()
	for _, row := range jsonConfig.Bookmarks {
		if row.Name == input.Name {
			this.RenderJson(40010, "书签名已经使用过了", nil)
			return
		}
	}

	// 添加书签
	rand := strconv.FormatInt(time.Now().UnixNano(), 10)
	jsonConfig.Bookmarks[rand] = *input

	// 持久化到文件
	err = SaveConfigJson(jsonConfig)
	if err != nil {
		panic(err)
	}

	this.RenderJson(200, "", map[string]string{
		"insertKey":  rand,
		"insertName": input.Name,
	})
}
