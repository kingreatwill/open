> * 原文地址：[https://dave.cheney.net/2016/03/19/should-methods-be-declared-on-t-or-t](https://dave.cheney.net/2016/03/19/should-methods-be-declared-on-t-or-t)
> * 原文作者：[Dave Cheney](https://dave.cheney.net/)
> * 译文地址：[https://github.com/watermelo/dailyTrans](https://github.com/watermelo/dailyTrans/blob/master/golang/method_use_T_or_pointer_T.md)
> * 译者：咔叽咔叽  
> * 译者水平有限，如有翻译或理解谬误，烦请帮忙指出

这篇文章是我几天前在 Twitter 上提出的[建议](https://twitter.com/davecheney/status/710604764640256000)的延续。

在 Go 中，对于任何类型 T 都存在类型 *T，表示获取 T 类型（T 表示你声明的类型）变量的地址。例如：

```go
type T struct { a int; b bool }
var t T    // t's type is T
var p = &t // p's type is *T
```

这两种类型，T 和 *T 是不同的，*T 不能替代 T（此规则是递归的，**T 会返回 *T 地址指向的值）。

你可以在任何类型上声明方法;也就是说，你在 package 中声明了一个类型。因此，你可以在这个类型上声明一个方法，他的接收者可以使用 T 或者 *T。或者是说声明接收者的类型为 T 是为了获取接收者值的副本，声明接收者的类型为 *T 是为了获取指向接收者值的指针（Go 中的方法只是函数的语法糖，它将接收者作为第一个形式参数传递）。那么问题就变成了，我们应该选择哪种方式？（如果该方法不改变它的接收者，它是否需要方法这种形式？）

显然，如果你的方法改变了接收者，那应该声明 * T。但是，如果该方法不改变其接收者，是否可以将其声明为 T 呢？

事实证明，这样做的场景非常有限。例如，众所周知，你不应该复制 sync.Mutex 值，因为它会使互斥锁失效。由于互斥锁控制对数据的访问，它们经常被包含在结构中：

```go
package counter

type Val struct {
        mu  sync.Mutex
        val int
}

func (v *Val) Get() int {
        v.mu.Lock()
        defer v.mu.Unlock()
        return v.val
}

func (v *Val) Add(n int) {
        v.mu.Lock()
        defer v.mu.Unlock()
        v.val += n
}
```

大多数 Go 程序员都知道我们应该使用指针接收者 *Val ，并在其上声明 Get 或 Add 方法。但是，任何嵌入 Val 以利用其零值的类型，也必须把方法的接收者设为指针类型，否则它可能复制其嵌入类型的值。

```go
type Stats struct {
        a, b, c counter.Val
}

func (s Stats) Sum() int {
        return s.a.Get() + s.b.Get() + s.c.Get() // whoops
}
```

对于切片类型，可能也会发生类似的陷阱，当然也有可能发生[意外的数据竞争](https://dave.cheney.net/2015/11/18/wednesday-pop-quiz-spot-the-race)。

简而言之，我认为你更应该在 *T 上声明方法，除非你有充分的理由不这样做。

## 相关文章
1. [What is the zero value, and why is it useful?](https://dave.cheney.net/2013/01/19/what-is-the-zero-value-and-why-is-it-useful)
2. [Ice cream makers and data races](https://dave.cheney.net/2014/06/27/ice-cream-makers-and-data-races)
3. [Slices from the ground up](https://dave.cheney.net/2018/07/12/slices-from-the-ground-up)
4. [The empty struct](https://dave.cheney.net/2014/03/25/the-empty-struct)