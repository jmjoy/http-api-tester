package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"sort"
	"strings"
)

type PluginController struct {
	*Controller
}

func NewPluginController() interface{} {
	return &PluginController{
		&Controller{},
	}
}

func (this *PluginController) Get() {
	this.RenderJson(200, "", GetConfigJson())
}

type PluginFunc func(BookmarkPlugin, *RequestStruct) (*RequestStruct, error)

var pluginHandlerPool = make(map[string]PluginFunc)

func HookPlugin(bp BookmarkPlugin, reqS *RequestStruct) (*RequestStruct, error) {
	fn, has := pluginHandlerPool[bp.Key]
	if !has {
		return nil, errors.New("No this plugin")
	}

	return fn(bp, reqS)
}

func init() {
	pluginHandlerPool["notuse"] = func(bp BookmarkPlugin, reqS *RequestStruct) (*RequestStruct, error) {
		return reqS, nil
	}

	pluginHandlerPool["md5signature"] = func(bp BookmarkPlugin, reqS *RequestStruct) (*RequestStruct, error) {
		keyName, has := bp.Data["keyname"]
		if !has {
			return nil, errors.New("Plugin keyname doesn't exist!")
		}
		password, has := bp.Data["password"]
		if !has {
			return nil, errors.New("Plugin password doesn't exist!")
		}

		argsM := make(map[string]string)
		query := reqS.URL.Query()

		for key, args := range query {
			argsM[key] = args[0]
		}

		if reqS.Method == "POST" {
			for key, args := range reqS.PostData {
				argsM[key] = args[0]
			}
		}

		argsKeys := make([]string, 0, len(argsM))
		for k := range argsM {
			argsKeys = append(argsKeys, k)
		}
		sort.Strings(argsKeys)

		values := make([]string, 0, 4)
		for _, key := range argsKeys {
			values = append(values, argsM[key])
		}
		values = append(values, password)

		text := strings.Join(values, "")
		md5text := fmt.Sprintf("%x", md5.Sum([]byte(text)))

		query.Add(keyName, md5text)
		reqS.URL.RawQuery = query.Encode()

		return reqS, nil
	}

	pluginHandlerPool["md5signature-noquery"] = func(bp BookmarkPlugin, reqS *RequestStruct) (*RequestStruct, error) {
		keyName, has := bp.Data["keyname"]
		if !has {
			return nil, errors.New("Plugin keyname doesn't exist!")
		}
		password, has := bp.Data["password"]
		if !has {
			return nil, errors.New("Plugin password doesn't exist!")
		}

		argsM := make(map[string]string)
		query := reqS.URL.Query()

		for key, args := range query {
			argsM[key] = args[0]
		}

		if reqS.Method == "POST" {
			for key, args := range reqS.PostData {
				argsM[key] = args[0]
			}
		}

		argsKeys := make([]string, 0, len(argsM))
		for k := range argsM {
			argsKeys = append(argsKeys, k)
		}
		sort.Strings(argsKeys)

		values := make([]string, 0, 4)
		for _, key := range argsKeys {
			values = append(values, argsM[key])
		}
		values = append(values, password)

		text := strings.Join(values, "")
		md5text := fmt.Sprintf("%x", md5.Sum([]byte(text)))

		query.Add(keyName, md5text)
		reqS.URL.RawQuery = query.Encode()

		return reqS, nil
	}

}
