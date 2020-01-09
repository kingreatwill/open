<!-- toc -->
[TOC]
# golang之defer简介

每次defer语句执行的时候，会把函数“压栈”，函数参数会被拷贝下来；当外层函数（非代码块，如一个for循环）退出时，defer函数按照定义的逆序执行；如果defer执行的函数为nil, 那么会在最终调用函数的产生panic.

defer语句并不会马上执行，而是会进入一个栈，函数return前，会按先后出的顺序执行。也说是说最先被定义的defer语句最后执行。先进后出的原因是后面定义的函数可能会依赖前面的资源，自然要先执行；否则，如果前面先执行，那后面函数的依赖就没有了。

在defer函数定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。作为函数参数，则在defer定义时就把值传递给defer，并被cache起来；作为闭包引用的话，则会在defer函数真正调用时根据整个上下文确定当前的值。

defer后面的语句在执行的时候，函数调用的参数会被保存起来，也就是复制了一份。真正执行的时候，实际上用到的是这个复制的变量，因此如果此变量是一个“值”，那么就和定义的时候是一致的。如果此变量是一个“引用”，那么就可能和定义的时候不一致。

```go
func main() {
	fmt.Println(increaseA()) // 0
    fmt.Println(increaseB()) // 1
    fmt.Println(f1())        // 10
	fmt.Println(f2())        // 5
    fmt.Println(f3())        // 0
    fmt.Println(f4())        // 2
    fmt.Println(f5())        // 2
}

func increaseA() int {
	var i int
	defer func() { i++ }()
	return i
}

func increaseB() (r int) {
	defer func() { r++ }()
	return r
}


func f1() (r int) {
	r = 5
	defer func() { r = r + 5 }()
	return r
}

func f2() (r int) {
	t := 5
	defer func() { t = t + 5 }()
	return t
}

func f3() (r int) {
	defer func(r int) { r = r + 5 }(r)
	return r
}

func f4() (r int) {
	i := 1
	defer func() {
		r++
	}()
	return i
}

func f5() (r int) {
	defer func() {
		r++
	}()
	return 1
}

```
return xxx 这条语句经过编译之后，变成了三条指令：
1. 返回值 = xxx
2. 调用defer函数
3. 空的return

f2 拆解后：
```go
func f2() (r int) {
     t := 5
     
     // 1. 赋值指令
     r = t
     
     // 2. defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
     func() {        
         t = t + 5
     }
     
     // 3. 空的return指令
     return
}
```

f4 拆解后：
```go
func f4() (r int) {
    i := 1
     // 1. 赋值指令
     r = i
	defer func() {
		r++
	}()
	// 3. 空的return指令
     return
}

```

f5 拆解后：
```go
func f5() (r int) {   
     // 1. 赋值指令
     r = 1
	defer func() {
		r++
	}()
	// 3. 空的return指令
     return
}

```


```go
package main

import "fmt"

type number int

func (n number) print()   { fmt.Println(n) }
func (n *number) pprint() { fmt.Println(*n) }

func main() {
	var n number

	defer n.print()               //0
	defer n.pprint()              //3
	defer func() { n.print() }()  //3
	defer func() { n.pprint() }() //3
	n = 3
	// output:
	// 3
	// 3
	// 3
	// 0
}
```
第四个defer语句是闭包，引用外部函数的n, 最终结果是3;
第三个defer语句同第四个；
第二个defer语句，n是引用，最终求值是3.
第一个defer语句，对n直接求值，开始的时候n=0, 所以最后是0;

## 三条基本规则
### 1. defer函数是在外部函数return后，按照后申明先执行（栈）的顺序执行的
延迟函数执行按后进先出顺序执行，即先出现的defer最后执行

```go
package main

import "fmt"

func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
// output:
// 3
// 2
// 1
```

### 2. defer函数的参数的值，是在申明defer时确定下来de
延迟函数的参数在defer语句出现时就已经确定下来了
注意：对于指针类型参数，规则仍然适用，只不过延迟函数的参数是一个地址值，这种情况下，defer后面的语句对变量的修改可能会影响延迟函数。

```go
package main

import "fmt"

func main() {
    i := 0

    defer fmt.Println(i)  //(0) 这也算是作为defer函数的参数
    defer func(j int) { fmt.Println(j) }(i)  //(0) 作为参数
    defer func() { fmt.Println(i) }()  //(1) 作为闭包（closure）进行引用

    i++
}
// output:
// 1
// 0
// 0
```

```go
package main

import "fmt"

type Person struct {
	name string
}

func main() {
	person := &Person{"Lilei"}

	defer fmt.Println(person.name)                        // person.name作为普通类型当做defer函数的参数
	defer fmt.Printf("%v\n", person)                      // 引用类型作为参数
	defer func(p *Person) { fmt.Println(p.name) }(person) // 同上
	defer func() { fmt.Println(person.name) }()           // 闭包引用，对引用对象属性的修改不影响引用

	person.name = "HanMeimei"
}
// output:
// HanMeimei
// HanMeimei
// &{HanMeimei}
// Lilei 
```

```go
package main

import "fmt"

type Person struct {
    name string
}

func main() {
    person := &Person{"Lilei"}

    defer fmt.Println(person.name)  // 同上，person.name作为普通类型当做defer函数的参数
    defer fmt.Printf("%v\n", person)  // 作为defer函数的参数，申明时指向“Lilei”
    defer func(p *Person) { fmt.Println(p.name) }(person)  // 同上
    defer func() { fmt.Println(person.name) }()  // 作为闭包引用，随着person的改变而指向“HanMeimei”

    person = &Person{"HanMeimei"}
}
// output:
// HanMeimei
// Lilei
// &{Lilei}
// Lilei 
```
### 3. defer函数可以读取和修改外部函数申明的返回值。
延迟函数可能操作主函数的具名返回值

```go
package main

import "fmt"

func main() {
	fmt.Printf("output: %d\n", f())
}

func f() (i int) {
	defer fmt.Println(i)                    // 参数引用
	defer func(j int) { fmt.Println(j) }(i) // 同上
	defer func() { fmt.Println(i) }()       // 闭包引用
	defer func() { i++ }()                  // 执行前，i=2
	defer func() { i++ }()                  // 执行前，i=1

	i++

	return
}

```
output
```
3
0
0
output: 3
```

### xxx
```go
package main

import "fmt"

func main() {
	var whatever [3]struct{}

	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}
// 2
// 2
// 2
```

##  defer实现原理
### 1 defer数据结构
源码包src/src/runtime/runtime2.go:_defer定义了defer的数据结构：
```go
type _defer struct {
    sp      uintptr   //函数栈指针
    pc      uintptr   //程序计数器
    fn      *funcval  //函数地址
    link    *_defer   //指向自身结构的指针，用于链接多个defer
}
```

我们知道defer后面一定要接一个函数的，所以defer的数据结构跟一般函数类似，也有栈地址、程序计数器、函数地址等等。

与函数不同的一点是它含有一个指针，可用于指向另一个defer，每个goroutine数据结构中实际上也有一个defer指针，该指针指向一个defer的单链表，每次声明一个defer时就将defer插入到单链表表头，每次执行defer时就从单链表表头取出一个defer执行。

下图展示一个goroutine定义多个defer时的场景：

![](img/defer-1.jpg)
从上图可以看到，新声明的defer总是添加到链表头部。

函数返回前执行defer则是从链表首部依次取出执行，不再赘述。

一个goroutine可能连续调用多个函数，defer添加过程跟上述流程一致，进入函数时添加defer，离开函数时取出defer，所以即便调用多个函数，也总是能保证defer是按FIFO方式执行的。

### 2 defer的创建和执行
源码包src/runtime/panic.go定义了两个方法分别用于创建defer和执行defer。

- deferproc()： 在声明defer处调用，其将defer函数存入goroutine的链表中；
- deferreturn()：在return指令，准确的讲是在ret指令前调用，其将defer从goroutine链表中取出并执行。

可以简单这么理解，在编译在阶段，声明defer处插入了函数deferproc()，在函数return前插入了函数deferreturn()。

return不是原子操作，执行过程是: 保存返回值(若有)-->执行defer（若有）-->执行ret跳转