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
	gPort       string
	gConfigPath string
	gViewDir    string
	gStaticDir  string
)

func init() {
	flag.StringVar(&gPort, "p", "8080", "服务器运行端口")
	flag.StringVar(&gConfigPath, "config", "config.json", "JSON配置文件路径")
	flag.StringVar(&gViewDir, "view", "view", "视图文件所在文件夹")
	flag.StringVar(&gStaticDir, "static", "static", "静态文件所在文件夹")
}

func main() {
	flag.Parse()

	app.Run(app.Config{
		Port:   gPort,
		DbPath: "http-api-tester.db",
	})
}
