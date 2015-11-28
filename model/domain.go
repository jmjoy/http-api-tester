package model

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

type upsertType string

const (
	UPSERT_ADD    upsertType = "ADD"
	UPSERT_UPDATE upsertType = "UPDATE"
)
