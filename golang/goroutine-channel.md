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


coroutine/协程,fiber/纤程,thread/线程,process/进程

goroutine是golang中的coroutine

goroutine是强制加上scheduler的纤程

纤程和协程的差异点在于，fiber/goroutine多了一个调度器，而这个调度器会把callback后的代码调度到其他线程去执行

如果还是原来线程在执行，那么就是coroutine
如果存在有被其他线程所执行的可能，那就是fiber

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

[goroutine 不是协程，项目描述不准确](https://github.com/panjf2000/ants/issues/60)
首先，goroutine 不是我们通常说的协作式调度，绝大部分的 go 程序都是没有 runtime.Gosched 的。如果认为 go scheduler 内部的抢占调度目前不完善，那 linux 内核在 2.几 某个版本之前也没有完善的抢占调度，难道 linux 线程也是协程吗。

另外，预期 Go 1.14 就有完全的抢占调度了，更跟协程不搭边了。

[Is a Go goroutine a coroutine?](https://stackoverflow.com/questions/18058164/is-a-go-goroutine-a-coroutine)
goroutine 是coroutine吗？ 答：不完全是
 [ Why goroutines instead of threads](https://golang.org/doc/faq#goroutines)



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