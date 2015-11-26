package controller

import (
	"io"
	"net/http"

	"github.com/jmjoy/http-api-tester/base"
	"github.com/jmjoy/http-api-tester/text"
)

type IndexController struct {
	*base.Controller
}

func NewIndexController(w http.ResponseWriter, r *http.Request) base.Restful {
	return &IndexController{
		Controller: base.NewController(w, r),
	}
}

// Get: index page
func (this *IndexController) Get() error {
	_, err := io.WriteString(this.W(), text.Text["view/index.html"])
	if err != nil {
		return base.NewStatusErrorFromError(http.StatusInternalServerError, err)
	}
	return nil
}

// Post: submit
func (this *IndexController) Post() error {
	return nil
}
