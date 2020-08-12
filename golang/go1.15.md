# Go 1.15 都有哪些值得关注的变化
https://tip.golang.org/doc/go1.15

正式版本：https://golang.org/doc/go1.15

## 运行时 Runtime


## 编译器 Compiler
与Go 1.14相比，通过消除某些类型的GC元数据和更积极地消除未使用的类型元数据，Go 1.15减少了典型的二进制文件大小约5%。

经过一个服务的测试：减少了3.8%，还是很不错的。

Go 1.15 adds a "-spectre" flag to enable Spectre mitigations for the compiler and assembler.

## 新的链接器 Linker
官方的设计文档地址：https://golang.org/s/better-linker，从命名看，是一个更好的链接器（这是废话）。

Go链接器现在使用的资源更少，速度更快，并且改进了代码质量。一般来说，对于大型Go程序，链接的速度大约快了20%，而占用的内存却少了30%左右。

## 略微小了些的二进制文件

```
$ go version
go version go1.14.2 linux/amd64
$ gotip version
go version devel +49f10f3797 Sat Apr 25 02:19:12 2020 +0000 linux/amd64
$ go build -ldflags='-w -s' -o 114 
$ gotip build -ldflags='-w -s' -o 115
$ du -b 114 115
5111808 114
5009408 115
```
2% 左右，虽然不多，但还在考虑进一步减小。
> gotip 新版本https://golang.org/dl/gotip

## 内嵌 tzdata（时区数据）

增加了一个新包：time/tzdata，当系统找不到时区数据时（比如 Windows 等），通过导入这个包，在程序中内嵌时区数据，也可以通过编译时传递  -tags timetzdata 来实现同样的效果。

具体查看这个 issue：https://github.com/golang/go/issues/38017 以及包 time/tzdata 的说明：https://tip.golang.org/pkg/time/tzdata/。

##4、增加 testing.TB.TempDir

测试生成临时文件挺常见的，这个为了更好的解决此问题。详情见 issue：https://github.com/golang/go/issues/35998。

## 增加 testing.T.Deadline

将 context 引入 testing 包。详情见 issue：https://github.com/golang/go/issues/28135。

## 关于 Ports 部分

darwin/386、darwin/arm 不再支持；riscv64 变得更好；linux/arm64 现在作为第一类 port 支持。

## API  的变动

1）net/url.URL RawFragment 和 EscapedFragment ，详情见 issue：https://github.com/golang/go/issues/37776；

2）net/url.URL.Redacted，详情见 issue：https://github.com/golang/go/issues/34855；

3）time.Ticker.Reset，我们知道 Timer 是有 Reset 的，这次为 Ticker 也增加，详情见 issue：https://github.com/golang/go/issues/33184；

4）regexp.Regexp.SubexpIndex，详情见 issue：https://github.com/golang/go/issues/32420；

5）sync.Map.LoadAndDelete，详情见 issue：https://github.com/golang/go/issues/33762；

6）crypto/tls.Dialer.DialContext，详情见 issue：https://github.com/golang/go/issues/18482；

还有其他一些 API 变动，不一一列举。

## 工具链

1）增加 go env GOMODCACHE：https://github.com/golang/go/issues/34527；

2）opt-in fallbacks in GOPROXY：https://github.com/golang/go/issues/37367；

3）vet：warn about string(int) 和 detect impossible interface assertions：https://github.com/golang/go/issues/32479 和 https://github.com/golang/go/issues/4483；

4）println 允许打印两个值。println(twoValues())；

5）panic：显示可打印的值而不是地址。比如：
```
type MyString string
panic(MyString("hello"))

现在打印：
panic: (main.MyString) (0x48aa00,0x4c0840)

期望打印：
panic: main.MyString("hello")
```

## 性能

1）在 amd64 上更好的写屏蔽；

2）在 Linux 上，forkAndExec 使用 dup3；

3）sha512 算法速度提升 15%；

4）ReadMemStats 延迟降低 95%；

5）关闭状态的 channel 接收速度提升 99%;

6）将小的 int 值转为 interface{} 不额外分配内存；

## 如何试用？
```
$ go install golang.org/dl/gotip
$ gotip download
```