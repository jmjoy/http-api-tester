package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RequestStruct struct {
	Method   string
	URL      *url.URL
	PostData url.Values
}

type SubmitController struct {
	*Controller
}

func NewSubmitController() interface{} {
	return &SubmitController{
		&Controller{},
	}
}

func (this *SubmitController) Post() {
	buf, err := ioutil.ReadAll(this.r.Body)
	if err != nil {
		this.RenderJson(400, "读取输入出错[罕见]", nil)
		return
	}

	// 解析输入JSON
	input := new(Bookmark)
	err = json.Unmarshal(buf, input)
	if err != nil {
		this.RenderJson(40001, "传入参数[JSON]解析出错: "+err.Error(), nil)
		return
	}

	reqS, err := this.getRequestStruct(input)
	if err != nil {
		this.RenderJson(40003, err.Error(), nil)
		return
	}

	reqS, err = HookPlugin(input.Plugin, reqS)
	if err != nil {
		this.RenderJson(40010, err.Error(), nil)
		return
	}

	respM := make(map[string]string, 2)
	err = this.submitTest(reqS, respM)
	if err != nil {
		this.RenderJson(40020, err.Error(), nil)
		return
	}

	this.RenderJson(200, "", respM)
}

func (this *SubmitController) getRequestStruct(bookmark *Bookmark) (*RequestStruct, error) {
	if bookmark.Url == "" {
		return nil, errors.New("请指定URL")
	}

	u, err := url.Parse(bookmark.Url)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, errors.New("未知协议：" + u.Scheme)
	}

	if u.Host == "" {
		return nil, errors.New("请指定host")
	}

	query := u.Query()
	postData := make(url.Values)

	for _, v := range bookmark.Args {
		switch strings.ToUpper(v.Method) {
		case "GET":
			query.Add(v.Key, v.Value)

		case "POST":
			postData.Add(v.Key, v.Value)

		default:
			return nil, errors.New("参数中包含未知请求方式")
		}
	}

	u.RawQuery = query.Encode()

	return &RequestStruct{
		Method:   bookmark.Method,
		URL:      u,
		PostData: postData,
	}, nil
}

func (this *SubmitController) submitTest(reqS *RequestStruct, respM map[string]string) error {
	var resp *http.Response
	var err error

	switch strings.ToUpper(reqS.Method) {
	case "GET":
		resp, err = http.Get(reqS.URL.String())

	case "POST":
		resp, err = http.PostForm(reqS.URL.String(), reqS.PostData)

	default:
		err = errors.New("未知请求方式")
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respM["status"] = resp.Status
	respM["requrl"] = reqS.URL.String()
	respM["reqBody"] = reqS.PostData.Encode()
	respM["test"] = string(buf)

	return nil
}
