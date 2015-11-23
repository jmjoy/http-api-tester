package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"text/template"
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
	http.HandleFunc("/static/", handleStatic)
	HandleRestful("/bookmark", NewBookmarkController)
	HandleRestful("/plugin", NewPluginController)
	HandleRestful("/submit", NewSubmitController)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.html").Parse(text["view/index.html"])
	if err != nil {
		panic(err)
	}
	t.Execute(w, map[string]string{
		"Config": GetConfigJsonString(),
	})
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	favicon, err := base64.StdEncoding.DecodeString(text["favicon.ico"])
	if err != nil {
		fmt.Println("[ERROR] favicon.ico")
		return
	}

	w.Write(favicon)
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	uS := r.URL.String()[1:] // remove fst "/"
	content, has := text[uS]
	if !has {
		fmt.Println(uS)
		return
	}

	ext := path.Ext(uS)
	var cType string
	switch ext {
	case ".css":
		cType = "text/css;charset=utf-8"
	case ".js":
		cType = "application/x-javascript"
	default:
		cType = "text/html;charset=utf-8"
	}
	w.Header().Set("Content-Type", cType)
	io.WriteString(w, content)
}
