CPU cache

https://www.toutiao.com/a6776788422771081739s

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//根据以上原理，CPU cache 在缓存数据时，并不是以单个字节为单位缓存的，而是以CPU cache line大小为单位缓存，CPU cache line 在一般的 x86 环境下为 64 字节。
//也就是说，即使从内存读取 1 个字节的数据，也会将邻近的 64 个字节一并缓存至 CPU cache 中。
//
//linux 下，可以通过getconf -a | grep CACHE命令获取 cache line 大小。
//
//这也是访问数组通常比链表快的原因之一。
// 一个uint64占 8 个字节
type Origin struct {
	a uint64
	b uint64
}

type WithPadding struct {
	a uint64
	_ [56]byte
	b uint64
	_ [56]byte
}

var num = 1000 * 1000

func OriginParallel() {
	var v Origin

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < num; i++ {
			atomic.AddUint64(&v.a, 1)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < num; i++ {
			atomic.AddUint64(&v.b, 1)
		}
		wg.Done()
	}()

	wg.Wait()
	_ = v.a + v.b
}

func WithPaddingParallel() {
	var v WithPadding

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < num; i++ {
			atomic.AddUint64(&v.a, 1)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < num; i++ {
			atomic.AddUint64(&v.b, 1)
		}
		wg.Done()
	}()

	wg.Wait()
	_ = v.a + v.b
}

func main() {
	var b time.Time

	b = time.Now()
	OriginParallel()
	fmt.Printf("OriginParallel. Cost=%+v.\n", time.Now().Sub(b))

	b = time.Now()
	WithPaddingParallel()
	fmt.Printf("WithPaddingParallel. Cost=%+v.\n", time.Now().Sub(b))
}

```

标准库中很多 _ 命名的没有意义的变量就是为了这个