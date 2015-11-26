package main

//import (
//    "bytes"
//    "encoding/json"
//    "io/ioutil"
//    "os"
//    "sync"
//)

//type JsonConfig struct {
//    Selected  string              `json:"selected"`
//    Bookmarks map[string]Bookmark `json:"bookmarks"`
//    Plugins   map[string]Plugin   `json:"plugins"`
//}

//type Bookmark struct {
//    Name   string         `json:"name"`
//    Method string         `json:"method"`
//    Url    string         `json:"url"`
//    Args   []Arg          `json:"args"`
//    Bm     Bm             `json:"bm"`
//    Plugin BookmarkPlugin `json:"plugin"`
//}

//type Arg struct {
//    Key    string `json:"key"`
//    Value  string `json:"value"`
//    Method string `json:"method"`
//}

//type Bm struct {
//    Switch bool `json:"switch"`
//    N      uint `json:"n"`
//    C      uint `json:"c"`
//}

//type BookmarkPlugin struct {
//    Key  string            `json:"key"`
//    Data map[string]string `json:"data"`
//}

//type Plugin struct {
//    Name   string            `json:"name"`
//    Fields map[string]string `json:"fields"`
//}

//var gConfigJsonMutex = new(sync.RWMutex)

//func GetConfigJson() *JsonConfig {
//    gConfigJsonMutex.Lock()
//    defer gConfigJsonMutex.Unlock()

//    src, err := ioutil.ReadFile(gConfigPath)
//    if err != nil {
//        panic(err)
//    }

//    jsonConfig := new(JsonConfig)
//    err = json.Unmarshal(src, jsonConfig)
//    if err != nil {
//        panic(err)
//    }

//    return jsonConfig
//}

//func GetConfigJsonString() string {
//    gConfigJsonMutex.Lock()
//    defer gConfigJsonMutex.Unlock()

//    src, err := ioutil.ReadFile(gConfigPath)
//    if err != nil {
//        panic(err)
//    }

//    dst := new(bytes.Buffer)
//    err = json.Compact(dst, src)
//    if err != nil {
//        panic(err)
//    }

//    return dst.String()
//}

//func SaveConfigJson(jsonConfig *JsonConfig) error {
//    buf, err := json.Marshal(jsonConfig)
//    if err != nil {
//        return err
//    }

//    buffer := new(bytes.Buffer)
//    err = json.Indent(buffer, buf, "", "    ")
//    if err != nil {
//        return err
//    }

//    gConfigJsonMutex.Lock()
//    defer gConfigJsonMutex.Unlock()

//    fw, err := os.OpenFile(gConfigPath, os.O_WRONLY|os.O_TRUNC, 0644)
//    if err != nil {
//        return err
//    }
//    defer fw.Close()

//    _, err = buffer.WriteTo(fw)
//    if err != nil {
//        return err
//    }

//    return nil
//}
