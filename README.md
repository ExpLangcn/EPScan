**停止更新（归档处理），请关注后续项目**

**[点击关注 Twitter](https://twitter.com/ExpLang_Cn) 以便快速了解我的动态.**

----

# EPScan
被动收集资产并自动进行SQL注入检测（插件化 自动Bypass）、XSS检测、RCE检测、敏感信息检测

## 注意

~~**由于Golang-Plus的特性，所以Windows不支持Bypass插件，Mac和Linux支持，请下载Plus对应系统的插件否则会报错。**~~

~~**暂时只支持Linux 和 Mac，如果报错 "panic: plugin: not implemented" 就是系统不支持Bypass插件**~~

**已改为通过程序内置解释器加载Bypass插件**

## 程序介绍
- [x] ~~移植SqlMap的检测模式~~
- [x] ~~SQL注入自动Bypass~~
- [x] ~~SQL注入支持Get/Post布尔盲注~~
- [x] ~~SQL注入支持请求头Get/Post报错注入~~
- [x] ~~SQL注入自动加载Bypass插件~~
- [x] ~~插件化Payload改造~~
- [x] ~~插件化Bypass模式~~
- [ ] 自动更新Bypass插件
- [ ] XSS漏洞检测
- [ ] Rce漏洞检测
- [ ] 敏感信息检测
- [ ] 主动爬虫+扫描

## 程序截图

<img src=https://user-images.githubusercontent.com/52586866/202387755-55cb5ad6-c4ea-4cf7-8eb5-be87ab20398a.png width=80% />

<img src=https://user-images.githubusercontent.com/52586866/202387795-09e707fc-4c37-48e5-9584-5595dcc0a27e.png width=80% />

<img src=https://user-images.githubusercontent.com/52586866/202387548-8b1dfe2b-0b15-4ceb-8385-f0e8d0779529.png width=80% />

<img src=https://user-images.githubusercontent.com/52586866/202387599-72faf4e8-d407-4b52-9228-74503ebbad8a.png width=80% />

## 参数教程

**一定要下载crt并安装证书，否则无法扫描https**

```
./EPScan -h

Usage of /EPScan:
  -Wlist string
        只扫描文件内的Host 文件内一行一条 不可带有http协议
	
  -blist string
        文件内的Host不扫描 文件内一行一条 不可带有http协议
	
  -deep 	bool 	// 选择后开启，默认不开启，不开效率快但是误报微高
        深入扫描（降低效率提高准确率）
	
  -pdelay 	string 	// 需带有http协议头，看例子
        流量转发（程序的流量转发到指定端口，例如Burp的8080端口: http://127.0.0.1:8080) 
	
  -pdir 	string	// 加载指定目录下所有的插件，默认不进行Bypass
        批量加载指定目录下的Bypass插件
	
  -plist 	string	// 查看插件信息，默认输出 "plus/" 目录下的插件信息
        查看插件列表和详细信息 (default "plus/")
	
  -port 	string	// 开启监听端口
        监听端口配置 (default "8899")
	
  -ptest 	string	// 测试指定插件效果 会输出原始Payload和处理后的Payload
        插件测试
	
  -run 		bool	// 开始被动式扫描，例如：./EPScan -pdir plus/ -port 8899 -run
        开始被动式扫描
	
  -iflength	bool	// 过滤Length为-1的结果（降低误报）默认不开启
  	过滤Length为-1的结果（降低误报）默认不开启

  -th		int	// 发包线程 默认10
  	发包线程 默认10
```

**开始被动式扫描**

**一定要下载crt并安装证书，否则无法扫描https**

```
./EPScan -port 8899 -run
```

**开始被动式扫描并加载Bypass插件目录**

```
./EPScan -pdir plus/ -port 8899 -run
```

## SQL注入Bypass插件编写

**联系微信：backxyh 贡献大于5个插件即可进入内测群（优先发布测试版和最新版，大量优秀Bypass方案及插件共享）**

### 例：

[**点击查看插件样例**](https://github.com/ExpLangcn/EPScan/blob/main/plus/unionalltounion.go)

```
package main //包名必须为 "main"

import (
	"strings"
)

var ( // 全局变量必须包含以下三个参数
	Plus_Author   = "ExpLang"        // 插件作者
	Plus_Version  = "1.0"            // 插件版本
	Plus_Describe = "Bypass_SafeDog" // 插件描述
)

func Bypass(Value string) string { // 主函数名必须为 "Bypass" 区分大小写，参数为String，返回值为String，思路：传入原始Payload，返回处理后的Payload。

	var Value_Bypass string
	var bypass_SafeDog_str = "/*x^x*/"
	Value_Bypass = strings.Replace(Value, "UNION", bypass_SafeDog_str+"UNION"+bypass_SafeDog_str, -1) // 寻找指定字符串并全部替换
	Value_Bypass = strings.Replace(Value_Bypass, "SELECT", bypass_SafeDog_str+"SELECT"+bypass_SafeDog_str, -1)

	return Value_Bypass // 返回处理后的Payload。
	
}

func INFO() []string {
	return []string{Plus_Author, Plus_Version, Plus_Describe} //输出插件信息
}

```
### 编译为插件：

```
./EPScan -pluslist

+----+---------+---------+-------------------------------+-------------------------------+
| ID | AUTHOR  | VERSION |           DESCRIBE            |             PATH              |
+----+---------+---------+-------------------------------+-------------------------------+
| 1  | ExpLang | 1.0     | Bypass_SafeDog                | plus/bypass_SafeDog.go        |
| 2  | ExpLang | 1.0     | 替换 "'" 单引号为 "%EF%BC%87"   | plus/bypass_apostrophemask.go |
| 3  | ExpLang | 1.0     | 测试/演示   		         | plus/bypass_name.go           |
+----+---------+---------+-------------------------------+-------------------------------+
```

### 测试插件

```
./EPScan -ptest "plus/bypass_name.go"
```

----

#### 😄 I’m ExpLang [**Twitter**](https://twitter.com/ExpLang_Cn) 欢迎关注fo～
