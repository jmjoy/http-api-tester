package model

import (
	"os"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/jmjoy/http-api-tester/app"
)

var testSrv *httptest.Server
var testData0 Data

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

	initData()

	m.Run()
}

func initData() {
	testData0 = Data{
		Method: "GET",
		Url:    testSrv.URL,
		Args: []Arg{
			Arg{"k1", "v1", "GET"},
			Arg{"k2", "v2", "GET"},
			Arg{"k3", "v3", "POST"},
			Arg{"k4", "v4", "POST"},
		},
	}

}
