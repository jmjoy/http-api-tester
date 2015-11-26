package base

import (
	"github.com/boltdb/bolt"
)

var globalDbPath string

type dbHelper func(string) error

// Db is a single instance of dbHelper
var Db dbHelper = func(dbPath string) error {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	globalDbPath = dbPath
	return nil
}

func (this dbHelper) Get(bucket string, key string) ([]byte, error) {
	db, err := bolt.Open(globalDbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var value []byte
	err = db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(bucket))
		if bk != nil {
			return ErrorBucketNotFound
		}
		value = bk.Get([]byte(key))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this dbHelper) Put(bucket string, key string, value []byte) error {
	db, err := bolt.Open(globalDbPath, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = bk.Put([]byte(key), value)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (this dbHelper) Delete(bucket string, key string) error {
	db, err := bolt.Open(globalDbPath, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(bucket))
		if bk != nil {
			return ErrorBucketNotFound
		}
		err = bk.Delete([]byte(key))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
