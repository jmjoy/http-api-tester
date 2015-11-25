package base

import (
	"encoding/json"
	"net/http"
)

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
				RenderApiError(w, err.(*apiStatusError))

			case *statusError: // system error
				statusErr := err.(*statusError)
				LogStatusError(r, statusErr)
				http.Error(w, statusErr.message, statusErr.code)
			}
		}
	})
}

func RenderApiError(w http.ResponseWriter, apiStatusErr *apiStatusError) {
	out := map[string]interface{}{
		"status":  apiStatusErr.code,
		"message": apiStatusErr.message,
		"data":    nil,
	}
	buf, err := json.Marshal(out)
	if err != nil {
		LogError(err)
		return
	}
	w.Write(buf)
}
