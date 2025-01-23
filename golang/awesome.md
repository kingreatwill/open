<!--toc-->
[TOC]
[awesome-go](https://github.com/avelino/awesome-go)

## 限流熔断降级
[【GO】golang 降级|熔断|限流实战](https://www.jianshu.com/p/be5c139c11e3)
[Go实现各类限流](https://juejin.cn/post/6959436201443426311) 推荐看看
### sentinel-golang 限流熔断
https://github.com/alibaba/sentinel-golang 2.5k

### ratelimit 限流
https://github.com/uber-go/ratelimit 3.7k

> 把请求数平均到耗时, 比如QPS=10/s 那么两个请求间隔100ms; 缺点是只能阻塞获取

### ratelimit 限流
https://github.com/juju/ratelimit 2.8k
功能更丰富


### x/time/rate 限流
golang.org/x/time/rate
该限流器是基于Token Bucket(令牌桶)实现的。
golang 提供了拓展库(golang.org/x/time/rate)提供了限流器组件
https://pkg.go.dev/golang.org/x/time/rate

功能更丰富

### gobreaker 熔断降级
https://github.com/sony/gobreaker 2.5k

### hystrix-go 熔断降级
https://github.com/afex/hystrix-go 4k


### bytedance Circuitbreaker/熔断
https://github.com/bytedance/gopkg/blob/develop/cloud/circuitbreaker/README_ZH.MD

### tollbooth
Simple middleware to rate-limit HTTP requests.
https://github.com/didip/tollbooth 2.4k


## 缓存
### singleflight
缓存击穿
将一组相同的请求合并成一个请求，实际上只会去请求一次，然后对所有的请求返回相同的结果。
"golang.org/x/sync/singleflight"
"github.com/zeromicro/go-zero/core/syncx"的 syncx.SingleFlight

> 如果第一个请求在执行了, 后面的请求阻塞等待获取结果(会指定key, 相同的key才会)

### KV DB
https://github.com/buraksezer/olric

https://github.com/dgraph-io/badger

go.etcd.io/bbolt
github.com/syndtr/goleveldb

## Goroutine
### Goroutine 泄漏防治神器 goleak
goroutine 作为 golang 并发实现的核心组成部分，非常容易上手使用，但却很难驾驭得好。我们经常会遭遇各种形式的 goroutine 泄漏，这些泄漏的 goroutine 会一直存活直到进程终结。它们的占用的栈内存一直无法释放、关联的堆内存也不能被 GC 清理，系统的可用内存会随泄漏 goroutine 的增多越来越少，直至崩溃！
https://github.com/uber-go/goleak


## io
### io
io.Reader and io.Writer
### bufio
bufio.Reader或bufio.Writer
### Metered

io.Reader and io.Writer的替代品, 支持统计字节数
https://github.com/samber/go-metered-io
- 性能监控: 实时监控数据传输量,及时发现性能问题
- 流量监控: 可实现精细化流量控制


## 分析/debug
### pprof
[Profiling](./pprof.md)
### visualgdb 工具
https://visualgdb.com/gdbreference/commands/x

### goweight 分析模块大小
https://github.com/jondot/goweight
```
$ go get github.com/jondot/goweight
$ cd current-project
$ goweight
```
### go-binsize-treemap 二进制大小分析
只支持linux的二进制, 可以看到每个包的大小(与goweight类似, 但是以图的形式展示)
```
$ go install github.com/nikolaydubina/go-binsize-treemap@latest
$ go tool nm -size <binary finename> | go-binsize-treemap > binsize.svg
```

### godebug:一个跨平台的Go程序调试工具
https://github.com/mailgun/godebug  已过时

### godebug:delve
https://github.com/go-delve/delve

### yaegi  go解释器
https://github.com/containous/yaegi
可以提供交互环境

https://github.com/topxeq/gotx

### 热更新
https://github.com/cosmtrek/air
https://github.com/air-verse/air

https://github.com/topics/live-reload?l=go

### 监控分析
https://github.com/nikolaydubina/go-recipes#monitoring

### gops 分析机器上运行了哪些go进程(类似jps)
go get -u github.com/google/gops
go install github.com/google/gops@latest
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

> 依次PID,PPID,进程名称,编译版本,进程路径

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

gops还有大量其它的功能

### assembly

SSA
Go 1.7开始，Go将原来的IR（Intermediate Representation，中间代码）转换成SSA（Static Single Assignment，静态单赋值）形式的IR
https://golang.design/gossa https://github.com/golang-design/ssaplayground

https://silverrainz.me/go-ssaviz/ https://github.com/SilverRainZ/go-ssaviz

汇编
https://godbolt.org/  https://github.com/compiler-explorer/compiler-explorer
https://github.com/loov/lensm
https://github.com/mmcloughlin/avo


AST
官方`go/ast`
```
package main

import (
  "go/ast"
  "go/parser"
  "go/token"
)

func main() {
  fs := token.NewFileSet()
  tr, _ := parser.ParseExpr("(3-1) * 5")
  ast.Print(fs, tr)
}
```

https://github.com/reflog/go2ast
https://github.com/xiazemin/ast_graph


## goreleaser 二进制包分发工具
go install github.com/goreleaser/goreleaser@latest

添加git 支持同时添加tag
```
git tag markdown/v1.0.5
git push -u origin main
git push -u origin markdown/v1.0.5
```
or
```
git tag -a markdown/v1.0.5 -m "markdown/v1.0.5"
git push origin markdown/v1.0.5
```

添加goreleaser 支持
```
goreleaser init
```

获取 Github Token
访问 [Settings / Developer Settings / Personal access tokens](https://github.com/settings/tokens)，点击 Generate new token 按钮，生成一个新的 Token，将 Token 保存到 `~/.config/goreleaser/github_token` 文件中 或者 `export GITHUB_TOKEN="YOUR_GH_TOKEN"`。

```
goreleaser release
```


goreleaser 的功能还是很强大的，同时支持github 的release，同时我们也可以配置docker


## 网络代理
- [**Caddy**](https://github.com/mholt/caddy) - 类似 Nginx 的 Web 服务器
- [Traefik](https://github.com/containous/traefik) - 反向代理&负载均衡
- [snail007/goproxy](https://github.com/snail007/goproxy) - golang 实现的高性能代理服务器
- [ProxyPool](https://github.com/henson/proxypool) - 采集免费的代理资源为爬虫提供有效的IP代理
- [frp](https://github.com/fatedier/frp) - 可用于内网穿透的高性能的反向代理应用
- [nps](https://github.com/cnlh/nps) - 一款轻量级、高性能、功能强大的内网穿透代理服务器
- [Pomerium](https://github.com/pomerium/pomerium) - 基于身份的反向代理
- [V2Ray](https://github.com/v2ray/v2ray-core)
- [V2Fly](https://github.com/v2fly/v2ray-core) - V2Ray 的社区版本
- [Tailscale](https://github.com/tailscale/tailscale) - WireGuard 解决方案
- [Clash](https://github.com/Dreamacro/clash) - 支持多种协议的多平台代理客户端
- [elazarl/goproxy](https://github.com/elazarl/goproxy) - HTTP 代理
- [oxy](https://github.com/vulcand/oxy) - Go middlewares for HTTP servers & proxies
- [ouqiang/goproxy](https://github.com/ouqiang/goproxy) - Go HTTP(S)代理库, 支持中间人代理解密HTTPS
- [pgrok](https://github.com/pgrok/pgrok) - 提供给穷人的内网穿透


## 网络相关
网络唤醒
https://github.com/seriousm4x/UpSnap

### tcpw
eBPF Talk: 使用 tcpw 获取 curl 的五元组信息
https://github.com/Asphaltt/tcpw/blob/main/README.md

### 抓包
https://github.com/google/gopacket
https://github.com/gopacket/gopacket

## http
### http中间件negroni
https://github.com/urfave/negroni
https://github.com/urfave/negroni#third-party-middleware

### HTTP cache/cachecontrol
HTTP caching proxy, implementing RFC 7234

https://github.com/pquerna/cachecontrol

https://github.com/darkweak/souin
https://github.com/darkweak/souin/blob/master/pkg/rfc

https://github.com/apache/trafficcontrol/tree/master/grove


### web框架/http server
- [gin](https://github.com/gin-gonic/gin)
- [Iris](https://github.com/kataras/iris)
- [echo](https://github.com/labstack/echo)
- [macaron](https://github.com/go-macaron/macaron)
- [mux](github.com/gorilla/mux)
- [Copper](https://github.com/gocopper/copper)
- [beego](https://github.com/beego/beego)
- [restful](https://github.com/emicklei/go-restful)

### fiber
https://github.com/gofiber/fiber
基于fasthttp
### hertz
https://github.com/cloudwego/hertz

## 文件
### 文件上传
tus（Terminated Uploads）
https://github.com/tus/tusd/tree/main

### vfs(Virtual File System)
https://github.com/rainycape/vfs
https://github.com/avfs/avfs
https://github.com/C2FO/vfs
https://github.com/spf13/afero

### 文件缓存
https://github.com/peterbourgon/diskv


## HTTP压测
- [Vegeta](https://github.com/tsenart/vegeta) - HTTP 负载压测工具
- [hey](https://github.com/rakyll/hey) - Web 压测工具
- [bombardier](https://github.com/codesenberg/bombardier) - Web 压测工具
- [go-wrk](https://github.com/tsliwowicz/go-wrk)
- [plow](https://github.com/six-ddc/plow)
- [Ddosify](https://github.com/ddosify/ddosify)

## 流量回放工具

### Goreplay
https://github.com/buger/goreplay/

[流量回放工具之 Goreplay 安装及初级使用](https://juejin.cn/post/6999586008698208263)

### Sharingan
https://kgithub.com/didi/sharingan

### 其它
- [tcpcopy](https://kgithub.com/session-replay-tools/tcpcopy)

## 分析网站使用的技术
### Wappalyzer
分析网站使用的技术
https://github.com/rverton/webanalyze/

https://github.com/wappalyzer/wappalyzer

### WhatRuns 

## 可视化
### Go Callvis
可视化 Go 程序的调用图
https://github.com/TrueFurby/go-callvis

## json
github.com/liamylian/json-hashids

序列化自动加密
### yaml/json转结构体
https://github.com/twpayne/go-jsonstruct

### encoding/json
https://github.com/bytedance/sonic
https://github.com/goccy/go-json

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

### Fyne
https://github.com/fyne-io/fyne 24.9k

Cross platform GUI in Go based on Material Design https://fyne.io/

跨平台支持手机端: `go run -tags mobile main.go`
```
fyne package -os android -appID my.domain.appname
fyne install -os android

fyne package -os ios -appID my.domain.appname
fyne package -os iossimulator -appID my.domain.appname
```
支持web
> If you’re using an older version of Go (<1.16), you should install fyne using `go get fyne.io/fyne/v2/cmd/fyne`
```
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne serve # 在线运行

fyne package -os web
```

```
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os darwin -icon myapp.png
fyne package -os linux -icon myapp.png
fyne package -os windows -icon myapp.png

# 在当前系统安装应用
fyne install -icon myapp.png
```


请注意，默认情况下，Windows应用程序是从命令提示符加载的，这意味着，如果单击图标，则可能会看到命令窗口。 要解决此问题，请在运行或构建命令中添加参数`-ldflags="-H windowsgui"`。

避免出现控制台:
`go run -ldflags="-H windowsgui"`

Prerequisites
https://fyne.io/develop/
Windows
1. Download Go from the download page and follow instructions
2. Install one of the available C compilers for windows, the following are tested with Go and Fyne:
    - MSYS2 with MingW-w64 - msys2.org
    - TDM-GCC - tdm-gcc.tdragon.net
    - Cygwin - cygwin.com
3. In Windows your graphics driver will already be installed, but it is recommended to ensure they are up to date.

环境安装
- Install Go
- Install Gcc
- `go get fyne.io/fyne/v2@latest` (or, if using Go before 1.16, then `go get fyne.io/fyne/v2`)
- You can test your installation using the [Fyne Setup](https://geoffrey-artefacts.fynelabs.com/github/andydotxyz/fyne-io/setup/latest/) app (检查环境: 检查你的环境是否可以进行fyne开发)

### wails
https://github.com/wailsapp/wails/ 25k

`go install github.com/wailsapp/wails/v2/cmd/wails@latest`

C/S模式，一个后端服务，一个前端页面作为UI。前端可以使用 Vue / React / Angular，可以说很适合偏前端的选手。

环境安装: https://wails.io/docs/gettingstarted/installation/
- Go 1.20+
- NPM (Node 15+)

> 可以运行`wails doctor`检查环境

创建项目: https://wails.io/docs/gettingstarted/firstproject
vue+ts: `wails init -n myproject -t vue-ts`


编译运行: https://wails.io/docs/guides/manual-builds/
`wails build` or `wails dev`
参数:
https://wails.io/docs/reference/cli/#build

### Gio
https://github.com/gioui/gio 1.7k
https://git.sr.ht/~eliasnaur/gio
https://gioui.org/

支持移动端和web端

`go install gioui.org/cmd/gogio@latest`
默认使用appicon.png做icon

开发: 
`go get gioui.org`


### webview
https://github.com/webview/webview 12.6k

Tiny cross-platform webview library for C/C++. Uses WebKit (GTK/Cocoa) and Edge WebView2 (Windows).

### 非go
#### Electron
https://github.com/electron/electron  C++ 86.5k -> 114k

#### TAURI
https://github.com/tauri-apps/tauri rust 7k -> 83.5k


## 脚手架/CLI tool
https://github.com/create-go-app/cli
https://github.com/go-nunu/nunu

## eval
github.com/Knetic/govaluate

## 微服务
https://github.com/douyu/jupiter

## book
《Go语法树入门》(开源免费图书/Go语言进阶/掌握抽象语法树/Go语言AST/LLVM/LLIR)
https://github.com/chai2010/go-ast-book


## log
- 日志收集
https://github.com/grafana/loki

- 变量调试神器
github.com/davecgh/go-spew

- wal(Write Ahead Log)
https://github.com/rosedblabs/wal
https://github.com/tidwall/wal

- log接口
https://github.com/rs/zerolog
https://github.com/uber-go/zap

- 日志分割/日志滚动
https://github.com/natefinch/lumberjack

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


## 代码质量检测工具
https://github.com/mgechev/revive

## tests
它是一个 Golang 命令行工具，它根据目标源文件的功能和方法签名生成表驱动测试。
https://github.com/cweill/gotests

## 优雅升级/平滑升级
https://lrita.github.io/2019/06/06/golang-graceful-upgrade/#so_reuseport-%E5%A4%9A%E8%BF%9B%E7%A8%8B

- facebookarchive/grace
- rcrowley/goagain
- jpillora/overseer
- cloudflare/tableflip

## grpc
### grpcui
他是一个 gRPC 的 Web 页面调试工具，提供交互式的调试界面。
https://github.com/fullstorydev/grpcui

grpcui -plaintext 127.0.0.1:9901
127.0.0.1:9901 是grpc server的地址

### bloomrpc
GUI Client for GRPC Services
https://github.com/uw-labs/bloomrpc

### insomnia
GraphQL、REST和gRPC的开源、跨平台API客户端。
https://github.com/Kong/insomnia

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

## js
### esbuild
https://github.com/evanw/esbuild
这是一个 JavaScript 打包和压缩程序。它用于打包 JavaScript 和 TypeScript 代码以在网络上分发。
我目前有两个基准测试用于衡量 esbuild 的性能。对于这些基准测试，esbuild 比我测试的其他 JavaScript 打包程序 快至少 100 倍 ：https://docs.breword.com/evanw-esbuild/

类似工具：
- Webpack ， Rollup 或 Parcel 用于打包
- Babel 或 TypeScript 用于转译
- Terser 或 UglifyJS 用于代码压缩

使用者：
- Vite
- Snowpack
> Vite，snowpack使用了esbuild

> esbuild 不可能替代 webpack、parcel 等构建工具

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

[golang将静态资源文件打包进二进制文件](https://zhuanlan.zhihu.com/p/351931501)
[golang1.16内嵌静态资源指南](https://www.cnblogs.com/apocelipes/p/13907858.html)

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

2020 年 10 月 30 日，Russ Cox 提交了最终的实现：[cmd/go: add //go:embed support](cmd/go: add //go:embed support)，意味着你在 tip 版本可以试用该功能了。Go1.16 版本会包含该功能。
Embed 设计提案：https://github.com/golang/proposal/blob/master/design/draft-embed.md
示例参考：https://github.com/mattn/go-embed-example
tip 相关文档：https://tip.golang.org


### go-bindata
https://github.com/go-bindata/go-bindata

## 视频下载

https://github.com/iawia002/annie
https://github.com/iawia002/lux

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

## go多版本
https://github.com/owenthereal/goup

## go测试
goc 是专为 Go 语言打造的一个综合覆盖率收集系统，尤其适合复杂的测试场景，比如系统测试时的代码覆盖率收集以及精准测试。
https://github.com/qiniu/goc

```
go tool cover -mode=count -var=CoverageVariableName xxxx.go
> 相信大家一定见过表示go覆盖率结果的coverprofile数据，类似下面: github.com/qiniu/goc/goc.go:21.13,23.2 1 1
其基本语义为 "文件:起始行.起始列,结束行.结束列 该基本块中的语句数量 该基本块被执行到的次数"
[聊聊 Go 代码覆盖率技术与最佳实践](https://xie.infoq.cn/article/ca1cc8ba293eddf793b3b0613)
```
## go Plugin
- go原生plugin
    - 实现机制基于动态链接so库
    - 跨语言支持较差且调用复杂，需解决不同语言的数据类型匹配问题

- github.com/hashicorp/go-plugin
    - 实现机制基于rpc调用，基于本地网络调用，调用性能高效
    - 插件可用多种语言实现，跨语言支持良好

## modules
https://github.com/nikolaydubina/go-recipes?tab=readme-ov-file#dependencies
### modgraphviz
golang.org/x/exp/cmd/modgraphviz
https://github.com/nikolaydubina/go-recipes?tab=readme-ov-file#-make-graph-of-upstream-modules-with-modgraphviz

### gmchart
github.com/PaulXu-cn/go-mod-graph-chart/gmchart

### import-graph
https://github.com/nikolaydubina/go-recipes?tab=readme-ov-file#-scrape-details-about-upstream-modules-and-make-graph-with-import-graph

## mock
- https://github.com/golang/mock
- https://github.com/brianvoe/gofakeit

## 机器学习
https://github.com/tensorflow/tensorflow/blob/master/tensorflow/go/README.md

[面部检测库](https://github.com/esimov/pigo)

## 流程图
https://github.com/blushft/go-diagrams
可以生成graphviz DOT file
go-diagrams实现了[diagrams](https://github.com/mingrammer/diagrams)的部分接口


## img/image
### Imagor
https://github.com/cshum/imagor
Imagor 是一个用 Go 编写的快速、支持 Docker 的图像处理服务器。
Imagor 使用最高效的图像处理库 libvips 之一（使用govips）。它通常比使用最快的 ImageMagick 和 GraphicsMagick 设置快4-8倍。
Imagor 是一个易于扩展的 Go 库，可以在任何 Unix 环境中安装和使用，并且可以使用 Docker 进行容器化。
Imagor 采用Thumbor URL 语法，涵盖了大多数 Web 图像处理用例。如果这些符合您的要求，Imagor 将是一种轻便、高性能的替代品。


## charts图表/plotting绘制/graphing绘图
### golang
https://github.com/vdobler/chart 695
https://github.com/Arafatk/glot 354
https://github.com/wcharczuk/go-chart 3.2k
https://github.com/gonum/plot 2k

### dotnet
https://github.com/ScottPlot/ScottPlot 1.1k
https://github.com/oxyplot/oxyplot 2.4k
https://github.com/Live-Charts/Live-Charts 5k
https://github.com/beto-rodriguez/LiveCharts2 771
https://github.com/microcharts-dotnet/Microcharts 1.6k
### 统计绘图
#### gnuplot
交互式绘图工具
http://www.gnuplot.info/

你可以在c#程序中编写数据文件，从c#调用gnuplot可执行文件，并在c#图片框中显示生成的图像。
#### Graph
http://www.padowan.dk/
#### Meta-chart.com
#### Infogr.am
https://infogram.com
#### Desmos
https://www.desmos.com/
#### Orange
https://orange.biolab.si/
#### GeoGebra
https://www.geogebra.org

#### AmCharts
https://www.amcharts.com



## 实时消息
https://github.com/topics/websocket?l=go
### centrifugo
https://github.com/centrifugal/centrifugo

> (WebSocket, HTTP-streaming, SSE/EventSource, GRPC, SockJS, WebTransport)

### melody
https://github.com/olahol/melody

## 扩展
### net/mem/cpu/disk
https://github.com/shirou/gopsutil

#### psutil
psutil是一个开源且跨平台的库，其提供了便利的函数用来获取操作系统的信息，如cpu、内存、磁盘、网络等信息。此外，psutil还可以用来进行进程管理，包括判断进程是否存在、获取进程列表、获取进程的详细信息等。psutil广泛应用于系统监控、进程管理、资源限制等场景。此外，psutil还提供了许多命令行工具提供的功能，包括ps, top, lsof, netstat, ifconfig, who, df, kill, free, nice, ionice, iostat, iotop, uptime, pidof, tty, taskset, pmap。有python、go、rust等版本。

### 配置文件处理库
[配置文件处理库](https://github.com/spf13/viper)

[命令行接口](https://github.com/spf13/cobra)
[环境变量](https://github.com/joho/godotenv)

[日期管理](https://github.com/golang-module/carbon)


### mapper
https://github.com/petersunbag/coven 27 (有用于生产经验)

这个也是很NB
https://github.com/jinzhu/copier  5.5k (有用于生产经验)


https://github.com/imdario/mergo  2.9k

https://github.com/mitchellh/mapstructure 7.9k

### collection/集合/容器
https://github.com/zyedidia/generic

- array2d: a 2-dimensional array.
- avl: an AVL tree.
- bimap: a bi-directional map; a map that allows lookups on both keys and values.
- btree: a B-tree.
- cache: a wrapper around map[K]V that uses a maximum size and evicts elements using LRU when full.
- hashmap: a hashmap with linear probing. The main feature is that the hashmap can be efficiently copied, using copy-on-write under the hood.
- hashset: a hashset that uses the hashmap as the underlying storage.
- heap: a binary heap.
- interval: an interval tree, implemented as an augmented AVL tree.
- list: a doubly-linked list.
- mapset: a set that uses Go's built-in map as the underlying storage.
- multimap: an associative container that permits multiple entries with the same key.
- queue: a First In First Out (FIFO) queue.
- rope: a generic rope, which is similar to an array but supports efficient insertion and deletion from anywhere in the array. Ropes are typically used for arrays of bytes, but this rope is generic.
- prope: a persistent version of the rope, which allows for keeping different versions of the rope with only a little extra time or memory.
- stack: a LIFO stack.
- trie: a ternary search trie.
- ulist: an un-rolled doubly-linked list.

- [字典树 (Trie)](https://github.com/dghubble/trie)


- [collection](https://github.com/bytedance/gopkg/tree/main/collection)
### 文档处理
#### 转PDF
##### Gotenberg 
https://github.com/gotenberg/gotenberg

### 工具库
#### 优秀的实验库
https://github.com/smallnest/exp/

#### samber/lo 
https://github.com/samber/lo


DI依赖注入
https://github.com/samber/do


Functional Programming
https://github.com/samber/mo
https://github.com/TeaEntityLab/fpGo
https://github.com/IBM/fp-go

#### lancet
https://github.com/duke-git/lancet

### hash

github.com/dchest/siphash
github.com/cespare/xxhash/v2


### 类型转换cast
github.com/spf13/cast

### 比较两个值是否相等go-cmp
github.com/com/google/go-cmp/cmp
github.com/com/google/go-cmp

### Pretty printing for Go values
github.com/kr/pretty 
### 字符串format

https://github.com/sirkon/go-format 功能最强大
```
type t struct {
	A     string
	Field int
}
var s = t{
	A:     "str",
	Field: 12,
}
var d struct {
	F     t
	Entry float64
}
d.F = s
d.Entry = 0.5
res := format.Formatg("${F.A} ${F.Field} $Entry", d)
// res = "str 12 0.500000"
```

https://github.com/chonla/format

```
var params = map[string]interface{}{
    "sister": "Susan",
    "brother": "Louis",
}
format.Printf("%<brother> loves %<sister>.", params)
```

https://github.com/go-ffmt/ffmt/blob/master/format.go
```
Format("hello {name}", "ffmt") to "hello ffmt"
```

https://github.com/Wissance/stringFormatter
```
FormatComplex("Hello {user} what are you doing here {app} ?", map[string]interface{}{"user":"vpupkin", "app":"mn_console"})
```

https://github.com/delicb/gstring
```
// outpits "Bracket {, }, key value, key value, key value"
gstring.Printm("Bracket {{, }}, {key}, {key}, {key}", map[string]interaface{}{"key", "key value"})
```

https://github.com/mgenware/go-string-format
```
strf.Format("🐆{0} {1} {0}", "leopard", -3)
// returns "🐆leopard -3 leopard" 
```

https://github.com/taloric/strfmt
```
format_string := "Today is a {what} {desc}"
args := make(map[string]string)
args["what"] = "wonderful"
args["desc"] = "day"
res,err := strfmt.FormatMap(format_string, &args)
fmt.Println(res)
```

https://github.com/yangchenxing/go-string-mapformatter
```
fmt.Println(mapformatter.MustFormat("Hello %(name|s), you owe me %(money|.2f) dollar.",
    map[string]interface{}{
        "name": "anyone",
        "money": 10.3,
    }))
// Output: Hello anyone, you owe me 10.30 dollar.
```

https://github.com/christianhujer/tfmt/blob/master/main.go


https://golangdocs.com/string-formatting-in-golang
https://zetcode.com/golang/string-format/

### go-humanize
https://github.com/dustin/go-humanize
```
fmt.Printf("That file is %s.", humanize.Bytes(82854982)) // That file is 83 MB.
fmt.Printf("This was touched %s.", humanize.Time(someTimeInstance)) // This was touched 7 hours ago.
```
### dateparse
支持多种格式
https://github.com/araddon/dateparse

### 驼峰、下划线等转换
https://github.com/WnP/go-sfmt


## 其它
### Tools工具
[Tools for Go projects](https://github.com/nikolaydubina/go-recipes)
### 隔离环境的命令行工具
https://github.com/jetpack-io/devbox

### 数据库迁移工具
https://github.com/golang-migrate/migrate

### 搜索引擎
https://github.com/prabhatsharma/zinc
### 一款用 SQL 方式查询 Git 仓库的开源项目进入 GitHub 趋势榜
https://github.com/augmentable-dev/gitqlite

### Webmail 服务器（SMTP、POP3、email）
https://github.com/inbucket/inbucket

### SSH

[massh](https://github.com/DiscoRiver/massh)：通过 SSH 方式运行 Linux 分布式 Shell 命令。

[sh](https://github.com/mvdan/sh): 一个支持 Bash 的 Shell 解析器、格式化器。