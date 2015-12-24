package model

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

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

	var response Response

	if err = this.submitTest(data, &response); err != nil {
		return
	}

	if err = this.submitBenckmark(data, data.Bm, &response); err != nil {
		return
	}

	return response, nil
}

func (this *submitModel) makeRequest(data Data) (*http.Request, error) {
	u, err := url.Parse(data.Url)
	if err != nil {
		return nil, err
	}

	querys := u.Query()
	forms := make(url.Values)

	for _, arg := range data.Args {
		switch arg.Method {
		case "GET":
			querys.Add(arg.Key, arg.Value)

		case "POST":
			forms.Add(arg.Key, arg.Value)
		}
	}

	u.RawQuery = querys.Encode()
	body := strings.NewReader(forms.Encode())

	return http.NewRequest(data.Method, u.String(), body)
}

func (this *submitModel) submitTest(data Data, response *Response) error {
	req, err := this.makeRequest(data)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	req, err = this.makeRequest(data)
	if err != nil {
		return err
	}

	reqBodyBuf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	respBodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response.ReqUrl = req.URL.String()
	response.ReqBody = string(reqBodyBuf)
	response.Status = resp.Status
	response.Test = string(respBodyBuf)

	return nil
}

func (this *submitModel) submitBenckmark(data Data, bm Bm, response *Response) error {
	if !bm.Switch {
		return nil
	}

	req, err := this.makeRequest(data)
	if err != nil {
		return err
	}

	bodyBuf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
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

	req, err = this.makeRequest(data)
	if err != nil {
		return err
	}

	text := (&boomer.Boomer{
		Request:     req,
		RequestBody: string(bodyBuf),
		N:           int(n),
		C:           int(c),
		Timeout:     35,
	}).Run()

	response.Bm = text

	return nil
}
