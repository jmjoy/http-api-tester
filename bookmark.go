package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"time"
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
		this.RenderJson(400, "读取输入出错[罕见]", nil)
		return
	}

	// 解析输入JSON
	var input map[string]interface{}
	err = json.Unmarshal(buf, input)
	if err != nil {
		this.RenderJson(40001, "传入参数[JSON]解析出错: "+err.Error(), nil)
		return
	}

	// 获取输入书签名字
	iname, ok := input["name"]
	if !ok {
		this.RenderJson(40002, "name字段缺失", nil)
		return
	}
	name, ok := iname.(string)
	if !ok {
		this.RenderJson(40002, "name字段类型不正确", nil)
		return
	}

	// 检查名字是否重复
	for _, row := range gJsonConfig.Bookmarks {
		in, ok := row["name"]
		if !ok {
			panic("配置文件有问题： Bookmarks的name")
		}
		n := in.(string)
		if n == name {
			this.RenderJson(40010, "书签名已经使用过了", nil)
			return
		}
	}

	// 添加书签
	input["name"] = name
	rand := strconv.FormatInt(time.Now().UnixNano(), 10)
	gJsonConfig.Bookmarks[rand] = input

	////////////
	buf, err = json.Marshal(gJsonConfig)
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	err = json.Indent(buffer, buf, "", "    ")
	if err != nil {
		panic(err)
	}
	buffer.WriteTo(os.Stdout)
	///////////////

	this.RenderJson(400, "", nil)
}
