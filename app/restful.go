package app

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("Regist router:", pattern)
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Handle request:", r.Method, r.URL)

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
