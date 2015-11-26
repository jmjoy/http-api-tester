package base

import (
	"github.com/boltdb/bolt"
)

var globalDbPath string

type dbHelper func(string) error

var Db dbHelper = func(dbPath string) error {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}
	if err = db.Close(); err != nil {
		return err
	}
	globalDbPath = dbPath
	return nil
}

func (this dbHelper) Query() {

}
