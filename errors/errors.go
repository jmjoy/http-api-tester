// user errros
package errors

import "github.com/jmjoy/http-api-tester/app"

// ApiStatusError
var (
	ErrBookmarkNameEmpty   = app.NewApiStatusError(1000, "书签名字不能为空")
	ErrBookmarkNameInvalid = app.NewApiStatusError(1001, "书签名字不合格")
	ErrBookmarkNotFound    = app.NewApiStatusError(1002, "书签不存在")
	ErrBookmarkExisted     = app.NewApiStatusError(1003, "书签已存在")

	ErrUrlEmpty           = app.NewApiStatusError(2000, "URL不能为空")
	ErrUrlUnknowScheme    = app.NewApiStatusError(2001, "URL未知协议：%s")
	ErrUrlEmptyHost       = app.NewApiStatusError(2002, "URL的Host不能为空")
	ErrUrlUnknowArgMethod = app.NewApiStatusError(2003, "URL的参数中包含未知请求方式：%s")
)
