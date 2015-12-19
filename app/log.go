package app

import (
	"fmt"
	"net/http"

	"github.com/ttacon/chalk"
)

type LogLevel int

const (
	LOG_LV_SUCC = iota
	LOG_LV_INFO
	LOG_LV_FAIL
)

func Log(lv LogLevel, e interface{}) {
	var color chalk.Color
	var tip string

	switch lv {
	case LOG_LV_SUCC:
		color = chalk.Green
		tip = "SUCC"

	case LOG_LV_INFO:
		color = chalk.Cyan
		tip = "INFO"

	case LOG_LV_FAIL:
		color = chalk.Red
		tip = "FAIL"

	default:
		panic("No this log level")
	}

	message := color.Color(fmt.Sprintf("[%s] %s", tip, e))
	fmt.Println(message)
}

func LogStatusError(r *http.Request, err *statusError) {
	message := fmt.Sprintf("<%d> %s (%s)", err.status, err.message, r.URL)
	Log(LOG_LV_FAIL, message)
}
