package router

import (
	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/controller"
)

var Routers = map[string]app.IController{
	"/":            controller.IndexController,
	"/favicon.ico": controller.FaviconController,
	"/static/":     controller.StaticController,
	"/bookmark":    controller.BookmarkController,
	"/bookmarks":   controller.BookmarksController,
}
