package main

import "strings"

var (
	Plus_Author   = "ExpLang"                                // 插件作者
	Plus_Version  = "1.0"                                    // 插件版本
	Plus_Describe = "替换 \"UNION ALL SELECT\" 为 UNION SELECT" // 插件描述
)

func Bypass(Value string) string {
	var Value_Bypass string
	Value_Bypass = strings.Replace(Value, "UNION ALL SELECT", "UNION SELECT", -1)
	return Value_Bypass
}

func INFO() []string {
	return []string{Plus_Author, Plus_Version, Plus_Describe}
}
