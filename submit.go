package main

//type RequestStruct struct {
//    Method   string
//    URL      *url.URL
//    PostData url.Values
//}

//type SubmitController struct {
//    *Controller
//}

//func NewSubmitController() interface{} {
//    return &SubmitController{
//        &Controller{},
//    }
//}

//func (this *SubmitController) Post() {
//    buf, err := ioutil.ReadAll(this.r.Body)
//    if err != nil {
//        this.RenderJson(400, "读取输入出错[罕见]", nil)
//        return
//    }

//    // 解析输入JSON
//    input := new(Bookmark)
//    err = json.Unmarshal(buf, input)
//    if err != nil {
//        this.RenderJson(40001, "传入参数[JSON]解析出错: "+err.Error(), nil)
//        return
//    }

//    reqS, err := this.getRequestStruct(input)
//    if err != nil {
//        this.RenderJson(40003, err.Error(), nil)
//        return
//    }

//    // plugin
//    reqS, err = HookPlugin(input.Plugin, reqS)
//    if err != nil {
//        this.RenderJson(40010, err.Error(), nil)
//        return
//    }

//    respM := make(map[string]string, 2)
//    // curl
//    err = this.submitTest(reqS, respM)
//    if err != nil {
//        this.RenderJson(40020, err.Error(), nil)
//        return
//    }

//    // benchmark
//    if input.Bm.Switch {
//        err = this.submitBm(reqS, input.Bm.N, input.Bm.C, respM)
//        if err != nil {
//            this.RenderJson(40030, err.Error(), nil)
//            return
//        }
//    }

//    // success
//    this.RenderJson(200, "", respM)
//}

//func (this *SubmitController) getRequestStruct(bookmark *Bookmark) (*RequestStruct, error) {
//    if bookmark.Url == "" {
//        return nil, errors.New("请指定URL")
//    }

//    u, err := url.Parse(bookmark.Url)
//    if err != nil {
//        return nil, err
//    }

//    if u.Scheme != "http" && u.Scheme != "https" {
//        return nil, errors.New("未知协议：" + u.Scheme)
//    }

//    if u.Host == "" {
//        return nil, errors.New("请指定host")
//    }

//    query := u.Query()
//    postData := make(url.Values)

//    for _, v := range bookmark.Args {
//        switch strings.ToUpper(v.Method) {
//        case "GET":
//            query.Add(v.Key, v.Value)

//        case "POST":
//            postData.Add(v.Key, v.Value)

//        default:
//            return nil, errors.New("参数中包含未知请求方式")
//        }
//    }

//    u.RawQuery = query.Encode()

//    return &RequestStruct{
//        Method:   bookmark.Method,
//        URL:      u,
//        PostData: postData,
//    }, nil
//}

//func (this *SubmitController) submitTest(reqS *RequestStruct, respM map[string]string) error {
//    var resp *http.Response
//    var err error

//    switch strings.ToUpper(reqS.Method) {
//    case "GET":
//        resp, err = http.Get(reqS.URL.String())

//    case "POST":
//        resp, err = http.PostForm(reqS.URL.String(), reqS.PostData)

//    default:
//        err = errors.New("未知请求方式")
//    }
//    if err != nil {
//        return err
//    }
//    defer resp.Body.Close()

//    buf, err := ioutil.ReadAll(resp.Body)
//    if err != nil {
//        return err
//    }

//    respM["status"] = resp.Status
//    respM["requrl"] = reqS.URL.String()
//    respM["reqBody"] = reqS.PostData.Encode()
//    respM["test"] = string(buf)

//    return nil
//}

//var gSubmitBmChan = make(chan struct{}, 1)

//func (this *SubmitController) submitBm(reqS *RequestStruct, n, c uint, respM map[string]string) error {

//    select {
//    case gSubmitBmChan <- struct{}{}:
//        defer func() {
//            <-gSubmitBmChan
//        }()

//        // run ab util
//        method := strings.ToUpper(reqS.Method)
//        if n > 2500 {
//            n = 2500
//        }
//        if c > 2500 {
//            c = 2500
//        }

//        cname := "./boom"
//        args := make([]string, 0, 10)
//        args = append(args, "-n")
//        args = append(args, fmt.Sprintf("%d", n))
//        args = append(args, "-c")
//        args = append(args, fmt.Sprintf("%d", c))
//        args = append(args, "-m")
//        args = append(args, method)

//        switch method {
//        case "GET":
//            args = append(args, "-T")
//            args = append(args, "text/plain")

//        case "POST":
//            args = append(args, "-T")
//            args = append(args, "application/x-www-form-urlencoded")
//            if len(reqS.PostData) > 0 {
//                args = append(args, "-d")
//                args = append(args, reqS.PostData.Encode())
//            }

//        default:
//            return errors.New("未知请求方式")
//        }
//        args = append(args, reqS.URL.String())

//        cmd := exec.Command(cname, args...)
//        out := new(bytes.Buffer)
//        cmd.Stdout = out
//        cmd.Env = os.Environ()
//        err := cmd.Run()
//        if err != nil {
//            return err
//        }

//        cAll := cname + " " + strings.Join(args, " ")
//        text := out.String()
//        respM["bmstatus"] = "true"
//        respM["bm"] = cAll + text[strings.Index(text, "\n"):]

//    case <-time.After(30 * time.Second):
//        respM["bmstatus"] = "false"
//        respM["bm"] = "30 second timeout"
//    }

//    return nil
//}
