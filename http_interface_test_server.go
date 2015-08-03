package main

import (
	"bytes"
	"crypto/md5"
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
	"sort"
	"strconv"
	"strings"
	"time"
)

func init() {
	flag.IntVar(&port, "p", 8080, "服务器运行端口")
	flag.Parse()
}

func main() {
	route()

	log.Printf("测试接口服务器在跑了，请访问 http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func route() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.HandleFunc("/submit", submitHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	HandleRestful("/bookmark", NewBookmarkController())
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("favicon.ico")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(w, f)
}

type Restful interface {
	SetWR(http.ResponseWriter, *http.Request)
	Get()
	Post()
	Put()
	Delete()
}

func HandleRestful(pattern string, rf Restful) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		rf.SetWR(w, r)

		switch r.Method {
		case "GET":
			rf.Get()

		case "POST":
			rf.Post()

		case "PUT":
			rf.Put()

		case "DELETE":
			rf.Delete()

		default:
			w.Write([]byte("I don't know this method, get out please!"))
		}
	})
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

	// 生成密钥，并添加到请求参数中
	query.Add(reqJson.Key, generateSecretKey(reqJson.Secret, reqJson.Args))

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

	if reqJson.N > 1000 {
		reqJson.N = 1000
	}
	if reqJson.C > 100 {
		reqJson.C = 100
	}

	command := "./boom"
	args := []string{
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

func generateSecretKey(secret string, args []Arg) string {
	argsMap := make(map[string]string)
	for i := range args {
		argsMap[args[i].Key] = args[i].Value
	}

	argsKeys := make([]string, 0, 4)
	for k := range argsMap {
		argsKeys = append(argsKeys, k)
	}
	sort.Strings(argsKeys)

	values := make([]string, 0, 4)
	for _, key := range argsKeys {
		values = append(values, argsMap[key])
	}
	values = append(values, secret)

	text := strings.Join(values, "")

	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
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
	Key    string
	Secret string
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

const VERSION = "0.3"
