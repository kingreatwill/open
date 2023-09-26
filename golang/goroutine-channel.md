
https://mp.weixin.qq.com/s/qDFgMzF1onowF33_pczP-Q
Go 特意模糊堆和栈之后，你对 Goroutine 栈了解多少？
https://www.toutiao.com/a6764649154305065475/
协程与通道
https://www.toutiao.com/a6769878961003430414
深入理解Golang之channel
https://juejin.im/post/5decff136fb9a016544bce67

goroutine究竟占用多少内存？
https://juejin.im/post/5d9ff459f265da5b8a5160f5

Golang 的 goroutine 是如何实现的？
https://www.zhihu.com/question/20862617/answer/921061289

goroutine背后的系统知识
http://www.sizeofvoid.net/goroutine-under-the-hood/

深度解密Go语言之channel
https://mp.weixin.qq.com/s/90Evbi5F5sA1IM5Ya5Tp8w

Go 中 goroutine 是如何协作与抢占
https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651439330&idx=3&sn=9ebccea18ceeb9e13474877c822c5ec0

[6.7 协作与抢占](https://changkun.de/golang/zh-cn/part2runtime/ch06sched/preemption/)

[Go 语言原本 基于1.14.0](https://github.com/changkun/go-under-the-hood)
https://changkun.de/golang/


[What Are JVM Fibers and Threads?](https://www.jrebel.com/blog/jvm-fibers-and-threads)

coroutine/协程,fiber/纤程,thread/线程,process/进程

goroutine是golang中的coroutine

goroutine是强制加上scheduler的纤程

纤程和协程的差异点在于，fiber/goroutine多了一个调度器，而这个调度器会把callback后的代码调度到其他线程去执行

如果还是原来线程在执行，那么就是coroutine
如果存在有被其他线程所执行的可能，那就是fiber



分类 |	OS Schedule (multicore support)|	User Schedule (Language/Library)
---|---|---
Preemptive|	thread	| goroutine?
Non-preemptive (cooperative)|	fiber	|coroutine


最主要的区别是调度方式，一个是协作式，一个是抢占式，而且本质上协程做不到并行，只能是并发，还有就是协程并不需要同步原语：锁、信号量等这些，因为不存在临界区。


**Fiber VS Coroutine VS Green Thread**
**纤程(Fiber)**是一个轻量级的线程,它使用协作完成多任务，而不是使用抢占使用多任务。正在运行的Fiber必须显示让出以此允许其他Fiber运行，这使得Fiber的实现
比内核线程和用户线程简单很多.Java有一个Fiber库  

**协程(Coroutine)**是一个组件，它被概括成多个子程序，允许多个进入点来暂停程序和恢复程序。但是和子程序又不一样，当前协程能够通过调用其他协程来退出，这就可能在稍后
返回到在原始协程中调用的它们的位置。

**Green Thread** 是由VM调用的线程，而不是由下层操作系统调用的线程.Green Thread模拟多线程环境而不用考虑本地OS的兼容能力,它们在用户空间中管理，而不是内核空间中
管理，它们可以在没有本地线程支持的环境中工作。

## 线程调度
定义： 系统为线程分配处理器使用权的过程。
分类： 主要调度方式分两种，分别是协同式线程调度和抢占式线程调度。

### 同步式线程调度 VS 抢占式线程调度

.|协同式线程调度 |	抢占式线程调度
---|---|---
控制权|	线程本身（线程执行完后，主动通知系统切换）|	系统决定
优点|	1.切换操作线程已知，控制简单 2.不存在线程同步问题|	线程执行时间可控，不会因为一个线程耽误整个进程
缺点|	执行时间不可控，一个线程可能耽误整个进程	|1.切换控制复杂 2.存在线程同步问题

[Golang 并发问题（五）goroutine 的调度及抢占](https://studygolang.com/articles/21227)
[Golang 抢占调度流程分析](https://studygolang.com/articles/21421)

Go1.14.1 发布，修正抢占调度等重要 bug
https://github.com/golang/go/issues?q=milestone%3AGo1.14.1+label%3ACherryPickApproved
https://github.com/golang/go/issues/37833

## [goroutine 不是协程，项目描述不准确](https://github.com/panjf2000/ants/issues/60)

首先，goroutine 不是我们通常说的协作式调度，绝大部分的 go 程序都是没有 runtime.Gosched 的。如果认为 go scheduler 内部的抢占调度目前不完善，那 linux 内核在 2.几 某个版本之前也没有完善的抢占调度，难道 linux 线程也是协程吗。

另外，预期 Go 1.14 就有完全的抢占调度了，更跟协程不搭边了。

[Is a Go goroutine a coroutine?](https://stackoverflow.com/questions/18058164/is-a-go-goroutine-a-coroutine)
goroutine 是coroutine吗？ 答：不完全是


## [Why called goroutines](https://golang.org/doc/effective_go.html?h=coroutine)
```
They're called goroutines because the existing terms—threads, coroutines, processes, and so on—convey inaccurate connotations. A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.

Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run. Their design hides many of the complexities of thread creation and management.

Prefix a function or method call with the go keyword to run the call in a new goroutine. When the call completes, the goroutine exits, silently. (The effect is similar to the Unix shell's & notation for running a command in the background.)

go list.Sort()  // run list.Sort concurrently; don't wait for it.

A function literal can be handy in a goroutine invocation.

func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
In Go, function literals are closures: the implementation makes sure the variables referred to by the function survive as long as they are active.

These examples aren't too practical because the functions have no way of signaling completion. For that, we need channels.
```


## Why goroutines instead of threads?
[Why goroutines instead of threads](https://golang.org/doc/faq#goroutines)
```
Goroutines are part of making concurrency easy to use. The idea, which has been around for a while, is to multiplex independently executing functions—coroutines—onto a set of threads. 
When a coroutine blocks, such as by calling a blocking system call, the run-time automatically moves other coroutines on the same operating system thread to a different, runnable thread so they won't be blocked. 
The programmer sees none of this, which is the point. 
The result, which we call goroutines, can be very cheap: they have little overhead beyond the memory for the stack, which is just a few kilobytes.

To make the stacks small, Go's run-time uses resizable, bounded stacks. 
A newly minted goroutine is given a few kilobytes, which is almost always enough. 
When it isn't, the run-time grows (and shrinks) the memory for storing the stack automatically, allowing many goroutines to live in a modest amount of memory. 
The CPU overhead averages about three cheap instructions per function call. It is practical to create hundreds of thousands of goroutines in the same address space. 
If goroutines were just threads, system resources would run out at a much smaller number.
```

## Concurrency: goroutines and channels
https://en.wanweibaike.com/wiki-Goroutine#Concurrency:_goroutines_and_channels
https://en.wikipedia.org/wiki/Go_(programming_language)#Concurrency
```
The Go language has built-in facilities, as well as library support, for writing concurrent programs. Concurrency refers not only to CPU parallelism, but also to asynchrony: letting slow operations like a database or network read run while the program does other work, as is common in event-based servers.[79]

The primary concurrency construct is the goroutine, a type of light-weight process. A function call prefixed with the go keyword starts a function in a new goroutine. The language specification does not specify how goroutines should be implemented, but current implementations multiplex a Go process's goroutines onto a smaller set of operating-system threads, similar to the scheduling performed in Erlang.[80]:10

While a standard library package featuring most of the classical concurrency control structures (mutex locks, etc.) is available,[80]:151–152 idiomatic concurrent programs instead prefer channels, which provide send messages between goroutines.[81] Optional buffers store messages in FIFO order[64]:43 and allow sending goroutines to proceed before their messages are received.[citation needed]

Channels are typed, so that a channel of type chan T can only be used to transfer messages of type T. Special syntax is used to operate on them; <-ch is an expression that causes the executing goroutine to block until a value comes in over the channel ch, while ch <- x sends the value x (possibly blocking until another goroutine receives the value). The built-in switch-like select statement can be used to implement non-blocking communication on multiple channels; see below for an example. Go has a memory model describing how goroutines must use channels or other operations to safely share data.[82]

The existence of channels sets Go apart from actor model-style concurrent languages like Erlang, where messages are addressed directly to actors (corresponding to goroutines). The actor style can be simulated in Go by maintaining a one-to-one correspondence between goroutines and channels, but the language allows multiple goroutines to share a channel or a single goroutine to send and receive on multiple channels.[80]:147

From these tools one can build concurrent constructs like worker pools, pipelines (in which, say, a file is decompressed and parsed as it downloads), background calls with timeout, "fan-out" parallel calls to a set of services, and others.[83] Channels have also found uses further from the usual notion of interprocess communication, like serving as a concurrency-safe list of recycled buffers,[84] implementing coroutines (which helped inspire the name goroutine),[85] and implementing iterators.[86]

Concurrency-related structural conventions of Go (channels and alternative channel inputs) are derived from Tony Hoare's communicating sequential processes model. Unlike previous concurrent programming languages such as Occam or Limbo (a language on which Go co-designer Rob Pike worked),[87] Go does not provide any built-in notion of safe or verifiable concurrency.[88] While the communicating-processes model is favored in Go, it is not the only one: all goroutines in a program share a single address space. This means that mutable objects and pointers can be shared between goroutines; see § Lack of race condition safety, below.[citation needed]
```

## Coroutine comparison with threads
https://en.wikipedia.org/wiki/Coroutine#Comparison_with_threads
```

Coroutines are very similar to threads. However, coroutines are cooperatively multitasked, whereas threads are typically preemptively multitasked. 
This means that coroutines provide concurrency but not parallelism. The advantages of coroutines over threads are that they may be used in a hard-realtime context (switching between coroutines need not involve any system calls or any blocking calls whatsoever), there is no need for synchronisation primitives such as mutexes, semaphores, etc. 
in order to guard critical sections, and there is no need for support from the operating system.

It is possible to implement coroutines using preemptively-scheduled threads, in a way that will be transparent to the calling code, but some of the advantages (particularly the suitability for hard-realtime operation and relative cheapness of switching between them) will be lost.
```


[wiki-Goroutine](https://en.wanweibaike.com/wiki-Goroutine)

## What are Goroutines?
https://golangbot.com/goroutines/
```
What are Goroutines?
Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can be thought of as light weight threads. The cost of creating a Goroutine is tiny when compared to a thread. Hence its common for Go applications to have thousands of Goroutines running concurrently.

Advantages of Goroutines over threads
Goroutines are extremely cheap when compared to threads. They are only a few kb in stack size and the stack can grow and shrink according to needs of the application whereas in the case of threads the stack size has to be specified and is fixed.
The Goroutines are multiplexed to fewer number of OS threads. There might be only one thread in a program with thousands of Goroutines. If any Goroutine in that thread blocks say waiting for user input, then another OS thread is created and the remaining Goroutines are moved to the new OS thread. All these are taken care by the runtime and we as programmers are abstracted from these intricate details and are given a clean API to work with concurrency.
Goroutines communicate using channels. Channels by design prevent race conditions from happening when accessing shared memory using Goroutines. Channels can be thought of as a pipe using which Goroutines communicate. We will discuss channels in detail in the next tutorial.
```


[Go语言常见坑](https://studygolang.com/articles/16949?fr=sidebar)

安装其它版本go
https://pkg.go.dev/golang.org/dl
1. go get golang.org/dl/go1.13.6
2. go1.13.6.exe download
3. go1.13.6 version

## 独占CPU导致其它Goroutine饿死

Goroutine是协作式调度, Goroutine本身不会主动放弃CPU:
```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
    }()

    for {} // 占用CPU
}

// go1.14 会输出
// go1.13 不会输出
```
解决的方法是在for循环加入runtime.Gosched()调度函数:
```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
    }()

    for {
        runtime.Gosched() //调度
    }
}
```
或者是通过阻塞的方式避免CPU占用:
```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
    }()

    select{}
}
```