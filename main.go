package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jmjoy/http-api-tester/router"
)

const VERSION = "0.5"

var (
	gPort       int
	gConfigPath string
	gViewDir    string
	gStaticDir  string
)

func init() {
	flag.IntVar(&gPort, "p", 8080, "服务器运行端口")
	flag.StringVar(&gConfigPath, "config", "config.json", "JSON配置文件路径")
	flag.StringVar(&gViewDir, "view", "view", "视图文件所在文件夹")
	flag.StringVar(&gStaticDir, "static", "static", "静态文件所在文件夹")
}

func main() {
	flag.Parse()

	router.Router()
	//route()

	log.Printf("测试接口服务器在跑了，请访问 http://localhost:%d\n", gPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", gPort), nil))
}

func route() {
	//http.HandleFunc("/", handleIndex)
	//http.HandleFunc("/favicon.ico", handleFavicon)
	//http.HandleFunc("/static/", handleStatic)
	//HandleRestful("/bookmark", NewBookmarkController)
	//HandleRestful("/plugin", NewPluginController)
	//HandleRestful("/submit", NewSubmitController)
}
