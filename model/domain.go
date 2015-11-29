package model

import "errors"

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

type Response struct {
	Status int
	Url    string
	Body   string
	Test   string
	Bm     string
}

type upsertType string

const (
	UPSERT_ADD    upsertType = "ADD"
	UPSERT_UPDATE upsertType = "UPDATE"
)

type pluginHandler func(Data) (Data, error)

var pluginHandlers map[string]pluginHandler

func RegisterPluginHandler(name string, handler pluginHandler) error {
	if _, has := pluginHandlers[name]; has {
		return errors.New("plugin has existed, CAN'T register again")
	}
	if handler == nil {
		return errors.New("handler CAN'T be nil")
	}
	pluginHandlers[name] = handler
	return nil
}
