package controller

import (
	"io"
	"net/http"

	"github.com/jmjoy/http-api-tester/base"
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

func (this *IndexController) Get() error {
	io.WriteString(this.W(), text.Text["view/index.html"])
	return nil
}
