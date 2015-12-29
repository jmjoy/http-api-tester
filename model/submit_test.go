package model

import (
	"reflect"
	"testing"
)

func TestSubmitTest(t *testing.T) {
	testData := Data{
		Method: "GET",
		Url:    testSrv.URL,
		Args: []Arg{
			Arg{"k1", "v1", "GET"},
			Arg{"k2", "v2", "GET"},
			Arg{"k3", "v3", "POST"},
			Arg{"k4", "v4", "POST"},
		},
	}

	testResp := Response{
		ReqUrl:  testSrv.URL + "?k1=v1&k2=v2",
		Status:  "200 OK",
		Test:    "k1=v1&k2=v2 k3=v3&k4=v4",
		ReqBody: "k3=v3&k4=v4",
	}

	resp, err := SubmitModel.Submit(testData)
	if err != nil {
		panic(err)
	}

	t.Logf("%#v", resp)
	if !reflect.DeepEqual(testResp, resp) {
		t.Fatal("response not equal!")
	}
}
