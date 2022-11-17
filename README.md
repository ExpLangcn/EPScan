# EPScan
被动收集资产并自动进行SQL注入检测（插件化 自动Bypass）、XSS检测、RCE检测、敏感信息检测

## 程序介绍
- [x] ~~移植SqlMap的检测模式
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

```
./EPScan -h

Usage of /EPScan:
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
```

**开始被动式扫描**

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
	Value_Bypass = strings.Replace(Value_Bypass, "AND", bypass_SafeDog_str+"AND"+bypass_SafeDog_str, -1)
	Value_Bypass = strings.Replace(Value_Bypass, "=", bypass_SafeDog_str+"="+bypass_SafeDog_str, -1)
	Value_Bypass = strings.Replace(Value_Bypass, " ", bypass_SafeDog_str, -1)
	Value_Bypass = strings.Replace(Value_Bypass, "information_schema.", "%20%20/*!%20%20%20%20INFOrMATION_SCHEMa%20%20%20%20*/%20%20/*^x^^x^*/%20/*!.*/%20/*^x^^x^*/", -1)
	Value_Bypass = strings.Replace(Value_Bypass, "FROM", bypass_SafeDog_str+"FROM"+bypass_SafeDog_str, -1)

	return Value_Bypass // 返回处理后的Payload。
}

```
### 编译为插件：

```
sudo go build --buildmode=plugin -o plus/bypass_name.so bypass/bypass_name/bypass_name.go
```

编译后可以看到plus目录下新增了一个名为："bypass_name"的so文件，运行程序：

```
./EPScan -pluslist

+----+---------+---------+-------------------------------+-------------------------------+
| ID | AUTHOR  | VERSION |           DESCRIBE            |             PATH              |
+----+---------+---------+-------------------------------+-------------------------------+
| 1  | ExpLang | 1.0     | Bypass_SafeDog                | plus/bypass_SafeDog.so        |
| 2  | ExpLang | 1.0     | 替换 "'" 单引号为 "%EF%BC%87"   | plus/bypass_apostrophemask.so |
| 3  | ExpLang | 1.0     | 测试/演示   		         | plus/bypass_name.so           |
+----+---------+---------+-------------------------------+-------------------------------+
```

### 测试插件

```
./EPScan -ptest "plus/bypass/bypass_name/bypass_name.go"
```

----

# Hi there 👋

<!--
**ExpLangcn/ExpLangcn** is a ✨ _special_ ✨ repository because its `README.md` (this file) appears on your GitHub profile.

Here are some ideas to get you started:

- 🔭 I’m currently working on ...
- 🌱 I’m currently learning ...
- 👯 I’m looking to collaborate on ...
- 🤔 I’m looking for help with ...
- 💬 Ask me about ...
- 📫 How to reach me: ...
- 😄 Pronouns: ...
- ⚡ Fun fact: ...
-->

![](https://komarev.com/ghpvc/?username=ExpLangcn)

- 😄 I’m ExpLang
- 我的后续安全项目将会发布在：**[Security-Magic-Weapon 组织](https://github.com/Security-Magic-Weapon) 点击查看！**
- Telegram频道：**[安全技术资讯聚合](https://t.me/sec_info) 实时更新国内外700+个黑客安全技术论坛/社区/博客/公众号的技术文章.**
- Telegram频道：**[安全资源整合互推](https://t.me/secyq) 一个用于收录Telegram上 安全行业 黑客技术 并且还会自动推送最新的CVE漏洞、CNVD漏洞、最新的红队工具等。**

## 我的个人知识星球

本星球会**每日更新**全网优秀安全资源包括但不限于：**安全工具**、**安全脚本**、**安全学习资料**、**安全商业产品的破解版**等，资源方向均与安全各领域相关。

本星球建有运营微信群可及时在群内反馈和互动（索取安全资源），并在运营微信群内拥有Bot机器人用于通知星球最新动态方便各位星友不错过任何优秀动态。

本星球会针对安全资源进行**严格的分类**，方便各位星友可以**快速定位**自己所需的资源，及给各位星友带来更好的**阅读体验**。

本星球会**每月**进行一次**优质资源统计**，会根据本阅读点赞/评论/阅读/下载最多的资源进行排序方便大家更好的浏览。

并且本星球将会**每季度**进行一次**星友知识共享直播**，仅限本星球的星友参与，直播会邀请星球内的部分大佬或在外部邀请在某领域有所建设的大佬进行**免费的知识共享**，并且本星球的续费折扣是平台的最低**5折折扣**！

<img src=https://tva1.sinaimg.cn/large/006y8mN6gy1h6tocodn91j30ku0bggm5.jpg width=40% />



---

<!-- 
[![ExpLangcn's GitHub stats](https://github-readme-stats.vercel.app/api?username=ExpLangcn)](https://twitter.com/ExpLang_Cn)
[![Top Langs](https://github-readme-stats.vercel.app/api/top-langs/?username=ExpLangcn&layout=compact)](https://twitter.com/ExpLang_Cn) -->

<div>
  <a href="https://twitter.com/ExpLang_Cn">
    <img align="left" height="160" src="https://github-readme-stats.vercel.app/api/top-langs/?username=ExpLangcn&layout=compact" />
  </a>
  <a href="https://github.com/ExpLangcn/ExpLangcn/edit/main/README.md">
    <img align="left" height="160" src="https://github-readme-stats.vercel.app/api?username=ExpLangcn&show_icons=true&count_private=true" />
  </a>
</div>
