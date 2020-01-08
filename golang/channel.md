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