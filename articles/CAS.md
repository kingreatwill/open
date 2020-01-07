# Compare And Swap

借用CPU提供的原子性指令来实现,CAS操作修改共享变量时候不需要对共享变量加锁，而是通过类似乐观锁的方式进行检查，本质还是不断的占用CPU 资源换取加锁带来的开销（比如上下文切换开销）

CAS操作具有原子性，在解决多线程操作共享变量安全上可以有效的减少使用锁所带来的开销，但是这是使用cpu资源做交换的。

## CAS使用场景#
- 使用一个变量统计网站的访问量；
- Atomic类操作；
- 数据库乐观锁更新。

## 各种语言相关函数
```
go atomic.CompareAndSwapInt32

https://docs.microsoft.com/zh-cn/dotnet/api/system.threading.interlocked?view=netcore-3.1

csharp InterlockedCompareExchangeObject
System.Threading.Interlocked 
CountdownEvent  计数

java compareAndSwapObject(AtomicReference)
Unsafe类
```
```go
package main

import (
	"fmt"

	"sync"

	"sync/atomic"
)

var (
	counter int32          //计数器
	wg      sync.WaitGroup //信号量
)

func main() {
	threadNum := 5000
	//1. 五个信号量
	wg.Add(threadNum)
	//2.开启5个线程
	for i := 0; i < threadNum; i++ {
		go incCounter(i)
	}
	//3.等待子线程结束
	wg.Wait()
	fmt.Println(counter)
}

func incCounter(index int) {
	defer wg.Done()
	spinNum := 0
	for {
		//2.1原子操作
		old := counter
		ok := atomic.CompareAndSwapInt32(&counter, old, old+1)
		if ok {
			break
		} else {
			spinNum++
		}
	}
	if spinNum > 0 {
		fmt.Printf("thread,%d,spinnum,%d\n", index, spinNum)
	}

}
```
atomic.CompareAndSwapInt32具有三个参数，第一个是变量的地址，第二个是变量当前值，第三个是要修改变量为多少，该函数如果发现传递的old值等于当前变量的值，则使用第三个变量替换变量的值并返回true，否则返回false。




## 参考

[并发编程的基石——CAS机制](https://www.cnblogs.com/54chensongxia/p/12160085.html)

[无锁机制----比较交换CAS Compare And Swap](https://blog.csdn.net/yanluandai1985/article/details/82686486)