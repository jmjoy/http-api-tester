package app

import (
	"net/http"
	"strings"
)

// Config is a struct named Config
type Config struct {
	// User defined
	Port   string
	DbPath string

	// App need
	Routers []map[string]IController
}

// Run is a function named Run
func Run(cfg Config) {
	port := cfg.Port
	if !strings.ContainsRune(port, ':') {
		port = "localhost:" + port
	}

	// init db config
	if err := base.Db(cfg.DbPath); err != nil {
		Log(base.LOG_LV_FAIL, err)
		return
	}

	Log(base.LOG_LV_INFO, "测试接口服务器在跑了，请访问 http://"+port)
	Log(base.LOG_LV_FAIL, http.ListenAndServe(port, nil))
}
