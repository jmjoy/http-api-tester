package router

import (
	"github.com/jmjoy/http-api-tester/base"
	"github.com/jmjoy/http-api-tester/controller"
)

func Router() {
	base.HandleRestful("/", controller.NewIndexController)
	base.HandleRestful("/favicon.ico", controller.NewFaviconController)
	base.HandleRestful("/static/", controller.NewStaticController)
	base.HandleRestful("/bookmark/", controller.NewBookmarkController)
}
