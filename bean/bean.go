package bean

import (
	"errors"
	"net/url"
)

// bookmarks
type BookmarkMap map[string]Data

// named bookmark
type Bookmark struct {
	Name string
	Data Data
}

// Submit Data
type Data struct {
	Method string
	Url    string
	Args   []Arg
	Bm     Bm
	Plugin Plugin
}

// Submit Arg
type Arg struct {
	Key    string
	Value  string
	Method string
}

// Benchmark data
type Bm struct {
	Switch bool
	N      uint
	C      uint
}

type Plugin struct {
	Key  string
	Data map[string]string
}

func DataDefault() Data {
	return Data{
		Method: "GET",
		Args:   []Arg{},
		Plugin: Plugin{
			Data: map[string]string{},
		},
	}
}

func (this Data) Validate() error {
	if this.Url == "" {
		return errors.New("请指定URL")
	}

	u, err := url.Parse(this.Url)
	if err != nil {
		return err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("未知协议：" + u.Scheme)
	}

	if u.Host == "" {
		return errors.New("请指定host")
	}

	for _, arg := range this.Args {
		switch arg.Method {
		case "GET", "POST":
		default:
			return errors.New("参数中包含未知请求方式: " + arg.Method)
		}
	}

	return nil
}

type Response struct {
	ReqUrl  string
	ReqBody string
	Status  string
	Test    string
	Bm      string
}

var pluginHandlers = make(map[string]PluginInfo)

type pluginHandler func(Data) (Data, error)

type PluginInfo struct {
	DisplayName string
	FieldNames  map[string]string
	Handler     pluginHandler
}

func (this PluginInfo) IsNull() bool {
	return this.DisplayName == "" || this.FieldNames == nil || this.Handler == nil
}

func RegisterPluginHandler(name string, info PluginInfo) error {
	if _, has := pluginHandlers[name]; has {
		return errors.New("plugin has existed, CAN'T register again")
	}
	if info.IsNull() {
		return errors.New("handler CAN'T be NULL")
	}
	pluginHandlers[name] = info
	return nil
}

func HookPlugin(data Data) (Data, error) {
	plugin, has := pluginHandlers[data.Plugin.Key]
	if !has {
		// if not exists, return default handler
		return data, nil
	}
	return plugin.Handler(data)
}

func init() {
	// default plugin: not use!
	RegisterPluginHandler("", PluginInfo{
		DisplayName: "不使用插件",
		FieldNames:  map[string]string{},
		Handler: func(data Data) (Data, error) {
			return data, nil
		},
	})
}
