package text

import (
	"encoding/base64"
	"io/ioutil"
	"path/filepath"
)

var (
	BasePath string
	IsDebug  bool
)

func ProvideString(path string) string {
	if IsDebug {
		buf, err := ioutil.ReadFile(filepath.Join(BasePath, path))
		if err != nil {
			panic(err)
		}
		return string(buf)
	}

	return Text[path]
}

func ProvideBytes(path string) []byte {
	if IsDebug {
		buf, err := ioutil.ReadFile(filepath.Join(BasePath, path))
		if err != nil {
			panic(err)
		}
		return buf
	}

	buf, err := base64.StdEncoding.DecodeString(Text[path])
	if err != nil {
		panic(err)
	}
	return buf
}
