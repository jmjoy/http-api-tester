package base

import (
	"encoding/json"
	"net/http"
)

type Restful interface {
	Get() error
	Post() error
	Put() error
	Delete() error
}

type controllerNewer func(http.ResponseWriter, *http.Request) Restful

func HandleRestful(pattern string, fn controllerNewer) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		rf := fn(w, r)

		var err error
		switch r.Method {
		case "GET":
			err = rf.Get()

		case "POST":
			err = rf.Post()

		case "PUT":
			err = rf.Put()

		case "DELETE":
			err = rf.Delete()

		default:
			err = ErrorMethodNotAllowed
		}

		if err != nil {
			switch err.(type) {
			case *apiStatusError: // api error
				apiStatusErr := err.(*apiStatusError)
				RenderJson(w, apiStatusErr.status, apiStatusErr.message, nil)

			case *statusError: // system error
				statusErr := err.(*statusError)
				LogStatusError(r, statusErr)
				http.Error(w, statusErr.message, statusErr.status)

			default: // unknow
				Log(LOG_LV_FAIL, err)
			}
		}
	})
}

func RenderJson(w http.ResponseWriter, status int, message string, data interface{}) {
	out := map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
	buf, err := json.Marshal(out)
	if err != nil {
		Log(LOG_LV_FAIL, err)
		return
	}
	_, err = w.Write(buf)
	if err != nil {
		Log(LOG_LV_FAIL, err)
		return
	}
}
