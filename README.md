# http-api-tester

HTTP接口测试工具

## Shortcut

![shortcut](https://raw.githubusercontent.com/jmjoy/http-api-tester/master/img/shortcut0.6.bmp)

## Download

Just one file ~! (Since v0.5)

[releases](https://github.com/jmjoy/http-api-tester/releases)

## Compile

### Requirements

    go get "github.com/boltdb/bolt" \
           "github.com/fatih/color" \
           "github.com/jmjoy/boomer" \
           "github.com/jmjoy/file2string"

### Install

    bower install
    sh create-view.sh
    go build
    
## Usage

    → ./http-api-tester --help
    Usage of ./http-api-tester:
      -db string
          数据库路径 (default "http-api-tester.db")
      -p string
          服务器运行端口 (default "8080")
    
## TODO

- [x] Support all http request method
- [x] Support more request body format
- [x] Adjust config file
- [x] Support history list
- [x] Support move header

## License

[MIT](https://github.com/jmjoy/http-interface-tester/blob/master/LICENSE)
