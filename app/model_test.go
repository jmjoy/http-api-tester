package app

import (
	"os"
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestMain(m *testing.M) {
	dbPath := "./model_test.db"
	defer os.Remove(dbPath)

	if err := InitDb(dbPath); err != nil {
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

func TestKeys(t *testing.T) {
	bucket := "test0"
	model := NewModel(bucket)

	keys, err := model.Keys()
	if err != nil {
		panic(err)
	}
	if len(keys) != 0 {
		t.Fatal("not empty?")
	}

	testKeys := []string{
		"1", "2", "3",
	}
	for _, k := range testKeys {
		model.Put(k, "data")
	}

	keys, err = model.Keys()
	if err != nil {
		panic(err)
	}
	t.Log(keys)
	if !reflect.DeepEqual(testKeys, keys) {
		t.Fatal("not equal?")
	}
}
