
![](https://github.com/polaris1119/golangweekly/raw/master/docs/imgs/issue054/history-of-go.png)


## 安装
1. 下载`wget -c https://dl.google.com/go/go1.20.5.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local`

2. 环境变量
这个可以通过添加下面的行到/etc/profile文件（系统范围内安装）或者$HOME/.profile文件（当前用户安装）：`export PATH=$PATH:/usr/local/go/bin`
保存文件，并且重新加载新的PATH 环境变量到当前的 shell 会话：`source ~/.profile`

设置`go env -w GOPROXY=https://goproxy.cn,direct`

安装 & 升级
```
wget https://golang.google.cn/dl/go1.21.4.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
vi /etc/profile
export PATH=$PATH:/usr/local/go/bin
source /etc/profile
```


## 在现有go版本中安装其他版本
https://pkg.go.dev/golang.org/dl
go get golang.org/dl/go1.10.7
go get golang.org/dl/go1.13.6

1. go get github.com/golang/dl/go1.13.9
2. go1.13.9 download 
3. go1.13.6 version

原理：go get github.com/golang/dl 库
每个版本install

会在 go/bin 下安装github.com/golang/dl/go1.13.9 并生成 go1.13.9.exe
go1.13.9 download 下载源码到用户目录下的sdk/go1.13.9 下

运行go1.13.6 version 相当于运行 sdk/go1.13.9/bin/go.exe version
### 使用docker编译不同版本
`docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.17 go build -v`

`-w, --workdir string`       指定工作目录


## 动态链接和静态链接
```
$file helloworld-default
helloworld-default: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
$ldd helloworld-default
   不是动态可执行文件
```

开启 CGO_ENABLED=1
```
// go-compilation/main-with-os-user.go

package main

import (
 "fmt"
 _ "os/user"
)

func main() {
 fmt.Println("hello, world")
}
```

```
$file helloworld-with-os-user
helloworld-with-os-user: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), dynamically linked (uses shared libs), not stripped

$ldd helloworld-with-os-user
    linux-vdso.so.1 =>  (0x00007ffcb8fd4000)
    libpthread.so.0 => /lib64/libpthread.so.0 (0x00007fb5d6fce000)
    libc.so.6 => /lib64/libc.so.6 (0x00007fb5d6c00000)
    /lib64/ld-linux-x86-64.so.2 (0x00007fb5d71ea000)
```
通过nm命令我们还可以查看Go程序依赖了哪些C库的符号：
```
$nm -a helloworld-with-os-user |grep " U "
                 U abort
                 U __errno_location
                 U fprintf
                 U fputc
                 U free
                 U fwrite
                 U malloc
                 U mmap
                 U munmap
                 U nanosleep
                 U pthread_attr_destroy
                 U pthread_attr_getstack
                 U pthread_attr_getstacksize
                 U pthread_attr_init
                 U pthread_cond_broadcast
                 U pthread_cond_wait
                 U pthread_create
                 U pthread_detach
                 U pthread_getattr_np
                 U pthread_key_create
                 U pthread_mutex_lock
                 U pthread_mutex_unlock
                 U pthread_self
                 U pthread_setspecific
                 U pthread_sigmask
                 U setenv
                 U sigaction
                 U sigaddset
                 U sigemptyset
                 U sigfillset
                 U sigismember
                 U stderr
                 U strerror
                 U unsetenv
                 U vfprintf
```

### 依赖使用cgo的外部go包（静态链接）
要想实现静态链接，我们需要找出外部go依赖的所有c库的.a文件(静态共享库)。以我们的go-sqlite3示例为例，go-sqlite3是sqlite库的go binding，它依赖sqlite库，同时所有第三方c库都依赖libc，我们还要准备一份libc的.a文件，下面我们就先安装这些：
```
$yum install -y gcc glibc-static sqlite-devel 
... ...

已安装:
  sqlite-devel.x86_64 0:3.7.17-8.el7_7.1                                                                                          

更新完毕:
  glibc-static.x86_64 0:2.17-326.el7_9.3            
```
接下来，我们就来以静态链接的方式在go-compilation/go-sqlite3-static下编译一下：
```
$go build -tags 'sqlite_omit_load_extension' -ldflags '-linkmode external -extldflags "-static"' demo

$file ./demo
./demo: ELF 64-bit LSB executable, x86-64, version 1 (GNU/Linux), statically linked, for GNU/Linux 2.6.32, BuildID[sha1]=c779f5c3eaa945d916de059b56d94c23974ce61c, not stripped
```
这里命令行中的`-tags 'sqlite_omit_load_extension'`用于禁用SQLite3的动态加载功能，确保更好的静态链接兼容性。而`-ldflags '-linkmode external -extldflags "-static"'`的含义是使用外部链接器(比如gcc linker)，并强制静态链接所有库。

### 优化可执行文件的大小
- 去除符号表和调试信息
在编译时使用-ldflags="-s -w"标志可以去除符号表和调试符号，其中-s用于去掉符号表和调试信息，-w用于去掉DWARF调试信息，这样能显著减小文件体积。
```
$go build -ldflags="-s -w" -o helloworld-default-nosym main.go
```


- 使用tinygo
[TinyGo](https://github.com/tinygo-org/tinygo)是一个Go语言的编译器，它专为资源受限的环境而设计，例如微控制器、WebAssembly和其他嵌入式设备。TinyGo的目标是提供一个轻量级的、能在小型设备上运行的Go运行时，同时尽可能支持Go语言的特性。tinygo的一大优点就是生成的二进制文件通常比标准Go编译器生成的文件小得多：
```
$tinygo build -o helloworld-tinygo main.go
```

### 依赖CGO依赖CGO
```
go list -deps -f '{{.ImportPath}}: {{.CgoFiles}}' ./...  | grep -v '\[\]'

go list -deps -f '{{.ImportPath}}: {{.CgoFiles}}'  main-with-os-user.go | grep -v '\[\]'
```


## 其它
https://go.dev/

https://godoc.org/

GolangCI-Lint 检查代码
自带go vet命令 检查代码
> https://github.com/timakin/bodyclose

更严格的gofmt https://github.com/mvdan/gofumpt

检查没有使用的函数https://github.com/mvdan/unparam

## Gopher China
https://github.com/fanjindong/GopherChina

## 字节对齐/内存对齐

[在线查看内存对齐](https://github.com/kingreatwill/go-sizeof-tips) 在线[查看](https://golang-sizeof.wcoder.com/)


[Go的内存对齐和指针运算详解和实践](https://www.cnblogs.com/sy270321/p/12205435.html)

[Go 的内存对齐和指针运算详解和实践](https://learnku.com/articles/39255)

[深挖Go系列之聊聊struct的内存对齐](https://mp.weixin.qq.com/s/TzY7aWbKtikQpfIySkJeQQ)

[内存对齐工具](https://github.com/dominikh/go-tools)

## 待整理

[推荐算法、相似度算法、布隆过滤器、均值算法、一致性Hash、数据结构、leetcode练习](https://github.com/hwholiday/algorithm_coding)

[Go 101](https://go101.org/article/101.html)
[Go 101 Github](https://github.com/go101/go101)
[Go 101 中文](https://gfw.go101.org)
[Go 101 中文Github](https://github.com/golang101/golang101)

[《Go语言编程》高清完整版电子书.pdf](../files/go/《Go语言编程》高清完整版电子书.pdf)

[golang参考GO知识图谱](https://www.toutiao.com/a1653862521797643)

[Go内存泄漏？不是那么简单!](https://colobu.com/2019/08/28/go-memory-leak-i-dont-think-so/)

[内存泄漏方案](https://go101.org/article/memory-leaking.html)

[Uber 编程指南](https://github.com/uber-go/guide)

[Go 编程语言规范原文](https://golang.org/ref/spec)

[Go 编程语言规范【译】](https://moego.me/golang_spec.html)

[高频 golang 服务接口超时排查&性能调优](https://mp.weixin.qq.com/s/rDjTqqR0q4VTSQrYFzbR7w)


[内存布局](https://gfw.go101.org/article/memory-layout.html)

[Memory Blocks](https://go101.org/article/memory-block.html)


[[译]Go：内存管理与内存分配](https://juejin.im/post/5ddcdc5df265da05c33fcad2)

[带你领略Go源码的魅力----Go内存原理详解](https://juejin.im/post/5ddcf31751882572d8212481)

[go内存模型](https://cloud.tencent.com/developer/article/1359184)

[Go内存分配那些事，就这么简单！](https://www.cnblogs.com/shijingxiang/articles/11466957.html)

[Go性能优化之路.pdf](../files/go/Go性能优化之路.pdf)

[编写和优化Go代码](https://github.com/dgryski/go-perfbook/blob/master/performance-zh.md)

[编写可维护的Go程序的建议](https://dave.cheney.net/practical-go/presentations/qcon-china.html)

[Go工具概述](https://www.alexedwards.net/blog/an-overview-of-go-tooling)

[Go 学习、Go 进阶、Go 实用工具类](https://github.com/hwholiday/learning_tools)

[Go性能提升快速指南](https://stephen.sh/posts/quick-go-performance-improvements)

[深入浅出Golang Runtime.ppt](../files/go/gopher-meetup-深入浅出Golang-Runtime-yifhao-full.pptx)

[深入浅出Golang Runtime](https://www.acao.cn/course/program/402.html)

[编译出dll C# 结合 Golang 开发](https://www.cnblogs.com/timeddd/p/11731160.html)

[详解 Go 内联优化](https://mp.weixin.qq.com/s/89QLSqGzH0yUM6cO6V4zmQ)

![](img/go.jpg)




