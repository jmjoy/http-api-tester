package model

import (
	"os"
	"testing"

	"github.com/jmjoy/http-api-tester/app"
)

func TestMain(m *testing.M) {
	dbPath := "./model_test.db"
	defer os.Remove(dbPath)

	if err := app.InitDb(dbPath); err != nil {
		panic(err)
	}
	m.Run()
}
