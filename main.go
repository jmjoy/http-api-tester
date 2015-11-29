package main

import (
	"flag"

	"github.com/jmjoy/http-api-tester/app"
)

const (
	NAME    = "http-api-tester"
	VERSION = "0.5"
)

var (
	gPort   string
	gDbPath string
)

func init() {
	flag.StringVar(&gPort, "p", "8080", "服务器运行端口")
	flag.StringVar(&gDbPath, "db", "http-api-tester.db", "数据库路径")
}

func main() {
	flag.Parse()

	app.Run(app.Config{
		Port:   gPort,
		DbPath: gDbPath,
	})
}
