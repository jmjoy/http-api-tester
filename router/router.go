package router

import (
	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/controller"
)

// TODO I want to use map[string]IController before but...
var Routers = map[string]func() app.IController{
	"/":            func() app.IController { return new(controller.IndexController) },
	"/favicon.ico": func() app.IController { return new(controller.FaviconController) },
	"/static/":     func() app.IController { return new(controller.StaticController) },
	"/bookmark":    func() app.IController { return new(controller.BookmarkController) },
	"/bookmarks":   func() app.IController { return new(controller.BookmarksController) },
}
