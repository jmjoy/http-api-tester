package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"
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
	if err := initDb(cfg.DbPath); err != nil {
		Log(LOG_LV_FAIL, err)
		return
	}

	Log(LOG_LV_INFO, "测试接口服务器在跑了，请访问 http://"+port)
	Log(LOG_LV_FAIL, http.ListenAndServe(port, nil))
}

var controllerPoolMap = make(map[string]*sync.Pool)

func HandleRestful(pattern string, c IController) {
	controllerPoolMap[pattern] = &sync.Pool{
		New: func() interface{} {
			return reflect.New(reflect.TypeOf(c)).Interface()
		},
	}

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c := controllerPoolMap[pattern].Get().(IController)
		defer controllerPoolMap[pattern].Put(c)

		c.SetR(r)
		c.SetW(w)

		var err error
		switch r.Method {
		case "GET":
			err = c.Get()

		case "POST":
			err = c.Post()

		case "PUT":
			err = c.Put()

		case "DELETE":
			err = c.Delete()

		default:
			err = new(ErrorMethodNotAllowed)
		}

		// handle error
		if err != nil {
			switch err.(type) {
			case IApiStatusError: // api error
				apiStatusErr := err.(IApiStatusError)
				status := errorGetStatus(apiStatusErr)
				message := errorGetMessage(apiStatusErr)
				if err := jsonRender(w, status, message, nil); err != nil {
					panic(err)
				}

			case IStatusError: // web or server error
				statusErr := err.(IStatusError)
				status := errorGetStatus(statusErr)
				message := errorGetMessage(statusErr)
				Log(LOG_LV_FAIL, fmt.Sprintf("<%d> %s (%s)", status, message, r.URL))
				http.Error(w, message, status)

			default: // system error
				panic(err)
			}
		}

	})
}

func jsonRender(w http.ResponseWriter, status int, message string, data interface{}) error {
	out := map[string]interface{}{
		"Status":  status,
		"Message": message,
		"Data":    data,
	}
	buf, err := json.Marshal(out)
	if err != nil {
		return err
	}
	_, err = w.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
