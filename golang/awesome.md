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