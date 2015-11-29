package plugin

import (
	"github.com/jmjoy/http-api-tester/bean"
)

func init() {
	bean.RegisterPluginHandler("md5signature", pluginMd5signature)
}

func pluginMd5signature(bean.Data) (bean.Data, error) {
	return bean.Data{}, nil
}
