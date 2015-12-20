package app

import (
	"encoding/json"
)

type Model struct {
	bucket string
}

func NewModel(bucket string) *Model {
	return &Model{
		bucket: bucket,
	}
}

func (this *Model) Get(key string, data interface{}) (ok bool, err error) {
	buf, err := Db.Get(this.bucket, key)
	if err != nil {
		if err == ErrBucketNotFound {
			err = nil
		}
		return
	}

	if buf == nil {
		return
	}

	if err = json.Unmarshal(buf, data); err != nil {
		return
	}

	return true, nil
}

func (this *Model) Put(key string, data interface{}) (err error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return
	}
	return Db.Put(this.bucket, key, buf)
}

func (this *Model) Delete(key string) (err error) {
	return Db.Delete(this.bucket, key)
}
