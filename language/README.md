[TOC]

![](../img/var-named.jpg)

### go build dll
[编译出 dll](https://www.cnblogs.com/timeddd/p/11731160.html)

```
go build --buildmode=c-shared -o Test.dll
dotnet
[DllImport(DLL_NAME, EntryPoint = "Test")]
```
### 六大主流语言代码漏洞分析
https://www.veracode.com/sites/default/files/pdf/resources/ipapers/security-flaw-heatmap/index.html

### 多线程

.NET 提供了两种机制，用于将线程本地存储 (Thread Local Storage,TLS) ：thread-relative static fields, and data slots.
- [ThreadStaticAttribute](https://docs.microsoft.com/en-us/dotnet/api/system.threadstaticattribute?view=net-5.0)
- [LocalDataStoreSlot](https://docs.microsoft.com/en-us/dotnet/api/system.localdatastoreslot?view=net-5.0)

dotnet的ThreadStatic对应java的ThreadLocal

dotnet的LocalDataStoreSlot

java volatile

### 让出 CPU 时间片

#### GO

用于让出 CPU 时间片
runtime.Gosched()

#### .net

Thread.Sleep()方法：是强制放弃 CPU 的时间片，然后重新和其他线程一起参与 CPU 的竞争。
用 Sleep()方法是会让线程放弃 CPU 的使用权。
用 SpinWait()方法是不会放弃 CPU 的使用权。

#### .net

await Task.Yield();// 让出时间片
Thread.SpinWait();// 不让出时间片

AsyncLocal 4.6+

4.5
using System.Runtime.Remoting;
using System.Runtime.Remoting.Messaging;
CallContext

#### java

Thread.yield(); //让出当前剩余的 CPU 时间片
Thread.onSpinWait();// 不让出时间片

TransmittableThreadLocal

## grpc Balance

```
conn, err := grpc.Dial(
    "",
    grpc.WithInsecure(),
    grpc.WithBalancer(grpc.RoundRobin(resolver.NewPseudoResolver([]string{
        "10.0.0.1:10000",
        "10.0.0.2:10000",
    }))),
)
```

## 代码格式化

### clang-format

下载https://github.com/llvm/llvm-project/releases

```
# win64
https://github.com/llvm/llvm-project/releases/download/llvmorg-10.0.0/LLVM-10.0.0-win64.exe

# ubuntu/debian
sudo apt-get install clang-format
```

文档
http://clang.llvm.org/docs/ClangFormat.html
支持 C/C++/Java/JavaScript/Objective-C/Protobuf/C#

代码格式支持：LLVM, Google, Chromium, Mozilla, WebKit
http://clang.llvm.org/docs/ClangFormatStyleOptions.html

[团队协作-代码格式化工具 clang-format](https://www.toutiao.com/i6886080589141639693)

Visual Studio Code 安装插件 clang-format

其它插件

- markdown-formatter
- shell-format
- Prettier - Code formatter

shift+alt+f
or
shift+ctrl+p

### Prettier
https://github.com/prettier

支持
- JavaScript (including experimental features)
- JSX
- Angular
- Vue
- Flow
- TypeScript
- CSS, Less, and SCSS
- HTML
- JSON
- GraphQL
- Markdown, including GFM and MDX
- YAML

## 命令行服务器(最小web服务器)/Simple Web Server
### rust
miniserve（最常用，功能全）
https://github.com/svenstaro/miniserve
cargo install miniserve
miniserve . -p 8080


simple-http-server（轻量）
https://github.com/TheWaWaR/simple-http-server
cargo install simple-http-server
simple-http-server -p 8080 -i 0.0.0.0 folder-name

mdBook（文档站点带内置 serve）
https://github.com/rust-lang/mdBook
https://rust-lang.github.io/mdBook/guide/installation.html
cargo install mdbook
mdbook serve


Trunk（Yew/Leptos 等）：
https://github.com/trunk-rs/trunk
cargo install trunk
trunk serve

### dotnet
https://github.com/natemcmaster/dotnet-serve
```
dotnet tool install --global dotnet-serve

dotnet serve -o
```

### jdk 18
经过一周的评审，[JEP 408](https://openjdk.java.net/jeps/408)，也就是 Simple Web Server 由 JDK 18 的 Proposed to Target 状态进入到了 Targeted。这个 JEP 提供了一个基于 HTTP 命令行的、最小化的、只提供静态文件的 Web 服务器。这个工具主要用于构建原型、临时编码和测试，特别是在培训环境中。这个 Web 服务器可以通过以下命令来启动：
`java -m jdk.httpserver [-b bind address] [-p port] [-d directory] [-h to show help message] [-o none|default|verbose]`

### python
#### python2
```
python -m SimpleHTTPServer 8888
python -m SimpleHTTPServer 8888 &
nohup python -m SimpleHTTPServer 8000 &
```
#### python3
```
python3 -m http.server 8888
python3 -m http.server 8888 &
nohup python3 -m http.server 8888 &
```

### node
```
npm install -g http-server

http-server <path> -a 0.0.0.0 -p 8080
```

### golang
#### devd
https://github.com/cortesi/devd
```
go install github.com/cortesi/devd/cmd/devd@latest

devd -w ./src http://localhost:8080

# Serve the current directory, open it in the browser (-o), and livereload when files change (-l):
devd -ol .
```
#### Swego
https://github.com/nodauf/Swego

#### algernon
https://github.com/xyproto/algernon

纯 Go 语言开发的 Web 开发工具，内置 Lua，Markdown，模板和 HTTP / 2，Redis，Mysql，Postgresql 支持

小型、独立的 Go Web 服务器。该项目是用 Go 编写的“快餐” Web 服务器，采用 BoltDB、Redis、MySQL 或 PostgreSQL 作为数据库，内置 Lua 解释器。所有功能全在一个独立可执行文件中，支持 Markdown 渲染、Lua 脚本、请求限制、用户和权限等。

#### gohttpserver
https://github.com/codeskyblue/gohttpserver

## 内存泄露/内存错误检测工具
### Address Sanitizer(ASan)
Address Sanitizer是谷歌的快速的内存错误检测工具，它非常快只拖慢程序2倍左右的速度，在这次使用过程中，也是深有体会。在GCC 4.9版本以上，就可以很好的使用了。

Sanitizers是谷歌发起的开源工具集，包括了AddressSanitizer, MemorySanitizer, ThreadSanitizer, LeakSanitizer，Sanitizers项目本是LLVM项目的一部分，但GNU也将该系列工具加入到了自家的GCC编译器中。GCC从4.8版本开始支持Address和Thread Sanitizer，4.9版本开始支持Leak Sanitizer和UB Sanitizer，这些都是查找隐藏Bug的利器。


编译时候添加选项：`-fsanitize=address  -fno-omit-frame-pointer -fno-optimize-sibling-calls  -O0`
```
g++  -fsanitize=address  -fno-omit-frame-pointer -fsanitize=leak    -use-after-free  -g main.c -o t1

gcc -g main.c -o t1  -fsanitize=leak -fsanitize=address  -fno-omit-frame-pointer

clang    -fsanitize=address  -fno-omit-frame-pointer -fsanitize=leak    -use-after-free  -g main.c -o t1
ASAN_OPTIONS=detect_leaks=1 ./t1
```

> yum install libasan -y
> yum install centos-release-scl-rh
> yum --enablerepo=centos-sclo-rh-testing install libasan5