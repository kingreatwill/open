<!--toc-->
[TOC]
[awesome-go](https://github.com/avelino/awesome-go)
## x
golang.org/x/time/rate
该限流器是基于Token Bucket(令牌桶)实现的。

## goweight 分析模块大小
```
$ cd current-project
$ goweight
```

## godebug:一个跨平台的Go程序调试工具
https://github.com/mailgun/godebug  已过时

NEW
https://github.com/derekparker/delve


## yaegi  go解释器
https://github.com/containous/yaegi
可以提供交互环境

https://github.com/topxeq/gotx

## 热更新
https://github.com/cosmtrek/air


https://github.com/topics/live-reload?l=go

## gops 分析机器上运行了哪些go进程
go get -u github.com/google/gops

```
C:\Users\35084>gops tree
...
├── 16712
│   └── 5988 (gops.exe) {go1.14.1}
├── 4728
│   ├── 16028 (com.docker.backend.exe) {go1.12.16}
│   └── 3708 (com.docker.proxy.exe) {go1.12.16}
└── 5172
    └── 12080 (gopls.exe) {go1.14.1}


C:\Users\35084>gops
18256 16712 gops.exe                go1.14.1  D:\go\bin\gops.exe
16028 4728  com.docker.backend.exe  go1.12.16 C:\Program Files\Docker\Docker\resources\com.docker.backend.exe
12080 5172  gopls.exe               go1.14.1  D:\go\bin\gopls.exe
3708  4728  com.docker.proxy.exe    go1.12.16 C:\Program Files\Docker\Docker\resources\com.docker.proxy.exe

C:\Users\35084>gops 3708
parent PID:     4728
threads:        12
memory usage:   0.058%
cpu usage:      0.001%
username:       DESKTOP-PK520IC\35084
cmd+args:       "com.docker.proxy.exe"  -dockerExe "C:\Program Files\Docker\Docker\resources\bin\docker.exe"  -host-names host.docker.internal,docker.for.win.host.internal,docker.for.win.localhost -gateway-names gateway.docker.internal,docker.for.win.gateway.internal,docker.for.win.http.internal -vm-names vm.docker.internal,docker-for-desktop,docker-desktop,kubernetes.docker.internal -host-ip 192.168.65.2 -gateway-ip 192.168.65.1 -vm-ip 192.168.65.3 -pki "C:\ProgramData\DockerDesktop\pki" -inject-hosts=True
elapsed time:   02:45:26
local/remote:   127.0.0.1:33499 <-> 0.0.0.0:0 (LISTEN)
local/remote:   127.0.0.1:53974 <-> :0 ()
```


## json
github.com/liamylian/json-hashids

序列化自动加密

## orm
facebook开源的新的go语言orm模块，An entity framework for Go
Simple, yet powerful ORM for modeling and querying data.
https://github.com/facebookincubator/ent


## DI


### Google wire 3.3k
https://github.com/google/wire

Wire 可以生成 Go 源码并在编译期完成依赖注入。它不需要反射机制或 [Service Locators](https://en.wikipedia.org/wiki/Service_locator_pattern)

好处：
1. 方便 debug，若有依赖缺失编译时会报错
2. 因为不需要 Service Locators， 所以对命名没有特殊要求
3. 避免依赖膨胀。生成的代码只包含被依赖的代码，而运行时依赖注入则无法作到这一点
4. 依赖关系静态存于源码之中， 便于工具分析与可视化

[Compile-time Dependency Injection With Go Cloud's Wire](https://blog.golang.org/wire)


[一文读懂 Go官方的 Wire](https://mp.weixin.qq.com/s/ZQKi9O7DRJ3qGWhDL9aTVg)

### Uber dig 1.3k
运行时依赖注入
https://github.com/uber-go/dig

### Facebook inject 1.2k(归档了，不更新)
运行时依赖注入
https://github.com/facebookarchive/inject

## spinal-case(脊柱) or snake_case(蛇形) or CamelCase(驼峰式) or KebabCase(短横线) or PascalCase(帕斯卡命名法) or PascalSnakeCase
https://github.com/iancoleman/strcase

## GUI


Cross platform GUI in Go based on Material Design https://fyne.io/
https://github.com/fyne-io/fyne


请注意，默认情况下，Windows应用程序是从命令提示符加载的，这意味着，如果单击图标，则可能会看到命令窗口。 要解决此问题，请在运行或构建命令中添加参数-ldflags -H = windowsgui。



Prerequisites
https://fyne.io/develop/
Windows
1. Download Go from the download page and follow instructions
2. Install one of the available C compilers for windows, the following are tested with Go and Fyne:
    - MSYS2 with MingW-w64 - msys2.org
    - TDM-GCC - tdm-gcc.tdragon.net
    - Cygwin - cygwin.com
3. In Windows your graphics driver will already be installed, but it is recommended to ensure they are up to date.


## eval
github.com/Knetic/govaluate

## 限流熔断
https://github.com/alibaba/sentinel-golang

https://github.com/afex/hystrix-go

golang 提供了拓展库(golang.org/x/time/rate)提供了限流器组件

## 微服务
https://github.com/douyu/jupiter

## book
《Go语法树入门》(开源免费图书/Go语言进阶/掌握抽象语法树/Go语言AST/LLVM/LLIR)
https://github.com/chai2010/go-ast-book


## log
https://github.com/grafana/loki


## 流媒体
monibuca  丰富的内置插件提供了流媒体服务器的常见功能，例如rtmp server、http-flv、视频录制、QoS等
https://github.com/Monibuca 100star


Darwin Streaming Server 没有维护了？
https://github.com/macosforge/dss


高性能开源RTSP流媒体服务器，基于go语言研发，维护和优化：RTSP推模式转发、RTSP拉模式转发、录像、检索、回放、关键帧缓存、秒开画面、RESTful接口、WEB后台管理、分布式负载均衡
https://github.com/EasyDarwin/EasyDarwin   4k 

https://github.com/pion/ion
https://github.com/pion/webrtc

IOT 平台？
https://github.com/nats-io/nats-streaming-server

## Chrome DevTools Protocol 
https://github.com/chromedp/chromedp
https://www.toutiao.com/i6843024797073408515

## 其它
### 一款用 SQL 方式查询 Git 仓库的开源项目进入 GitHub 趋势榜
https://github.com/augmentable-dev/gitqlite


## grpc
### grpcui
他是一个 gRPC 的 Web 页面调试工具，提供交互式的调试界面。
https://github.com/fullstorydev/grpcui

grpcui -plaintext 127.0.0.1:9901
127.0.0.1:9901 是grpc server的地址

## dingtalk
```
package dingtalk
​
import (
    "bytes"
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
    "strconv"
    "time"
)
​
// SendMessage 发送钉钉机器人消息
func SendMessage(url, message string, ats ...string) (respContent string, err error) {
    c := &http.Client{
        Timeout: time.Second * 30,
    }
    data := map[string]interface{}{
        "msgtype": "text",
        "text":    map[string]string{"content": message},
    }
    if len(ats) != 0 {
        isAtAll := false
        atMobiles := []string{}
        for i := range ats {
            if ats[i] == "all" {
                isAtAll = true
            } else {
                atMobiles = append(atMobiles, ats[i])
            }
        }
        data["at"] = map[string]interface{}{
            "isAtAll":   isAtAll,
            "atMobiles": atMobiles,
        }
    }
    b, _ := json.Marshal(data)
    resp, err := c.Post(url, "application/json", bytes.NewReader(b))
    if err != nil {
        return "post请求失败", err
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode == 200 {
        return string(body), nil
    }
    return string(body), errors.New(strconv.Itoa(resp.StatusCode))
}

package main
​
import (
    "fmt"
    "dingtalk"
)
​
func main() {
    result, err := dingtalk.SendMessage("https://oapi.dingtalk.com/robot/send?access_token=XXXXXX", "测试消息通知", "all")
    if err != nil {
        fmt.Println("发送失败", result)
        return
    }
    fmt.Println("发送成功", result)
}
```

## WebAssembly
从Go到JavaScript的编译器，用于在浏览器中运行Go代码
https://github.com/gopherjs/gopherjs


https://github.com/hexops/vecty

## git

### git sql
https://github.com/augmentable-dev/askgit

也可以跑docker
docker run -v F:/github/openjw/openself:/repo:ro augmentable/askgit "SELECT * FROM commits"

## Markdown
将 markdown 中的 go 代码块进行格式化。
https://github.com/po3rin/gofmtmd

https://github.com/JohannesKaufmann/html-to-markdown

## 嵌入文件
https://github.com/golang/proposal/blob/master/design/draft-embed.md#background

- github.com/alecthomas/gobundle
- github.com/GeertJohan/go.rice
- github.com/go-playground/statics
- github.com/gobuffalo/packr
- github.com/knadh/stuffbin
- github.com/mjibson/esc
- github.com/omeid/go-resources
- github.com/phogolabs/parcello
- github.com/pyros2097/go-embed
- github.com/rakyll/statik
- github.com/shurcooL/vfsgen
- github.com/UnnoTed/fileb0x
- github.com/wlbr/templify
- perkeep.org/pkg/fileembed

### go-bindata
https://github.com/go-bindata/go-bindata

## 视频下载

https://github.com/iawia002/annie

抖音	    https://www.douyin.com	✓			
哔哩哔哩	https://www.bilibili.com	✓		✓	✓
半次元	    https://bcy.net		✓		
pixivision	https://www.pixivision.net		✓		
优酷	    https://www.youku.com	✓			✓
YouTube	    https://www.youtube.com	✓		✓	
爱奇艺	    https://www.iqiyi.com	✓			
芒果TV	    https://www.mgtv.com	✓			
糖豆广场舞	http://www.tangdou.com	✓		✓	
Tumblr	    https://www.tumblr.com	✓	✓		
Vimeo	    https://vimeo.com	✓			
Facebook	https://facebook.com	✓			
斗鱼视频	https://v.douyu.com	✓			
秒拍	    https://www.miaopai.com	✓			
微博	    https://weibo.com	✓			
Instagram	https://www.instagram.com	✓	✓		
Twitter	    https://twitter.com	✓			
腾讯视频	https://v.qq.com	✓			
网易云音乐	https://music.163.com	✓			
音悦台	    https://yinyuetai.com	✓			
极客时间	https://time.geekbang.org	✓			
Pornhub	    https://pornhub.com	✓			
XVIDEOS	    https://xvideos.com	✓			
聯合新聞網	https://udn.com	✓			
TikTok	    https://www.tiktok.com	✓			
好看视频    https://haokan.baidu.com	✓


## go包搜索
https://api.godoc.org/search?q=etcd
https://github.com/clearcodecn/gosearch