package controller

import (
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

// Get current bookmark config
func (this *BookmarkController) Get() error {
	data, err := this.model.GetCurrent()
	if err != nil {
		return base.NewApiStatusErrorFromError(4000, err)
	}

	return this.RenderJson(data)
}

//func (this *BookmarkController) Get() error {
//    querys := this.R().URL.Query()
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

//    jsonConfig.Selected = key
//    SaveConfigJson(jsonConfig)

//    this.RenderJson(200, "", jsonConfig)

//    return nil
//}

//func (this *BookmarkController) Post() error {

//    buf, err := ioutil.ReadAll(this.R().Body)
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

//    // 检查名字是否重复
//    jsonConfig := GetConfigJson()
//    for _, row := range jsonConfig.Bookmarks {
//        if row.Name == input.Name {
//            this.RenderJson(40010, "书签名已经使用过了", nil)
//            return
//        }
//    }

//    // 添加书签
//    rand := strconv.FormatInt(time.Now().UnixNano(), 10)
//    jsonConfig.Bookmarks[rand] = *input

//    // 持久化到文件
//    err = SaveConfigJson(jsonConfig)
//    if err != nil {
//        panic(err)
//    }

//    this.RenderJson(200, "", map[string]string{
//        "insertKey":  rand,
//        "insertName": input.Name,
//    })
//}

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
