package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"

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
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.HandleFunc("/static/", handleStatic)
	HandleRestful("/bookmark", NewBookmarkController)
	HandleRestful("/plugin", NewPluginController)
	HandleRestful("/submit", NewSubmitController)
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	favicon, err := base64.StdEncoding.DecodeString(text["favicon.ico"])
	if err != nil {
		fmt.Println("[ERROR] favicon.ico")
		return
	}

	w.Write(favicon)
}

var staticExtMimeMap = map[string]string{
	".css": "text/css;charset=utf-8",
	".js":  "application/x-javascript",
	".map": "text/map;charset=utf-8",
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	uS := r.URL.String()[1:] // remove fst "/"
	content, has := text[uS]
	if !has {
		fmt.Println(uS)
		return
	}

	ext := path.Ext(uS)
	var buf []byte
	var cType string
	mimeType, has := staticExtMimeMap[ext]
	if has {
		cType = mimeType
		buf = []byte(content)
	} else { // is not textual file
		cType = ""
		var err error
		buf, err = base64.StdEncoding.DecodeString(content)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	w.Header().Set("Content-Type", cType)
	w.Write(buf)
}
