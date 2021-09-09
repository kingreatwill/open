# Go 1.15 都有哪些值得关注的变化
https://tip.golang.org/doc/go1.15

正式版本：https://golang.org/doc/go1.15

## 运行时 Runtime
- Converting a small integer value into an interface value no longer causes allocation.
意思是说，将小整数转换为接口值不再需要进行内存分配。小整数是指 0 到 255 之间的数。

```go
package smallint

func Convert(val int) []interface{} {
    var slice = make([]interface{}, 100)
    for i := 0; i < 100; i++ {
        slice[i] = val
    }

    return slice
}

package smallint_test

import (
    "testing"
    "test/smallint"
)

func BenchmarkConvert(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := smallint.Convert(12)
        _ = result
    }
}
```
分别使用 Go1.14 和 Go1.15 版本进行测试,接着讲 smallint_test.go 中调用 Convert 的参数由 12 改为 256，再次使用 Go1.15 运行.


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

# go1.17
[草案](https://github.com/godghdai/deployment-logs/blob/main/Go/Type%20Parameters%20Proposal.md)
使用是 -G 标识做为泛型的开关。

计划如下：

-G=0：继续使用传统的类型检查器。
-G=1：使用 type2，但不支持泛型。
-G=2：使用 type2，支持泛型。
在完成 types2 的错误和现有的错误的开发协调后，计划在 Go 1.17 将 -G=1 设置为默认值。

未来也许可以在 Go 1.18 中放弃对 -G=0 的支持，这样后续在默认启用 -G=2 上会变得更容易。
```
package main

import (
    "fmt"
)

type Addable interface {
type int, int8, int16, int32, int64,
    uint, uint8, uint16, uint32, uint64, uintptr,
    float32, float64, complex64, complex128,
    string
}

func add[T Addable](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(add(1,2))
    fmt.Println(add("1", "2"))
} 
```
`go run -gcflags=all=-G=3 main.go`
`go run -gcflags=-G=3 main.go`