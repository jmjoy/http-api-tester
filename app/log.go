package app

import (
	"github.com/fatih/color"
)

type LogLevel int

const (
	LOG_LV_SUCC LogLevel = iota
	LOG_LV_INFO
	LOG_LV_FAIL
)

func Log(lv LogLevel, e interface{}) {
	var attr color.Attribute
	var tip string

	switch lv {
	case LOG_LV_SUCC:
		attr = color.FgGreen
		tip = "SUCC"

	case LOG_LV_INFO:
		attr = color.FgCyan
		tip = "INFO"

	case LOG_LV_FAIL:
		attr = color.FgRed
		tip = "FAIL"

	default:
		panic("No this log level")
	}

	color.New(attr).Printf("[%s] %s\n", tip, e)
}
