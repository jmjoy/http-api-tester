package main

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

var gJsonConfig = new(JsonConfig)
