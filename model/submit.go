package model

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/jmjoy/http-api-tester/bean"
)

type SubmitModel struct {
}

func NewSubmitModel() *SubmitModel {
	return new(SubmitModel)
}

func (this *SubmitModel) Submit(data bean.Data) (bean.Response, error) {
	if err := data.Validate(); err != nil {
		return bean.Response{}, err
	}

	data, err := bean.HookPlugin(data)
	if err != nil {
		return bean.Response{}, err
	}

	req, err := this.makeRequest(data)
	if err != nil {
		return bean.Response{}, err
	}

	var response bean.Response

	if err = this.submitTest(req, &response); err != nil {
		return bean.Response{}, err
	}

	if err = this.submitBenckmark(req, &response); err != nil {
		return bean.Response{}, err
	}

	return response, nil
}

func (this *SubmitModel) makeRequest(data bean.Data) (*http.Request, error) {
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

func (this *SubmitModel) submitTest(req *http.Request, response *bean.Response) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

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

func (this *SubmitModel) submitBenckmark(req *http.Request, response *bean.Response) error {
	return nil
}
