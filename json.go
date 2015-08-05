package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type JsonConfig struct {
	Selected  string
	Bookmarks map[string]Bookmark
	Plugins   map[string]Plugin
}

type Bookmark struct {
	Name, Method, Url string
	Args              []Arg
	Bm                Bm
	Plugin            BookmarkPlugin
}

type Arg struct {
	Key, Value, Method string
}

type Bm struct {
	Switch bool
	N, C   uint
}

type BookmarkPlugin struct {
	Key  string
	Data map[string]string
}

type Plugin struct {
	Name   string
	Fields map[string]string
}

var gConfigJsonMutex = new(sync.RWMutex)

func GetConfigJson() *JsonConfig {
	gConfigJsonMutex.Lock()
	defer gConfigJsonMutex.Unlock()

	src, err := ioutil.ReadFile(gConfigPath)
	if err != nil {
		panic(err)
	}

	jsonConfig := new(JsonConfig)
	err = json.Unmarshal(src, jsonConfig)
	if err != nil {
		panic(err)
	}

	return jsonConfig
}

func GetConfigJsonString() string {
	gConfigJsonMutex.Lock()
	defer gConfigJsonMutex.Unlock()

	src, err := ioutil.ReadFile(gConfigPath)
	if err != nil {
		panic(err)
	}

	dst := new(bytes.Buffer)
	err = json.Compact(dst, src)
	if err != nil {
		panic(err)
	}

	return dst.String()
}

func SaveConfigJson(jsonConfig *JsonConfig) error {
	buf, err := json.Marshal(jsonConfig)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	err = json.Indent(buffer, buf, "", "    ")
	if err != nil {
		return err
	}

	gConfigJsonMutex.Lock()
	defer gConfigJsonMutex.Unlock()

	fw, err := os.Open(gConfigPath)
	if err != nil {
		return err
	}
	defer fw.Close()

	_, err = buffer.WriteTo(fw)
	if err != nil {
		return err
	}

	return nil
}
