package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

// Get: get bookmark config by name or current
func (this *BookmarkController) Get() error {
	var data model.Data
	var err error

	name := this.R().URL.Query().Get("name")
	if name == "" { // get current bookmark if name is empty
		data, err = this.model.GetCurrent()
	} else {
		data, err = this.model.Get(name)
	}

	if err != nil {
		return base.NewApiStatusError(4000, err)
	}

	return this.RenderJson(data)
}

// Post: add bookmark config
func (this *BookmarkController) Post() error {
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
func (this *BookmarkController) Put() error {
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
func (this *BookmarkController) Delete() error {
	name := this.R().URL.Query().Get("name")
	if err := this.model.Delete(name); err != nil {
		return base.NewApiStatusError(4000, err)
	}
	return this.RenderJson(nil)
}

// for Post and Put: upsert data
func (this *BookmarkController) parseBookmarkFromBody() (model.Bookmark, error) {
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

//func (this *BookmarkController) Put() error {
//    buf, err := ioutil.ReadAll(this.r.Body)
//    if err != nil {
//        this.RenderJson(400, "读取输入出错[罕见]", nil)
//        return
//    }

//    // 解析输入JSON
//    input := new(Bookmark)
//    err = json.Unmarshal(buf, input)
//    if err != nil {
//        this.RenderJson(40001, "传入参数[JSON]解析出错: "+err.Error(), nil)
//        return
//    }

//    // 检查名字是否存在
//    jsonConfig := GetConfigJson()
//    names := make(map[string]string)
//    for key, row := range jsonConfig.Bookmarks {
//        names[row.Name] = key
//    }
//    editKey, ok := names[input.Name]
//    if !ok {
//        this.RenderJson(40010, "书签名不存在", nil)
//        return
//    }

//    // 修改书签
//    jsonConfig.Bookmarks[editKey] = *input

//    // 持久化到文件
//    err = SaveConfigJson(jsonConfig)
//    if err != nil {
//        panic(err)
//    }

//    this.RenderJson(200, "", nil)
//}

//func (this *BookmarkController) Delete() error {
//    querys := this.r.URL.Query()
//    key := querys.Get("key")
//    if key == "" {
//        this.RenderJson(400, "请输入书签的键", nil)
//        return
//    }

//    jsonConfig := GetConfigJson()
//    _, has := jsonConfig.Bookmarks[key]
//    if !has {
//        this.RenderJson(400, "所选书签找不到", nil)
//        return
//    }
//    if jsonConfig.Selected == key {
//        this.RenderJson(400, "不能删除使用中的书签", nil)
//        return
//    }

//    // 删除
//    delete(jsonConfig.Bookmarks, key)

//    SaveConfigJson(jsonConfig)

//    this.RenderJson(200, "", nil)
//}
