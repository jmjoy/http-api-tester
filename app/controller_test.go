package app

import (
	"net/http"
	"reflect"
	"testing"
)

type MyController struct {
	*Controller
}

func TestController(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	c0 := MyController{
		Controller: &Controller{R: req},
	}
	t.Logf("%#v", c0)

	c := new(MyController)
	reflect.ValueOf(c).Elem().FieldByName("Controller").Set(reflect.ValueOf(&Controller{R: req}))
	// c.Controller.Reset(nil, req)
	t.Logf("%#v", c)

	if !reflect.DeepEqual(c0, c) {
		// t.Fatal("not equal!")
	}
}
