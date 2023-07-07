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

# go1.18

[工作区模式](https://mp.weixin.qq.com/s/pR2rOWM2cgOk191IbifwGQ)

# go1.21

- 低版本的go编译器将拒绝编译高版本的go module(go.mod中go version标识最低版本) 

- 支持WASI(WebAssembly System Interface)
Go从1.11版本就开始支持将Go源码编译为wasm二进制文件，并在支持wasm的浏览器环境中运行。

不过WebAssembly绝不仅仅被设计为仅限于在Web浏览器中运行，核心的WebAssembly语言是独立于其周围环境的，WebAssembly完全可以通过API与外部世界互动。在Web上，它自然使用浏览器提供的现有Web API。然而，在浏览器之外，之前还没有一套标准的API可以让WebAssembly程序使用。这使得创建真正可移植的非Web WebAssembly程序变得困难。WebAssembly System Interface(WASI)是一个填补这一空白的倡议，它有一套干净的API，可以由多个引擎在多个平台上实现，并且不依赖于浏览器的功能（尽管它们仍然可以在浏览器中运行）。

Go 1.21将增加对WASI的支持，初期先支持WASI Preview1版本，之后会支持WASI Preview2版本，直至最终WASI API版本发布！目前我们可以使用GOOS=wasip1 GOARCH=wasm将Go源码编译为支持WASI的wasm程序，下面是一个例子：
```go
// main.go
package main            
                        
func main() {           
    println("hello")    
}   
```

```
$ GOARCH=wasm GOOS=wasip1 gotip build -o main.wasm main.go
```

开源的wasm运行时有很多，[wazero](https://wazero.io/)是目前比较火的且使用纯Go实现的wasm运行时程序，安装wazero后，可以用来执行上面编译出来的main.wasm：
```
$curl https://wazero.io/install.sh
$wazero run main.wasm
hello 
```

- 增加$GOROOT/go.env
Go 1.21将增加一个全局层次上的go.env，放在$GOROOT下面，目前默认的go.env为：
```
// $GOROOT/go.env

# This file contains the initial defaults for go command configuration.
# Values set by 'go env -w' and written to the user's go/env file override these.
# The environment overrides everything else.

# Use the Go module mirror and checksum database by default.
# See https://proxy.golang.org for details.
GOPROXY=https://proxy.golang.org,direct
GOSUMDB=sum.golang.org
```
我们仍然可以通过go env -w命令修改user级的env文件来覆盖上述配置，当然最高优先级的是OS用户环境变量，如果在OS用户环境变量文件(比如.bash_profile、.bashrc)中设置了Go的环境变量值，比如GOPROXY等，那么以OS用户环境变量为优先。

- **在 1.20 中处于预览阶段的启用配置文件引导优化 (PGO) 功能现已正式 GA。**
Profile-guided optimization (PGO) 是计算机编程中的一种编译器优化技术, 使用配置文件引导的优化。
如果主软件包目录中存在名为default.pgo的文件，go命令将使用它来启用 PGO 构建。

PGO的原理很简单，那就是先把程序跑起来，收集程序运行过程中的数据。然后编译器再根据收集到的程序运行时数据来分析程序的行为，进而做针对性的性能优化。

```
$ curl -o cpu.pprof "http://localhost:8080/debug/pprof/profile?seconds=30"
$ mv cpu.pprof default.pgo
$ go build -pgo=auto -o withpgo
```

> -pgo既可以支持指定的profiling文件，也可以支持auto模式。
如果是auto模式，会自动寻找程序主目录下名为default.pgo的profiling文件。
Go 1.20版本里，-pgo选项的默认值是off，我们必须添加-pgo=auto来开启PGO优化。
https://go.dev/doc/pgo