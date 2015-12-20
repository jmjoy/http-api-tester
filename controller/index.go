package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jmjoy/http-api-tester/bean"
	"github.com/jmjoy/http-api-tester/model"
	"github.com/jmjoy/http-api-tester/text"
)

type IndexController struct {
	*base.Controller
}

func NewIndexController(w http.ResponseWriter, r *http.Request) base.Restful {
	return &IndexController{
		Controller: base.NewController(w, r),
	}
}

// Get:
func (this *IndexController) Get() error {
	act := this.R().URL.Query().Get("act")
	if act == "initData" {
		return this.initData()
	}
	return this.indexPage()
}

func (this *IndexController) indexPage() error {
	_, err := io.WriteString(this.W(), text.Text["view/index.html"])
	if err != nil {
		return base.NewStatusError(http.StatusInternalServerError, err)
	}
	return nil
}

func (this *IndexController) initData() error {
	data, err := model.NewBookmarkModel().GetCurrent()
	if err != nil {
		return base.NewApiStatusError(4000, err)
	}

	renderData := map[string]interface{}{
		"Data":    data,
		"Plugins": bean.PluginPool,
	}
	return this.RenderJson(renderData)
}

// Post: submit
func (this *IndexController) Post() error {
	data, err := this.parseDataFromBody()
	if err != nil {
		return err
	}
	response, err := model.NewSubmitModel().Submit(data)
	if err != nil {
		return base.NewApiStatusError(4000, err)
	}
	return this.RenderJson(response)
}

// get Data from body <JSON format>
func (this *IndexController) parseDataFromBody() (bean.Data, error) {
	// Get Body
	buf, err := ioutil.ReadAll(this.R().Body)
	if err != nil {
		return bean.Data{}, base.NewApiStatusError(4000, fmt.Errorf("Read body error: %s", err))
	}

	// 解析输入JSON
	var data bean.Data
	if len(buf) != 0 {
		if err = json.Unmarshal(buf, &data); err != nil {
			return bean.Data{}, base.NewApiStatusError(4000, fmt.Errorf("Unmarshal body error: %s", err))
		}
	}

	return data, nil
}
