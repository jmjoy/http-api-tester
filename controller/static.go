package controller

import (
	"path"

	"github.com/jmjoy/http-api-tester/app"
	"github.com/jmjoy/http-api-tester/text"
)

var extMimeMap = map[string]string{
	".css": "text/css;charset=utf-8",
	".js":  "application/x-javascript",
	".map": "text/map;charset=utf-8",
}

type StaticController struct {
	*app.Controller
}

func (this *StaticController) Get() error {
	uS := this.R.URL.String()[1:] // remove fst "/"
	ext := path.Ext(uS)

	var buf []byte
	var cType string

	if mimeType, has := extMimeMap[ext]; has {
		cType = mimeType
		content := text.ProvideString(uS)
		if content == "" {
			return app.ErrNotFound
		}
		buf = []byte(content)

	} else { // is not textual file
		buf = text.ProvideBytes(uS)
		if buf == nil {
			return app.ErrNotFound
		}
	}

	this.W.Header().Set("Content-Type", cType)
	this.W.Write(buf)
	return nil
}
