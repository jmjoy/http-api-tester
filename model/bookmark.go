package model

import "github.com/jmjoy/http-api-tester/base"

type BookmarkModel struct {
	*base.Model
}

func NewBookmarkModel() *BookmarkModel {
	return &BookmarkModel{
		Model: base.NewModel("bookmark"),
	}
}
