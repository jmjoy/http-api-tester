package app

import (
	"net/http"
	"strings"

	"github.com/jmjoy/http-api-tester/base"
	"github.com/jmjoy/http-api-tester/router"
)

type Config struct {
	Port   string
	DbPath string
}

func Run(cfg Config) {
	port := cfg.Port
	if !strings.ContainsRune(port, ':') {
		port = "localhost:" + port
	}

	router.Router()

	// init db config
	if err := base.Db(cfg.DbPath); err != nil {
		base.Log(base.LOG_LV_FAIL, err)
		return
	}

	base.Log(base.LOG_LV_INFO, "测试接口服务器在跑了，请访问 http://"+port)
	base.Log(base.LOG_LV_FAIL, http.ListenAndServe(port, nil))
}
