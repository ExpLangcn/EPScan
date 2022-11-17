package main //包名必须为 "main"

import (
	"strings"
)

var ( // 全局变量必须包含以下三个参数
	Plus_Author   = "ExpLang" // 插件作者
	Plus_Version  = "1.0"     // 插件版本
	Plus_Describe = "插件样例"    // 插件描述
)

func Bypass(Value string) string { // 主函数名必须为 "Bypass" 区分大小写，参数为String，返回值为String，思路：传入原始Payload，返回处理后的Payload。
	var Value_Bypass string
	var Bypass_Payload = "处理一下"
	Value_Bypass = strings.Replace(Value, "UNION", Bypass_Payload+"UNION"+Bypass_Payload, -1) // 寻找指定字符串并全部替换
	return Value_Bypass                                                                       // 返回处理后的Payload。
}
