package model

import (
	goerrors "errors"
	"net/http"
	"net/url"

	"strings"

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

// Header
type Header struct {
	Key   string
	Value string
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
	Method       string
	Url          string
	Args         []Arg
	Headers      []Header
	Bm           Bm
	Plugin       Plugin
	Enctype      string
	JsonContent  string
	PlainContent string
}

func DataDefault() Data {
	return Data{
		Method:  "GET",
		Args:    []Arg{},
		Headers: []Header{},
		Bm: Bm{
			N: 100,
			C: 10,
		},
		Plugin: Plugin{
			Key:  PLUGIN_DEFAULT_NAME,
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

type RequestMaker struct {
	Method      string
	Url         *url.URL
	ContentType string
	Body        string
	Headers     []Header
}

func NewRequestMaker(data Data) (reqMaker *RequestMaker, err error) {
	u, err := url.Parse(data.Url)
	if err != nil {
		return
	}

	var contentType string
	var body string

	switch data.Enctype {
	case "x_www":
		contentType = "application/x-www-form-urlencoded"

		querys := u.Query()
		forms := make(url.Values)

		for _, arg := range data.Args {
			switch arg.Method {
			case "GET":
				querys.Add(arg.Key, arg.Value)

			case "POST":
				forms.Add(arg.Key, arg.Value)
			}
		}

		u.RawQuery = querys.Encode()

		body = forms.Encode()

	case "json":
		contentType = "text/json"
		body = data.JsonContent

	case "plain":
		contentType = "text/plain"
		body = data.PlainContent
	}

	reqMaker = &RequestMaker{
		Method:      data.Method,
		Url:         u,
		ContentType: contentType,
		Body:        body,
		Headers:     data.Headers,
	}
	return
}

func (this *RequestMaker) NewRequest() (request *http.Request, err error) {
	request, err = http.NewRequest(
		this.Method,
		this.Url.String(),
		strings.NewReader(this.Body),
	)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", this.ContentType)

	for _, header := range this.Headers {
		request.Header.Set(http.CanonicalHeaderKey(header.Key), header.Value)
	}

	return
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
	RegisterPluginHandler(PLUGIN_DEFAULT_NAME, PluginInfo{
		DisplayName: "不使用插件",
		FieldNames:  map[string]string{},
		Handler: func(data Data) (Data, error) {
			return data, nil
		},
	})
}
