package main

import (
	"flag"

	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/router"
	"github.com/jmjoy/http-api-tester/text"
)

const (
	NAME    = "http-api-tester"
	VERSION = "0.5"

	IS_DEBUG = true
)

var (
	gPort   string
	gDbPath string
)

func init() {
	flag.StringVar(&gPort, "p", "8080", "服务器运行端口")
	flag.StringVar(&gDbPath, "db", NAME+".db", "数据库路径")
}

func main() {
	flag.Parse()

	text.IsDebug = IS_DEBUG
	text.BasePath = "."

	app.Run(app.Config{
		Port:    gPort,
		DbPath:  gDbPath,
		Routers: router.Routers,
	})
}
