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
	Routers map[string]func() IController
}

// Run is a function named Run
func Run(cfg Config) {
	port := cfg.Port
	if !strings.ContainsRune(port, ':') {
		port = "localhost:" + port
	}

	// init db config
	if err := InitDb(cfg.DbPath); err != nil {
		Log(LOG_LV_FAIL, err)
		return
	}

	// register restful router
	for pattren, fn := range cfg.Routers {
		HandleRestful(pattren, fn)
	}

	Log(LOG_LV_INFO, "测试接口服务器在跑了，请访问 http://"+port)
	Log(LOG_LV_FAIL, http.ListenAndServe(port, nil))
}

var controllerPoolMap = make(map[string]*sync.Pool)

func HandleRestful(pattern string, fn func() IController) {
	controllerPoolMap[pattern] = &sync.Pool{
		// New: func() interface{} {
		// TODO Find why can't new a substruct
		// return reflect.New(reflect.TypeOf(c)).Elem().Interface()
		// },
		New: func() interface{} {
			return fn()
		},
	}

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c := controllerPoolMap[pattern].Get().(IController)
		defer controllerPoolMap[pattern].Put(c)

		// ResetController(c, w, r)
		// TODO Here alway new a Controller
		reflect.ValueOf(c).Elem().FieldByName("Controller").
			Set(reflect.ValueOf(&Controller{R: r, W: w}))
		// c.Reset(w, r)

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
			err = ErrMethodNotAllowed
		}

		// handle error
		if err != nil {
			switch err.(type) {
			case *ApiStatusError: // api error
				apiStatusErr := err.(*ApiStatusError)
				status := apiStatusErr.status
				message := apiStatusErr.message
				if err := jsonRender(w, status, message, nil); err != nil {
					panic(err)
				}

			case *StatusError: // web or server error
				statusErr := err.(*StatusError)
				status := statusErr.status
				message := statusErr.message
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
