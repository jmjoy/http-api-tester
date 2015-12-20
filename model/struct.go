package model

import (
	goerrors "errors"
	"net/url"

	"github.com/jmjoy/http-api-tester/errors"
)

// bookmarks
type BookmarkMap map[string]Data

// named bookmark
type Bookmark struct {
	Name string
	Data Data
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

// Submit Data
type Data struct {
	Method string
	Url    string
	Args   []Arg
	Bm     Bm
	Plugin Plugin
}

func DataDefault() Data {
	return Data{
		Method: "GET",
		Args:   []Arg{},
		Bm: Bm{
			N: 100,
			C: 10,
		},
		Plugin: Plugin{
			Data: map[string]string{},
		},
	}
}

func (this Data) Validate() error {
	if this.Url == "" {
		return errors.ErrUrlEmpty
	}

	u, err := url.Parse(this.Url)
	if err != nil {
		return err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.ErrUrlUnknowScheme.NewMessageSpf(u.Scheme)
	}

	if u.Host == "" {
		return errors.ErrUrlEmptyHost
	}

	for _, arg := range this.Args {
		switch arg.Method {
		case "GET", "POST":
		default:
			return errors.ErrUrlUnknowArgMethod.NewMessageSpf(arg.Method)
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

var pluginPool = make(map[string]PluginInfo)

func PluginPool() map[string]PluginInfo {
	return pluginPool
}

type pluginHandler func(Data) (Data, error)

type PluginInfo struct {
	DisplayName string
	FieldNames  map[string]string
	Handler     pluginHandler `json:"-"`
}

func (this PluginInfo) IsNull() bool {
	return this.DisplayName == "" || this.FieldNames == nil || this.Handler == nil
}

func RegisterPluginHandler(name string, info PluginInfo) error {
	if _, has := pluginPool[name]; has {
		return goerrors.New("plugin has existed, CAN'T register again")
	}
	if info.IsNull() {
		return goerrors.New("handler CAN'T be NULL")
	}
	pluginPool[name] = info
	return nil
}

func HookPlugin(data Data) (Data, error) {
	plugin, has := pluginPool[data.Plugin.Key]
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
