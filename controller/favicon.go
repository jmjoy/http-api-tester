package controller

import (
	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/text"
)

type FaviconController struct {
	*app.Controller
}

// Get: favicon.ico
func (this *FaviconController) Get() (err error) {
	favicon := text.ProvideBytes("favicon.ico")
	_, err = this.W.Write(favicon)
	return
}
