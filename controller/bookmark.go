package controller

import (
	"net/http"

	"github.com/jmjoy/http-api-tester/base"
)

type BookmarkController struct {
	*base.Controller
}

func NewBookmarkController(w http.ResponseWriter, r *http.Request) base.Restful {
	return &BookmarkController{
		Controller: base.NewController(w, r),
	}
}

// Get: favicon.ico
func (this *BookmarkController) Get() error {
	return nil
}
