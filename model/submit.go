package model

import (
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

func (this *SubmitModel) Submit(data bean.Data) error {
	if err := data.Validate(); err != nil {
		return err
	}

	data, err := bean.HookPlugin(data)
	if err != nil {
		return err
	}

	req, err := this.makeRequest(data)
	if err != nil {
		return err
	}

	_ = req

	return nil
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
