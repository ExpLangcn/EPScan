package main

import (
	"strings"
)

var (
	Plus_Author   = "ExpLang"                                      // 插件作者
	Plus_Version  = "1.0"                                          // 插件版本
	Plus_Describe = "替换 \"'\" 单引号为 \"\\'\" 和 \"\"\" 双引号为 \"\\\"\"" // 插件描述
)

func Bypass(Value string) string {
	var Value_Bypass string
	Value_Bypass = strings.Replace(Value, "'", "\\'", -1)
	Value_Bypass = strings.Replace(Value_Bypass, "\"", "\\\"", -1)
	return Value_Bypass
}

func INFO() []string {
	return []string{Plus_Author, Plus_Version, Plus_Describe}
}
