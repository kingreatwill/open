> * 原文地址：[官方文档 context](https://golang.org/pkg/context/)
> * 译文地址：[https://github.com/watermelo/dailyTrans](https://github.com/watermelo/dailyTrans/blob/master/golang/context.md)
> * 译者：咔叽咔叽  
> * 译者水平有限，如有翻译或理解谬误，烦请帮忙指出

在刚刚过去的 2019 gopher china 大会上 context 概念被多次提起，看得出来 context 在 golang 的世界中是一个非常重要的知识点，所以有必要对 context 有一个基本的使用和认知。官方文档解释和示例都比较详细正规，本着学习的态度翻译一遍加深理解，该文章翻译自官方文档 [context](https://golang.org/pkg/context/)。

## 概览
context 包定义了 Context 类型，它在 API 边界和进程之间传递截止时间，取消信号和其他请求作用域的值。

服务收到请求应该要创建一个 Context，对服务的响应应该要接受一个 Context。它们之间的函数调用链必须传递 Context，也可以使用 WithCancel，WithDeadline，WithTimeout 或 WithValue 等方法创建派生 Context 替换它。取消 Context 后，也会取消从中派生的所有 Context。

WithCancel，WithDeadline 和 WithTimeout 函数接受 Context（父）并返回派生的 Context（子）和一个 CancelFunc 函数。调用 CancelFunc 函数会取消该派生的子 Context 及其孙子 Context，删除父项对子项的引用，并停止任何关联的计时器。如果没有调用 CancelFunc 会泄漏子项和孙子项，直到父项被取消或计时器触发。 go vet 工具检查是否在所有控制流路径上使用了 CancelFuncs。

使用 Contexts 的程序应遵循这些规则，以保持各个包的接口一致，并启用静态分析工具来检查上下文的传递：

不要将 Contexts 存储在结构类型中；相反，要将 Context 明确地传递给需要它的每个函数。 Context 应该是第一个参数，通常命名为 ctx：

```golang
func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}
```

即使函数允许，也不要传递 nil Context。如果你不确定要使用哪个 Context，请传递 context.TODO。

仅将上下文的值用于 API 边界和进程之间的请求作用域数据，而不是将可选参数传递给函数。

可以将相同的 Context 传递给在不同 goroutine 中运行的函数；Contexts 对于同时使用多个 goroutine 是安全的。

有关服务中使用 Contexts 的示例代码，请参考[https://blog.golang.org/context](https://blog.golang.org/context)
。

## index

## 变量
Canceled 是上下文取消时，通过 Context.Err 返回的错误。

```golang
var Canceled = errors.New("context canceled")
```

DeadlineExceeded 是在上下文超过截止时间时，通过 Context.Err 返回的错误。

```golang
var DeadlineExceeded error = deadlineExceededError{}
```

## 函数 [WithCancel](https://golang.org/src/context/context.go?s=8226:8290#L219)

```golang
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

WithCancel 返回带有新 Done channel 的父副本。返回的上下文的 Done channel 在调用返回的取消函数或父上下文的 Done channel 关闭时关闭，取决于谁先发生。

取消此上下文会释放与其相关的资源，因此代码应在此上下文中的操作完成后立即调用 cancel。

##### 示例
此示例演示了使用可取消的上下文来防止 goroutine 泄漏。在示例函数的最后，gen 启动的 goroutine 将返回，并且不会造成 goroutine 泄漏。

```golang
package main

import (
	"context"
	"fmt"
)

func main() {
	// gen 在单独的 goroutine 中生成整数并将它们发送到返回的 channel。
	// 一旦消费了生成的整数，gen 的调用者需要取消上下文，从而不会泄漏 gen 启动的内部 goroutine。
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // 返回以致不泄露 goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们消费完整数后调用取消函数

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
```

[Run in playground](https://play.golang.org/p/i2NTpVhzpIP)  

## 函数 [WithDeadline](https://golang.org/src/context/context.go?s=12173:12241#L373)

```golang
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
```

WithDeadline 返回父上下文的副本，其截止日期调整为不迟于 d。如果父级的截止日期早于 d，则 WithDeadline(parent, d)在语义上等同于 parent。返回的上下文的 Done channel 在超过截止时间后，调用返回的取消函数时或父上下文的 Done channel 关闭时关闭，三者取决于谁先发生。

取消此上下文会释放与其关联的资源，因此代码应在此上下文中的操作完成后立即调用 cancel。

##### 示例
这个例子传递一个带有任意截止时间的上下文来告诉一个阻塞的函数它应该在超时的时候丢弃它的任务。

```golang
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 即使 ctx 将要过期，在任何情况下最好也要调用它的取消函数。
	// 如果不这样做，可能会使上下文及其父级的活动时间超过必要时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
```

[Run in playground](https://play.golang.org/p/G_tpG96NxU_A)  

## 函数 [WithTimeout](https://golang.org/src/context/context.go?s=14162:14239#L440)

```golang
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

WithTimeout 返回 WithDeadline(parent, time.Now().Add(timeout))。

取消此上下文会释放与其关联的资源，因此代码应在此上下文中运行的操作完成后立即调用 cancel：

```golang
func slowOperationWithTimeout(ctx context.Context) (Result, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()  // 如果 slowOperation 在超时之前完成，则释放资源
completes before timeout elapses
	return slowOperation(ctx)
}
```

##### 示例
此示例传递具有超时的上下文，以告知一个阻塞的函数在超时后它应该丢弃它的任务。

```golang
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 传递一个带超时的上下文，以告知一个阻塞的函数在超时后它应该丢弃它的任务。
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) //打印  "context deadline exceeded"
	}

}
```

[Run in playground](https://play.golang.org/p/SfYFNZGzShR)  

## 类型 [CancelFunc](https://golang.org/src/context/context.go?s=7806:7828#L211)

CancelFunc 通知一个操作丢弃它的任务。 CancelFunc 不等待任务停止。第一次调用后，CancelFunc 的后续调用将失效。

```golang
type CancelFunc func()
```

## 类型 [Context](https://golang.org/src/context/context.go?s=2437:5895#L52)

一个 Context 可以跨 API 边界传递截止日期，取消信号和其他值。

Context 的方法可以由多个 goroutine 同时调用。

```golang
type Context interface {
        // Deadline 返回完成的任务的时间，即取消此上下文的时间。
        // 如果没有设置截止时间，Deadline 返回 ok == false。
        // 对截止日期的连续调用返回相同的结果。
        Deadline() (deadline time.Time, ok bool)

        // 当任务完成时，即此上下文被取消，Done 会返回一个关闭的channel。
        // 如果此上下文一直不被取消，Done 返回 nil。对 Done 的连续调用会返回相同的值。
        //
        // 当取消函数被调用时，WithCancel 使 Done 关闭; 
        // 在截止时间到期时，WithDeadline 使 Done 关闭;
        // 当超时的时候，WithTimeout使 Done 关闭。
        //
        // Done 可以使用 select 语句:
        //
        //  // Stream 使用 DoSomething 生成值并将它们发送到 out，
        //  // 直到 DoSomething 返回错误或 ctx.Done 关闭。
        //  func Stream(ctx context.Context, out chan<- Value) error {
        //  	for {
        //  		v, err := DoSomething(ctx)
        //  		if err != nil {
        //  			return err
        //  		}
        //  		select {
        //  		case <-ctx.Done():
        //  			return ctx.Err()
        //  		case out <- v:
        //  		}
        //  	}
        //  }
        //
        // 查看 https://blog.golang.org/pipelines 获得更多关于怎么使用 Done channel 去取消的例子
        Done() <-chan struct{}

        // 如果 Done 尚未关闭，则 Err 返回 nil。
        // 如果 Done 关闭，Err 会返回一个非nil的错误，原因：
        // 如果上下文被取消，则调用 Canceled;
        // 如果上下文的截止时间已过，则调用 DeadlineExceeded。
        // 在 Err 返回非 nil 错误后，对 Err 的连续调用返回相同的错误。
        Err() error

        // Value 返回与此上下文关联的 key 的值，如果没有值与 key 关联，则返回nil。使用相同的 key 连续调用 Value 会返回相同的结果。
        //
        // 仅将上下文的值用于API边界和进程之间的请求作用域数据，而不是将可选参数传递给函数。
        //
        // key 标识上下文中的特定值。
        // 在上下文中存储值的函数通常在全局变量中分配一个 key，然后使用该 key 作为 context.WithValue 和 Context.Value 的参数。
        // key 可以是支持比较的任何类型
        // 包应该将 key 定义为非导出类型以避免冲突。
        //
        // 定义 Context key 的包应该为使用该 key 存储的值提供类型安全的访问：
        //
        // 	// 包使用者定义一个存储在上下文中的 User 类型。
        // 	package user
        //
        // 	import "context"
        //
        // 	// User 是上下文中值的类型。
        // 	type User struct {...}
        //
        // 	// key 是此程序包中定义的 key 的非导出类型。
        // 	// 这可以防止与其他包中定义的 key 冲突。
        // 	type key int
        //
        // 	// userKey 是上下文中 user.User 值的 key。它是不可以被导出的。
        // 	// 客户端使用 user.NewContext 和 user.FromContext 而不是直接使用 key。
        // 	var userKey key
        //
        // 	// NewContext 返回一个带有值为 u 的新的上下文。
        // 	func NewContext(ctx context.Context, u *User) context.Context {
        // 		return context.WithValue(ctx, userKey, u)
        // 	}
        //
        // 	// FromContext 返回存储在 ctx 中的 User 值（如果有的话）。
        // 	func FromContext(ctx context.Context) (*User, bool) {
        // 		u, ok := ctx.Value(userKey).(*User)
        // 		return u, ok
        // 	}
        Value(key interface{}) interface{}
}
```

### 函数 [Background](https://golang.org/src/context/context.go?s=7302:7327#L196)

```golang
func Background() Context
```

Background 返回一个非 nil 的空 Context。它永远不会被取消，没有值，也没有截止时间。它通常由 main 函数初始化和测试使用，并作为请求的顶级 Context。

### 函数 [TODO](https://golang.org/src/context/context.go?s=7590:7609#L204)

```golang
func TODO() Context
```

TODO 返回一个非 nil 的空 Context。当不清楚使用哪个 Context 或者它还不可用时（因为周围的函数尚未扩展为接受 Context 参数），代码应该使用 context.TODO。

### 函数 [WithValue](https://golang.org/src/context/context.go?s=14954:15014#L457)

```golang
func WithValue(parent Context, key, val interface{}) Context
```
WithValue 返回父级的副本，其中与 key 关联的值为 val。

仅将上下文的值用于 API 边界和进程之间的请求作用域数据，而不是将可选参数传递给函数。

提供的 key 必须是可比较的，不应该是字符串类型或任何其他内置类型，以避免使用上下文的包之间产生冲突。 WithValue 的使用者应该为 keys 定义他们自己的自定义类型。为了避免在分配 interface{}时指定，上下文的 keys 通常具有具体类型 struct {}。或者，导出的上下文的 key 变量的静态类型应该是指针或接口。

##### 示例
此示例展示如何将值传递给上下文以及如何检索它（如果存在）。

```golang
package main

import (
	"context"
	"fmt"
)

func main() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

}
```
[Run in playground](https://play.golang.org/p/QiFNd_jJkVI)  








