package text

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	BasePath string
	IsDebug  bool
)

func ProvideString(path string) string {
	if IsDebug {
		if !isFileExists(path) {
			return ""
		}
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
		if !isFileExists(path) {
			return nil
		}
		buf, err := ioutil.ReadFile(filepath.Join(BasePath, path))
		if err != nil {
			panic(err)
		}
		return buf
	}

	content, has := Text[path]
	if !has {
		return nil
	}
	buf, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		panic(err)
	}
	return buf
}

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
