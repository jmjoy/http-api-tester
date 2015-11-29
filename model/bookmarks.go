package model

import (
	"encoding/json"
	"fmt"

	"github.com/jmjoy/http-api-tester/base"
)

type BookmarksModel struct {
}

func NewBookmarksModel() *BookmarksModel {
	return new(BookmarksModel)
}

func (this *BookmarksModel) Get(name string) (Data, error) {
	if err := this.validateBookmarkName(name); err != nil {
		return Data{}, err
	}

	bookmark, err := base.Db.Get("bookmarks", name)

	fmt.Println(string(bookmark), err)

	if err != nil {
		return Data{}, err
	}

	fmt.Println("bookmark get no error")

	fmt.Println("bookmark", string(bookmark))

	if bookmark == nil {
		return Data{}, base.ErrorBookmarkNotFound
	}

	var data Data
	err = json.Unmarshal(bookmark, &data)
	return data, err
}

func (this *BookmarksModel) Upsert(bookmark Bookmark, typ upsertType) error {
	if err := this.validateBookmarkName(bookmark.Name); err != nil {
		return err
	}

	// check is exists? when add
	if typ == UPSERT_ADD {
		buf, _ := base.Db.Get("bookmarks", bookmark.Name)
		if buf != nil {
			return fmt.Errorf("该书签名字已经存在了")
		}
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

func (this *BookmarksModel) Delete(name string) error {
	if err := this.validateBookmarkName(name); err != nil {
		return err
	}

	if err := base.Db.Delete("bookmarks", name); err != nil {
		return err
	}

	return nil
}

func (this *BookmarksModel) validateBookmarkName(name string) error {
	if name == "" {
		return fmt.Errorf("书签名字不能为空")
	}
	// TODO 暂时允许所有名字
	return nil
}
