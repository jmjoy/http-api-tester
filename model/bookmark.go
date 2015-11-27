package model

import (
	"encoding/json"
	"fmt"

	"github.com/jmjoy/http-api-tester/base"
)

type BookmarkModel struct {
}

func NewBookmarkModel() *BookmarkModel {
	return new(BookmarkModel)
}

func (this *BookmarkModel) Get(name string) (Data, error) {
	fmt.Println("bookmark get:", name)

	bookmark, err := base.Db.Get("bookmarks", name)

	fmt.Println(string(bookmark), err)

	if err != nil {
		return this.handleGetError(err)
	}

	fmt.Println("bookmark get no error")

	if bookmark == nil {
		return this.DefaultData(), nil
	}

	fmt.Println(string(bookmark))

	var data Data
	err = json.Unmarshal(bookmark, &data)
	return data, err
}

func (this *BookmarkModel) GetCurrent() (Data, error) {
	name, err := base.Db.Get("bookmark", "selected")
	if err != nil {
		return this.handleGetError(err)
	}

	return this.Get(string(name))
}

func (this *BookmarkModel) Add(bookmark Bookmark) error {
	if err := this.validateBookmarkName(bookmark.Name); err != nil {
		return err
	}

	// check is exists?
	buf, _ := base.Db.Get("bookmarks", bookmark.Name)
	if buf != nil {
		return fmt.Errorf("该书签名字已经存在了")
	}

	buf, err := json.Marshal(bookmark.Data)
	if err != nil {
		return err
	}

	if err = base.Db.Put("bookmarks", bookmark.Name, buf); err != nil {
		return err
	}

	return nil
}

func (this *BookmarkModel) DefaultData() Data {
	return Data{
		Method: "GET",
		Args:   []Arg{},
		Plugin: Plugin{
			Data: map[string]string{},
		},
	}
}

func (this *BookmarkModel) handleGetError(err error) (Data, error) {
	if err == base.ErrorBucketNotFound {
		return this.DefaultData(), nil
	}
	return Data{}, err
}

func (this *BookmarkModel) validateBookmarkName(name string) error {
	// TODO 暂时允许所有名字
	return nil
}
