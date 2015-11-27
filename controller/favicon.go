package controller

import (
	"encoding/base64"
	"net/http"

	"github.com/jmjoy/http-api-tester/base"
	"github.com/jmjoy/http-api-tester/text"
)

type FaviconController struct {
	*base.Controller
}

func NewFaviconController(w http.ResponseWriter, r *http.Request) base.Restful {
	return &FaviconController{
		Controller: base.NewController(w, r),
	}
}

// Get: favicon.ico
func (this *FaviconController) Get() error {
	favicon, err := base64.StdEncoding.DecodeString(text.Text["favicon.ico"])
	if err != nil {
		return base.NewStatusError(http.StatusInternalServerError, err)
	}

	_, err = this.W().Write(favicon)
	if err != nil {
		return base.NewStatusError(http.StatusInternalServerError, err)
	}

	return nil
}
