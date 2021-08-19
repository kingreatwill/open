# Profiling 
## pyroscope
https://github.com/pyroscope-io/pyroscope

- Ruby (via rbspy)
- Python (via py-spy)
- Go (via pprof)
- Linux eBPF (via `profile.py` from bcc-tools)
- PHP (via phpspy)
- .NET (via dotnet trace)
- Java (coming soon)

## pprof
[深度解密Go语言之 pprof](https://qcrao.com/2019/11/10/dive-into-go-pprof/)

[深度解密Go语言之 pprof](https://www.cnblogs.com/qcrao-2018/p/11832732.html)

https://github.com/golang/go/wiki/Performance




gops 分析机器上运行了哪些go进程
https://shockerli.net/post/golang-tool-gops/

godebug:一个跨平台的Go程序调试工具
go tool trace


查看运行时间，详细
/usr/bin/time -v go run test2.go


go tool pprof -http=:1234 http://10.244.28.10:8080/debug/pprof/profile?seconds=30 
go tool pprof -http=:1234 http://10.244.28.10:8080/debug/pprof/allocs?seconds=30

- 分析工具：GODEBUG
- 分析工具：go tool pprof
- 分析工具：go tool trace
[好东西 - 性能优化](file/gopher-meetup-2020-陈一枭.pdf)




堆的信息
```
go tool pprof -http :9090 http://ip:port/debug/pprof/heap
```

cpu（CPU Profiling）: $HOST/debug/pprof/profile，默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件
block（Block Profiling）：$HOST/debug/pprof/block，查看导致阻塞同步的堆栈跟踪
goroutine：$HOST/debug/pprof/goroutine，查看当前所有运行的 goroutines 堆栈跟踪
heap（Memory Profiling）: $HOST/debug/pprof/heap，查看活动对象的内存分配情况
mutex（Mutex Profiling）：$HOST/debug/pprof/mutex，查看导致互斥锁的竞争持有者的堆栈跟踪
threadcreate：$HOST/debug/pprof/threadcreate，查看创建新OS线程的堆栈跟踪


pprof可以比较两个时间点的分配的内存的差值

1. 首先确保你已经配置了 pprof 的 http 路径， 可以访问`http://ip:port/debug/pprof/`查看(如果你没有修改默认的 pprof 路径)
2. 导出**时间点 1**的堆的 profile: `curl -s http://127.0.0.1:8080/debug/pprof/heap > base.heap`, 我们把它作为基准点
3. 喝杯茶，等待一段时间后导出**时间点 2**的堆的 profile: `curl -s http://127.0.0.1:8080/debug/pprof/heap > current.heap`
4. 现在你就可以比较这两个时间点的堆的差异了: `go tool pprof --base base.heap current.heap`

5. 使用web命令会生成一个 SVG 文件，可能你需要使用浏览器打开它。
或者你直接使用命令打开 web 界面: go tool pprof --http :9090 --base base.heap current.heap。


通过比较差值，就容易看到哪些地方产生的内存"残留"的比较多，没有被内存释放，极有可能是内存泄漏的点。

go tool pprof -inuse_space http://127.0.0.1:9999/debug/pprof/heap。输入top10可以看出前十占用内存情况，这里我是直接输入png导出图片来查看，以便以后比较。还有两个参数可以选择，-inuse_space顾名思义是正在使用的内存，-alloc_space是已经分配的内存，本次我是一直用-inuse_space进行分析。


go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60
执行该命令后，需等待 60 秒（可调整 seconds 的值），pprof 会进行 CPU Profiling。结束后将默认进入 pprof 的交互式命令模式，可以对分析的结果进行查看或导出。


go tool pprof -svg    http://localhost:8080/debug/pprof/heap > cpu.svg


### 安装 Graphviz 
https://graphviz.gitlab.io/_pages/Download/Download_windows.html

go tool pprof --http :9090 http://localhost:8080/debug/pprof/heap


方法一：
$ go tool pprof -http=:8080 cpu.prof
方法二：
$ go tool pprof cpu.prof 
$ (pprof) web

## 另一种可视化数据的方法是火焰图，需手动安装原生 PProf 工具：

（1） 安装 PProf

$ go get -u github.com/google/pprof
（2） 启动 PProf 可视化界面:

$ pprof -http=:8080 cpu.prof
（3） 查看 PProf 可视化界面

打开 PProf 的可视化界面时，你会明显发现比官方工具链的 PProf 精致一些，并且多了 Flame Graph（火焰图）

它就是本次的目标之一，它的最大优点是动态的。调用顺序由上到下（A -> B -> C -> D），每一块代表一个函数，越大代表占用 CPU 的时间更长。同时它也支持点击块深入进行分析！

## go tool trace trace.out
```
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
```

## 也可以写入文件

golang 的性能分析库在 runtime/pprof 里，主要提供下面几个接口

// 堆栈分析
func WriteHeapProfile(w io.Writer) error
// cpu分析
func StartCPUProfile(w io.Writer) error

方法1
```
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }
…

```

https://github.com/hatlonely/easygolang/blob/master/pprof/pprof.go



## 参考

go tool pprof 使用介绍 ：https://segmentfault.com/a/1190000016412013

Go 内存监控介绍：https://golang.org/src/runtime/mstats.go

Go 内存优化介绍：https://blog.golang.org/profiling-go-programs

高性能Go服务内存分配：https://segment.com/blog/allocation-efficiency-in-high-performance-go-services

Go stack 优化分析：https://studygolang.com/articles/10597

Go内存泄漏？不是那么简单! https://colobu.com/2019/08/28/go-memory-leak-i-dont-think-so/


## 内存不归还系统

```
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {

	go func() {
		var a []int
		for i := 0; i < 2000; i++ {
			a = make([]int, i*100000)
		}
		a = nil
		fmt.Printf("%v\n", a)
	}()

	go func() {
		for {
			m := runtime.MemStats{}
			d := debug.GCStats{}
			runtime.ReadMemStats(&m)
			debug.ReadGCStats(&d)
			fmt.Printf("%v\t%v\t%v\n", d.NumGC, m.HeapIdle/1024/1024, m.HeapInuse/1024/1024)
			time.Sleep(time.Second)
		}
	}()
	for {
		runtime.GC()
		time.Sleep(time.Second)
	}

}
```

```
package main

import (
    "fmt"
    "runtime"
    "runtime/debug"
    "time"
)

func main() {
    var a []int
    for i := 0; i < 200; i++ {
        a = make([]int, i*100000)
    }
    a = nil
    fmt.Printf("%v\n", a)
    for {
        m := runtime.MemStats{}
        d := debug.GCStats{}
        runtime.ReadMemStats(&m)
        debug.ReadGCStats(&d)
        fmt.Printf("%v\t%v\t%v\n", d.NumGC, m.HeapIdle/1024/1024, m.HeapInuse/1024/1024)
        time.Sleep(time.Second)
    }
}
```


https://www.cnblogs.com/luckcs/articles/4107647.html
```
package main
 
import (  
    "fmt" 
    "math/rand" 
    "runtime" 
    "time"
)  
 
func makeBuffer() []byte {  
    return make([]byte, rand.Intn(5000000)+5000000)  
}
 
func main() {  
    pool := make([][]byte, 20)
 
    var m runtime.MemStats  
    makes := 0  
    for {  
        b := makeBuffer()
        makes += 1
        i := rand.Intn(len(pool))
        pool[i] = b
 
        time.Sleep(time.Second)
 
        bytes := 0
 
        for i := 0; i < len(pool); i++ {
            if pool[i] != nil {
                bytes += len(pool[i])
            }
        }
 
        runtime.ReadMemStats(&m)
        fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
            m.HeapIdle, m.HeapReleased, makes)
    }
}
```

