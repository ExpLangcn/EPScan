# EPScan
被动收集资产并自动进行SQL注入检测（插件化 自动Bypass）、XSS检测、RCE检测、敏感信息检测

## 程序介绍
- [x] ~~自动更新Bypass插件~~
- [x] ~~SQL注入自动Bypass~~
- [x] ~~SQL注入支持Get/Post布尔盲注~~
- [x] ~~SQL注入支持请求头Get/Post报错注入~~
- [x] ~~SQL注入自动加载Bypass插件~~

* SQL注入漏洞检测

  * 移植SqlMap的检测模式

  * 插件化Payload改造

  * 插件化Bypass模式

  * 支持Get/Post报错注入

  * 支持Get/Post布尔盲注

  * 支持请求头Get/Post报错注入

  * 自动加载Bypass插件


- [ ] XSS漏洞检测

- [ ] Rce漏洞检测

- [ ] 敏感信息检测

## 程序截图

## SQL注入Bypass插件编写

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

----

### Hi there 👋

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

### 我的个人知识星球

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
