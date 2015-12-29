package model

import (
	"io/ioutil"
	"net/http"

	"github.com/jmjoy/boom/boomer"
)

var SubmitModel = &submitModel{}

type submitModel struct{}

func (this *submitModel) Submit(data Data) (resp Response, err error) {
	if err = data.Validate(); err != nil {
		return
	}

	data, err = HookPlugin(data)
	if err != nil {
		return
	}

	reqMaker, err := NewRequestMaker(data)
	if err != nil {
		return
	}

	var response Response

	if err = this.submitTest(data, reqMaker, &response); err != nil {
		return
	}

	if err = this.submitBenckmark(data, data.Bm, reqMaker, &response); err != nil {
		return
	}

	return response, nil
}

func (this *submitModel) submitTest(data Data, reqMaker *RequestMaker, response *Response) error {
	req, err := reqMaker.NewRequest()
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response.ReqUrl = req.URL.String()
	response.ReqBody = reqMaker.PostForm.Encode()
	response.Status = resp.Status
	response.Test = string(respBodyBuf)

	return nil
}

func (this *submitModel) submitBenckmark(data Data, bm Bm, reqMaker *RequestMaker, response *Response) error {
	if !bm.Switch {
		return nil
	}

	// limit N, C reduce server pressure.
	var n, c uint
	if bm.N >= 1000 {
		n = 1000
	} else {
		n = bm.N
	}
	if bm.C >= 500 {
		c = 500
	} else {
		n = bm.C
	}

	req, err := reqMaker.NewRequest()
	if err != nil {
		return err
	}

	text := (&boomer.Boomer{
		Request:     req,
		RequestBody: reqMaker.PostForm.Encode(),
		N:           int(n),
		C:           int(c),
		Timeout:     35,
	}).Run()

	response.Bm = text

	return nil
}
