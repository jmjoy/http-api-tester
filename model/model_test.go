package model

import (
	"os"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/jmjoy/http-api-tester/app"
)

var testSrv *httptest.Server

func TestMain(m *testing.M) {
	dbPath := "./model_test.db"
	defer os.Remove(dbPath)

	// init db
	if err := app.InitDb(dbPath); err != nil {
		panic(err)
	}

	// test server
	testSrv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			panic(err)
		}
		content := r.URL.Query().Encode() + " " + r.PostForm.Encode()
		w.Write([]byte(content))
	}))
	defer testSrv.Close()

	m.Run()
}
