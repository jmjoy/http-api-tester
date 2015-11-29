package plugin

import (
	"crypto/md5"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/jmjoy/http-api-tester/bean"
)

func init() {
	handler := func(data bean.Data) (bean.Data, error) {
		keyName, has := data.Plugin.Data["keyName"]
		if !has {
			return bean.Data{}, errors.New("md5 signature key name DOESN't exist!")
		}

		password, has := data.Plugin.Data["password"]
		if !has {
			return bean.Data{}, errors.New("md5 signature password DOESN'T exist!")
		}

		argMap := make(map[string]string, len(data.Args))
		argKeys := make([]string, 0, len(data.Args))
		for _, arg := range data.Args {
			// if method == "GET", only when args method is "GET", arg will be signature
			if data.Method != "GET" || arg.Method == "GET" {
				argKeys = append(argKeys, arg.Key)
			}

			argMap[arg.Key] = arg.Value
		}
		sort.Strings(argKeys)

		values := make([]string, 0, len(argKeys))
		for _, argKey := range argKeys {
			values = append(values, argMap[argKey])
		}

		values = append(values, password)
		text := strings.Join(values, "")
		md5Text := fmt.Sprintf("%x", md5.Sum([]byte(text)))
		data.Args = append(data.Args, bean.Arg{
			Key:    keyName,
			Value:  md5Text,
			Method: "GET",
		})

		return data, nil
	}

	bean.RegisterPluginHandler("md5signature", bean.PluginInfo{
		DisplayName: "MD5签名认证",
		FieldNames: map[string]string{
			"keyName":  "密钥名称",
			"password": "密码",
		},
		Handler: handler,
	})
}
