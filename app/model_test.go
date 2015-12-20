package app

import (
	"os"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestMain(m *testing.M) {
	dbPath := "./model_test.db"
	defer os.Remove(dbPath)

	if err := initDb(dbPath); err != nil {
		panic(err)
	}
	m.Run()
}

func TestCRUD(t *testing.T) {
	bucket := "test"
	testkey := "testkey"

	model := NewModel(bucket)

	var data map[string]string
	has, err := model.Get(testkey, &data)
	if err != nil {
		panic(err)
	}
	if has {
		t.Fatal("Has data?")
	}

	person := Person{
		Name: "jmjoy",
		Age:  23,
	}
	err = model.Put(testkey, person)
	if err != nil {
		panic(err)
	}

	var person0 Person
	has, err = model.Get(testkey, &person0)
	if err != nil {
		panic(err)
	}
	if !has {
		t.Fatal("No data?")
	}

	if person != person0 {
		t.Fatal("Not euqal!")
	}

	err = model.Delete(testkey)
	if err != nil {
		panic(err)
	}

	has, err = model.Get(testkey, &person0)
	if err != nil {
		panic(err)
	}
	if has {
		t.Fatal("Has data?")
	}
}
