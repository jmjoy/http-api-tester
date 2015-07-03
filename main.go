package main

import (
	"strconv"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

func init() {
	flag.IntVar(&port, "p", 8080, "服务器运行端口")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("index.html")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		io.Copy(w, f)
	})

	http.HandleFunc("/submit", submitHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	respData := make(map[string]interface{})

	if err := submitTest(w, r, respData); err != nil {
		handleErr(err, w)
		return
	}

	respJson := Resp{
		Status: 200,
		Data:   respData,
	}

	buf, err := json.Marshal(respJson)
	if err != nil {
		handleErr(err, w)
		return
	}

	//fmt.Println(string(buf))

	w.Write(buf)
}

func submitTest(w http.ResponseWriter, r *http.Request, respData map[string]interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	//fmt.Println(string(buf))

	var reqJson Req
	err = json.Unmarshal(buf, &reqJson)
	if err != nil {
		return err
	}

	//fmt.Println(reqJson)
	//fmt.Println()

	u, err := url.Parse(reqJson.Url)
	if err != nil {
		return err
	}

	query := u.Query()
	postData := make(url.Values)

	for _, v := range reqJson.Args {
		switch strings.ToUpper(v.Method) {
		case "GET":
			query.Add(v.Key, v.Value)

		case "POST":
			postData.Add(v.Key, v.Value)

		default:
			return errors.New("参数中包含未知请求方式")
		}
	}

	u.RawQuery = query.Encode()

	//fmt.Println(u)
	//fmt.Println(postData)
	//fmt.Println(strings.ToUpper(reqJson.Method))
	//fmt.Println()

	aTime := time.Now()

	var resp *http.Response
	switch strings.ToUpper(reqJson.Method) {
	case "GET":
		resp, err = http.Get(u.String())

	case "POST":
		resp, err = http.PostForm(u.String(), postData)

	default:
		return errors.New("未知请求方式")
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	useTime := time.Now().Sub(aTime).Seconds()

	respData["Status"] = resp.StatusCode
	respData["Data"] = string(buf)
	respData["Times"] = useTime

	if err := submitBoom(u.String(), &reqJson, respData, postData); err != nil {
		return err
	}

	return nil
}

func submitBoom(urlStr string, reqJson *Req, respData map[string]interface{}, postData url.Values) error {
	command := "./boom"
	args := []string {
		"-n", strconv.Itoa(reqJson.N),
		"-c", strconv.Itoa(reqJson.C),
		"-m", strings.ToUpper(reqJson.Method),
		"-d", postData.Encode(),
		urlStr,
	}
	
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}

	for i := range args {
		command += " " + args[i]
	}

	respData["Boom"] = command + "\n" + out.String()
	
	return nil
}

func handleErr(err error, w http.ResponseWriter) {
	if err != nil {
		resp := Resp{
			Status: 400,
			Msg:    err.Error(),
		}
		buf, _ := json.Marshal(resp)
		//log.Println(string(buf))
		w.Write(buf)
	}
}

type Req struct {
	Method string
	Url    string
	Args   []Arg
	C      int
	N      int
}

type Arg struct {
	Key    string
	Value  string
	Method string
}

type Resp struct {
	Status int
	Msg    string
	Data   map[string]interface{}
}

var port int
