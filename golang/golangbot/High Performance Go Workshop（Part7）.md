<!-- toc -->
[TOC]
# High Performance Go Workshop（Part7）

## 7\. Tips 和 trips

  

这是一些优化 Go 代码的技巧。

  

### 7.1. Goroutines

  

goroutines 是 Go 的关键特性，它非常适合现代硬件。

  

goroutine 很容易使用，而且开销很低，你可以认为它们几乎是 0 开销的。

  

Go 运行时是为以成千上万个 goroutine 为标准的程序编写的，成千上万个没什么的。

  

但是，每个 goroutine 确实消耗了 goroutine 的堆栈的最小内存量，该堆栈当前至少为 2k。

  

2048*1000000 个 goroutines ==2GB 内存，它们还没有做任何事情。

  

#### 7.1.1. 知道什么时候停止 goroutine

  

Goroutines 的启动成本和运行成本都很低，但就内存占用而言，它们的成本是有限的；你不能创建无限数量的 goroutine。

  

每次在程序中使用 `go` 关键字启动 goroutine 时，都必须知道该 goroutine 将**如何**以及**何时**退出。

  

在你的设计中，某些 goroutine 可能会一直运行到程序退出。这些 goroutine 非常罕见，不会成为规则的例外。

  

如果你不知道答案，那就是潜在的内存泄漏，因为 Goroutines 会将其堆栈的内存固定在堆上，以及可以从堆栈访问的任何堆分配变量。

  

> 永远不要在不知道如何停止的情况下开启一个 goroutine 。轻灵划译

  

#### 7.1.2. 延伸阅读

  

*   [Concurrency Made Easy](https://www.youtube.com/watch?v=yKQOunhhf4A&index=16&list=PLq2Nv-Sh8EbZEjZdPLaQt1qh_ohZFMDj8) (video)

  

*   [Concurrency Made Easy](https://dave.cheney.net/paste/concurrency-made-easy.pdf) (slides)

  

*   [Never start a goroutine without knowning when it will stop](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_never_start_a_goroutine_without_knowning_when_it_will_stop) (Practical Go, QCon Shanghai 2018)  
    

  

  

### 7.2. Go 对某些请求使用高效的网络轮询

  

Go 运行时使用有效的操作系统轮询机制（kqueue、epoll、windows IOCP 等）处理网络 io。许多等待的 goroutine 将由单个操作系统线程提供服务。

  

但是，对于本地文件 IO，Go 不实现任何 IO 轮询。`*os.File` 上的每个操作在进行过程中消耗一个操作系统线程。

  

大量使用本地文件 IO 会导致程序产生成百上千个线程；可能超出操作系统的允许范围。

  

你的磁盘子系统可能无法处理数百或数千个并发 IO 请求。

  

> 要限制并发阻塞 IO 的数量，可以使用 goroutines 池或缓冲通道作为信号量。
> 
>   
> 
``` 
var semaphore = make(chan struct{}, 10)   
func processRequest(work *Work) {  
    semaphore <- struct{}{} // acquire semaphore 
    // process request 
    <-semaphore // release semaphore 
}
```

  

  

### 7.3. 注意你的应用程序中的 IO 乘数

  

如果你正在编写服务器程序，那么它的主要工作是多路复用通过网络连接的客户端和存储在应用程序中的数据。

  

大多数服务器程序接受一个请求，做一些处理，然后返回一个结果。这听起来很简单，但根据结果，它可以让客户机在您的服务器上消耗大量（可能是无限制的）资源。以下是一些需要注意的事项：

  

*   每个传入请求的 IO 请求数量；单个客户端请求生成多少 IO 事件？它的平均值可能是 1，如果从缓存中服务了许多请求，则可能小于 1。

  

*   服务查询所需的读取量；是固定的、N+1 还是线性的（读取整个表以生成最后一页结果）。

  

如果相对而言，如果内存是非常慢的，那么 IO 就是龟速了，所以你应该不惜一切代价避免这样做。最重要的是避免在请求的上下文中执行 IO，不要让用户等待你的磁盘子系统写入磁盘，甚至读取磁盘。

  

  

### 7.4. 使用流 IO 接口

  

尽可能避免将数据读入一个 `[]byte` 并将其四处传递。

  

根据请求，你可能最终读取兆字节（或者更大！）把数据存入内存。这给 GC 带来了巨大的压力，这将增加应用程序的平均延迟。

  

而是使用 `io.Reader` 和 `io.Writer` 来构造处理管道，以限制每个请求使用的内存量。

  

为了提高效率，如果使用大量 `io.Copy`，请考虑实现 `io.ReaderFrom` / `io.WriterTo`。这些接口效率更高，避免将内存复制到临时缓冲区中。

  

  

### 7.5. 超时,超时,超时

  

在不知道要花多长时间的情况下，不要启动 IO 操作。

  

你需要对使用 `SetDeadline`、`SetReadDeadline` 和 `SetWriteDeadline` 发出的每个网络请求设置超时。

  

  

### 7.6. Defer 是昂贵的，是吗？

  

`defer` 代价高昂，因为它必须为 defer 参数记录一个闭包。

  
```
defer mu.Unlock()
```
  

等价于

  
```
defer func() {  mu.Unlock() }()
```
  

如果工作量很小，那么 `defer` 是开销很大的，像常见的例子，比如在 struct 变量或 map 查找释放互斥锁。在这种情况下，你可以选择避免 `defer` 。

  

这是一个牺牲可读性和维护性来赢得性能的例子。

  

一直要重新考虑这些决定。

  

### 7.7. 避免 Finalisers

  

Finalisation 是一种将行为附加到即将被垃圾收集的对象上的技术。

  

因此，finalisation 是不确定的。

  

若要运行 finaliser ，任何对象都不能访问该对象。如果你意外地在字典中保留了对该对象的引用，则该对象不会被最终确定。

  

Finalisers 作为 gc 周期的一部分运行，这意味着他们何时运行是不可预测的，并使他们与减少 gc 操作的目标相悖。

  

如果你有一个很大的堆，并且调整了你的应用程序产生了最低程度的垃圾，那么 finaliser 可能不会持续很长时间。

  

### 7.8. 最小化 cgo

  

cgo 允许 Go 程序调用 C 库。

  

C 代码和 Go 代码在两个不同的世界中，cgo 穿越了它们之间的边界。

  

这种转换不是免费的，并且取决于它在代码中的位置，成本可能会很大。

  

cgo 调用类似于阻塞 IO，它们在操作期间消耗线程。

  

不要在多重循环中调用 C 代码。

  

#### 7.8.1. 实际上，可以避免使用 cgo

  

cgo 有产生很大的开销。

  

为了获得最佳性能，我建议在应用程序中避免使用 cgo。

  

*   如果c代码需要很长时间，那么 cgo 开销就不那么重要了。

  

*   如果你使用 cgo 调用一个非常短的 C 函数（其中开销最明显），请在Go 中重写该代码。

  

*   如果你在多重循环中使用的是一大块开销很大的 C 代码，那么为什么要使用 Go 呢？

  

##### 延伸阅读

  

*   [cgo is not Go](http://dave.cheney.net/2016/01/18/cgo-is-not-go)

  

### 7.9. 始终使用最新发布的 Go 版本

  

旧版本的 Go 永远不会变得更好。他们永远不会得到 bug 修复或优化。

  

*   不应使用 GO 1.4。

  

*   Go 1.5 和 1.6 的编译器速度较慢，但它生成的代码更快，GC 也更快。

  

*   Go 1.7 的编译速度比 1.6 提高了 30%，链接速度提高了 2 倍（比 Go 以前的任何版本都好）

  

*   Go 1.8 在编译速度上会有较小的提高（此时此刻），但在非英特尔体系结构的代码质量上会有显著的提高。

  

*   Go 1.9-1.12 继续提高生成代码的性能，修复错误，改进内联和调试。

  

> 旧版本的 Go 没有更新。**不要使用它们**。使用最新的版本，你会得到最好的性能。（跟 Node 一样啊，V8一发布新版本，node 升级下程序就跑的更快了）

  

#### 7.9.1. 延伸阅读

  

*   [Go 1.7 toolchain improvements](http://dave.cheney.net/2016/04/02/go-1-7-toolchain-improvements)

  

*   [Go 1.8 performance improvements](http://dave.cheney.net/2016/09/18/go-1-8-performance-improvements-one-month-in)

  

  

## 最后的问题和结论

  

> Readable means reliable — Rob Pike

  

从最简单的代码开始。

  

测量。分析代码以确定瓶颈，不要猜测。

  

如果运行的稳定，就停下来。你不需要优化所有内容，只需要优化代码中性能最关键的部分。

  

随着应用程序的增长或流量模式的发展，性能热点将发生变化。

  

不要把不是性能关键的复杂代码放在一边，如果瓶颈转移到其他地方，就用更简单的方式重写它们。

  

总是尽可能地编写最简单的代码，编译器会针对普通代码进行优化。

  

更短的代码是更快的代码；Go 不是 C++，不要指望编译器解开复杂的抽象（一时分不清是在黑还是夸。。）。

  

代码越短，代码就越小；这对CPU的缓存很重要。

注意分配，尽量避免不必要的分配。

  

> I can make things very fast if they don’t have to be correct. — Russ Cox

  

性能和可靠性同样重要。

  

我认为一个非常快的服务器如果频繁地 panics，死锁 或 OOMs，那么这样是毫无价值的。

  

不要为了可靠性而牺牲性能。
