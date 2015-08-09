package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

const VERSION = "0.4"

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
	flag.Parse()
}

func main() {
	route()

	log.Printf("测试接口服务器在跑了，请访问 http://localhost:%d\n", gPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", gPort), nil))
}

func route() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	HandleRestful("/bookmark", NewBookmarkController)
	HandleRestful("/plugin", NewPluginController)
	HandleRestful("/submit", NewSubmitController)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.html").ParseFiles(filepath.Join(gViewDir, "index.html"))
	if err != nil {
		panic(err)
	}
	t.Execute(w, map[string]string{
		"Config": GetConfigJsonString(),
	})
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("favicon.ico")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(w, f)
}
