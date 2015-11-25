package controller

import (
	"net/http"

	"github.com/jmjoy/http-api-tester/base"
)

type IndexController struct {
	*base.Controller
}

func NewIndexController(w http.ResponseWriter, r *http.Request) *IndexController {
	return &IndexController{
		Controller: base.NewController(w, r),
	}
}

func (this *IndexController) Get() error {
	return base.NewApiStatusError(450, "fuck")
}
