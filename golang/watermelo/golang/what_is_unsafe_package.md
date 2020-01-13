> * 原文地址：[https://peter.bourgon.org/blog/2019/09/11/programming-with-errors.html](https://peter.bourgon.org/blog/2019/09/11/programming-with-errors.html)
> * 原文作者：[Peter](https://peter.bourgon.org/about/)
> * 译文地址：[https://github.com/watermelo/dailyTrans](https://github.com/watermelo/dailyTrans/blob/master/golang/programming_with_errors.md)
> * 译者：咔叽咔叽  
> * 译者水平有限，如有翻译或理解谬误，烦请帮忙指出

Go 1.13 引入了一个增强的[package errors](https://golang.org/pkg/errors/)，大致标准化了错误处理。就个人而言，我觉得它的 API 令人有点困惑。本文提供一些如何更有效使用它的参考。

## 创建 errors
sentinel errors（译者注：表示在此错误中断，程序不会继续往下处理）和以前一样。将它们命名为 ErrXxx，使用 errors.New 来创建它们。

```go
var ErrFoo = errors.New("foo error")
```

错误类型基本上也和以前一样。将它们命名为 XxxError，并确保它们具有 Error 方法，以满足 error 接口。

```go
type BarError struct {
    Reason string
}

func (e BarError) Error() string {
    return fmt.Sprintf("bar error: %s", e.Reason)
}
```

如果你的错误类型包装了另一个错误，就需要提供 Unwrap 方法。

```go
type BazError struct {
    Reason string
    Inner  error
}

func (e BazError) Error() string {
    if e.Inner != nil {
        return fmt.Sprintf("baz error: %s: %v", e.Reason, e.Inner)
    }
    return fmt.Sprintf("baz error: %s", e.Reason)
}

func (e BazError) Unwrap() error {
    return e.Inner
}
```

## 包装和返回错误
默认情况下，当你在函数中遇到错误并需要将其返回给调用者时，可以通过 [fmt.Errorf](https://golang.org/pkg/fmt/#Errorf) 的 `％w` 格式，使用相关上下文包装错误。

```go
func process(j Job) error {
    result, err := preprocess(j)
    if err != nil {
         return fmt.Errorf("error preprocessing job: %w", err)
    }
```

此过程称为错误注解。需要避免返回未注解的错误，因为这可能使调用者不知道出错的地方在哪里。

另外，考虑通过自定义错误类型（如上面的 BazError）包装错误，以获得更复杂的用例。

```go
p := getPriority()
widget, err := manufacture(p, result)
if err != nil {
    return ManufacturingError{Priority: p, Error: err}
}
```

## 错误检查
大多数情况下，当你收到错误时，不需要关心细节。如果你的代码执行失败了，你需要报出错误（例如记录它）并继续;或者，如果无法继续，可以使用上下文来注解错误，并将其返回给调用者。

如果你想知道收到的是哪个错误，可以用 [errors.Is](https://golang.org/pkg/errors/#Is) 检查 sentinel errors，也可以用 [errors.As](https://golang.org/pkg/errors/#As)来检查错误值。

```go
err := f()
if errors.Is(err, ErrFoo) {
    // you know you got an ErrFoo
    // respond appropriately
}

var bar BarError
if errors.As(err, &bar) {
    // you know you got a BarError
    // bar's fields are populated
    // respond appropriately
}
```

errors.Is 和 errors.As 会尝试以递归的方式解包错误来找到匹配项。[此代码](https://play.golang.org/p/GorSR6HTWzf)**演示了基本的错误包装和检查技术**（译者注：需要科学上网，把这段代码贴到文章末尾了）。查看 `func a()` 中检查的顺序，然后尝试更改 `func c()` 返回的错误，以获得关于运行的流程。

正如[文档](https://golang.org/pkg/errors/)所述，更偏向使用 errors.Is 来检查普通等式，例如 `if err == ErrFoo` ;更偏向使用 errors.As 来断言普通类型，例如 `if e，ok := err.(MyError)`，因为普通版本不执行 unwrap 操作。如果你明确不希望调用者 unwrap 错误，可以为 `fmt.Errorf` 提供不同的格式化动词，例如 `％v` ;或者不要在错误类型上提供 `Unwrap` 方法。但这些例不是常见的。

## 示例
```go
package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	i, err := a()
	log.Printf("i=%d err=%v", i, err)
}

//
//
//

func a() (int, error) {
	i, err := b()
	if errors.Is(err, ErrFoo) {
		return 0, fmt.Errorf("tragedy: %w", err)
	}

	var bar BarError
	if errors.As(err, &bar) {
		return 0, fmt.Errorf("comedy: %w", err)
	}

	var baz BazError
	if errors.As(err, &baz) {
		return 0, fmt.Errorf("farce: %w", err)
	}

	return i, nil
}

func b() (int, error) {
	if err := c(); err != nil {
		return 0, fmt.Errorf("error executing c: %w", err)
	}
	return 1, nil
}

func c() error {
	// return ErrFoo
	// return BarError{Reason: "😫"}
	// return BazError{Reason: "☹️"}
	return BazError{Reason: "😟", Inner: ErrFoo}
}

//
//
//

var ErrFoo = errors.New("foo error")

//
//
//

type BarError struct {
	Reason string
}

func (e BarError) Error() string {
	return fmt.Sprintf("bar error: %s", e.Reason)
}

//
//
//

type BazError struct {
	Reason string
	Inner  error
}

func (e BazError) Unwrap() error {
	fmt.Println("fuck")
	return e.Inner
}

func (e BazError) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("baz error: %s: %v", e.Reason, e.Inner)
	}
	return fmt.Sprintf("baz error: %s", e.Reason)
}
```
