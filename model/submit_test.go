package model

import (
	"reflect"
	"testing"
)

func TestSubmitTest(t *testing.T) {
	testResp := Response{
		ReqUrl:  testSrv.URL + "?k1=v1&k2=v2",
		Status:  "200 OK",
		Test:    "k1=v1&k2=v2 k3=v3&k4=v4",
		ReqBody: "k3=v3&k4=v4",
	}

	resp, err := SubmitModel.Submit(testData0)
	if err != nil {
		panic(err)
	}

	t.Logf("%#v", resp)
	if !reflect.DeepEqual(testResp, resp) {
		t.Fatal("response not equal!")
	}
}

func TestSubmitBenckmark(t *testing.T) {
	testData1 := testData0
	testData1.Bm = Bm{
		Switch: true,
		N:      10,
		C:      1,
	}

	resp, err := SubmitModel.Submit(testData1)
	if err != nil {
		panic(err)
	}

	t.Log(resp.Bm)
	if resp.Bm == "" {
		t.Fatal("no benchmark data?")
	}
}
