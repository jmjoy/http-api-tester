package controller

import (
	"encoding/base64"
	"net/http"
	"path"

	"github.com/jmjoy/http-api-tester/base"
	"github.com/jmjoy/http-api-tester/text"
)

type StaticController struct {
	*base.Controller
	extMimeMap map[string]string
}

func NewStaticController(w http.ResponseWriter, r *http.Request) base.Restful {
	extMimeMap := map[string]string{
		".css": "text/css;charset=utf-8",
		".js":  "application/x-javascript",
		".map": "text/map;charset=utf-8",
	}

	return &StaticController{
		Controller: base.NewController(w, r),
		extMimeMap: extMimeMap,
	}
}

func (this *StaticController) Get() error {
	uS := this.R().URL.String()[1:] // remove fst "/"
	content, has := text.Text[uS]
	if !has {
		return base.ErrorNotFound
	}

	ext := path.Ext(uS)
	var buf []byte
	var cType string
	mimeType, has := this.extMimeMap[ext]
	if has {
		cType = mimeType
		buf = []byte(content)
	} else { // is not textual file
		cType = ""
		var err error
		buf, err = base64.StdEncoding.DecodeString(content)
		if err != nil {
			return base.NewStatusErrorFromError(http.StatusInternalServerError, err)
		}
	}

	this.W().Header().Set("Content-Type", cType)
	this.W().Write(buf)
	return nil
}
