package router

import (
	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/controller"
)

var Routers = map[string]app.IController{
	"/":            new(controller.IndexController),
	"/favicon.ico": new(controller.FaviconController),
	"/static/":     new(controller.StaticController),
	"/bookmark":    new(controller.BookmarkController),
	"/bookmarks":   new(controller.BookmarksController),
}
