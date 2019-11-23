High Performance Go Workshop（Part2）
===================================

2.基准化分析法
--------

  

> Measure twice and cut once. — Ancient proverb

  

在我们试图提高一段代码的性能之前，首先我们必须了解它当前的性能。

  

本节重点介绍如何使用go测试框架构建有用的基准，并给出避免这些陷阱的实用技巧。

  

### 2.1. 基准的基本原则

  

在基准测试之前，必须有一个稳定的环境才能获得可重复的结果。

  

*   机器必须空闲——不要在共享硬件上配置文件，不要在等待长时间基准运行时浏览web 网页。

  

*   注意节能和热缩放。这些在现代笔记本电脑上几乎是不可避免的。

  

*   避免使用虚拟机和共享云主机；对于一致的测量，它们可能太嘈杂。

  

如果你有钱的话，就买专用的性能测试硬件。机架安装后禁用所有的电源管理和热量缩放，永远不要更新这些机器上的软件。从系统管理员的角度来看，最后一点是糟糕的建议，但是如果软件更新改变了内核或库的执行方式——想想 Spectre 补丁——这将使以前的基准测试结果失效。

  

对于我们其他人，有一个运行前后的样本，并运行它们多次，以获得一致的结果。

  

### 2.2. 使用 test package 进行基准测试

  

`testing`package 内置了对编写基准测试的支持。如果我们有这样一个简单的函数：

```golang
func Fib(n int) int {  
    switch n { 
        case 0: return 0 
        case 1: return 1 
        case 2: return 2 
        default: return Fib(n-1) + Fib(n-2) 
    } 
}
```
  

我们可以使用 `testing` package 采用这种形式为函数编写基准测试。

  
```golang
func BenchmarkFib20(b *testing.B) {  
    for n := 0; n < b.N; n++ { 
        Fib(20) // run the Fib function b.N times 
    } 
}
```
  

> 基准测试函数与测试共存，放在 `_test.go` 文件里。  

  

基准测试与测试类似，唯一真正的区别是它们采用的是 a *testing.B 而不是 a *testing.T。这两种类型都实现了 testing.TB 接口，该接口提供了 Errorf()、Fatalf() 和 FailNow()  等。

  

#### 2.2.1. 运行 package’s benchmarks

  

由于基准使用 `testing` 包，因此它们通过 go test 命令执行。但是，默认情况下，当调用 go test 时，基准被排除在外。

  

要在包中显式运行基准测试，请使用 `-bench` flag。--bench 接受一个正则表达式，该表达式与要运行的基准的名称匹配，因此调用包中所有基准的最常用方法是 -bench=.

  

下面是一个例子：

  
```
% go test -bench=. ./examples/fib/
goos: darwin
goarch: amd64
BenchmarkFib20-8           30000             40865 ns/op 
PASS 
ok      _/Users/dfc/devel/high-performance-go-workshop/examples/fib     1.671s
```
  

  

> `go test` 也会在匹配在基准之前的所有测试，所以如果你在一个包中有很多测试，或者它们需要很长时间才能运行，你可以通过提供一个不匹配的正则表达式来排除它们。  
>   
> 
> go test -run=^$

  

  

#### 2.2.2. 基准测试工作原理

  

每个基准函数用不同的 b.N 值调用，这是基准应该运行的迭代次数。

  

b.N 从1开始，如果基准函数在1秒内完成，则 b.N 增加，基准函数再次运行。

  

b.N 以近似的顺序增加；1、2、3、5、10、20、30、50、100 等等。基准测试框架试图变得聪明，如果它看到 b.N 的小值相对较快地完成，它将更快地增加迭代次数。

  

在上面的例子中，BenchmarkFib20-8 发现大约 30000 次循环迭代只需要一秒钟。在此基础上，基准框架计算出每个操作的平均时间为 40865 ns。

  

> `-8` 后缀用于运行该测试 `GOMAXPROCS` 的值。 这个数字,（如 `GOMAXPROCS` ）, 默认为Go 进程可见的的数量。你可以使用 `-cpu` 命令修改此值，该命令需要值列表来运行基准测试。

  
```
% go test -bench=. -cpu=1,2,4 ./examples/fib/ 
goos: darwin 
goarch: amd64 
BenchmarkFib20             30000             39115 ns/op 
BenchmarkFib20-2           30000             39468 ns/op 
BenchmarkFib20-4           50000             40728 ns/op 
PASS 
ok      _/Users/dfc/devel/high-performance-go-workshop/examples/fib     5.531s
```
  

> 这显示了以 1、2 和 4 核运行基准测试。在这种情况下，该 flag 对结果几乎没有影响，因为该基准是完全顺序的。

  

  

#### 2.2.3. 提高基准测试的精度

  

`fib` 函数是一个人为故意的例子——除非你写了 TechPower web server 服务器基准——否则你的业务不太可能依赖于计算 Fibonaci 序列中第 20 个数字的速度。但是，基准测试确实提供了一个有效基准测试示例。

  

具体来说，你希望你的基准测试运行数万次迭代，以便获得每个操作的良好平均值。如果你的基准测试只运行100或10次迭代，那么这些运行的平均值可能有很高的标准偏差。如果你的基准测试运行了数百万或数十亿次迭代，那么平均值可能非常精确，但受制于代码布局和对齐的不确定性。

  

要增加迭代次数，可以使用 -benchtime 标志增加基准时间。例如：

  
```
% go test -bench=. -benchtime=10s ./examples/fib/ 
goos: darwin 
goarch: amd64 
BenchmarkFib20-8          300000             39318 ns/op 
PASS 
ok      _/Users/dfc/devel/high-performance-go-workshop/examples/fib     20.066s
```
  

运行相同的基准测试，直到它达到 b.N 的值，返回时间超过10秒。当我们运行的时间延长了10倍时，迭代的总数就增加了10倍。结果没有太大变化，这是我们所期望的。

  

> 为什么总时间是20秒，而不是10秒？

  

如果你有一个基准测试，该基准测试运行数百万次或数十亿次，导致每次操作的时间在微或纳米秒范围内，那么你可能会发现，由于热缩放、内存局部性、后台处理、GC 等原因，基准测试是不稳定的。

  

对于每次操作以 10 纳秒或一位数纳秒计算的时间，指令重新排序和代码对齐的相对论效应将对基准时间产生影响。

  

可以使用 -count 命令多次处理此运行基准：

  
```
% go test -bench=Fib1 -count=10 ./examples/fib/ 
goos: darwin 
goarch: amd64 
BenchmarkFib1-8         2000000000               1.99 ns/op 
BenchmarkFib1-8         1000000000               1.95 ns/op 
BenchmarkFib1-8         2000000000               1.99 ns/op 
BenchmarkFib1-8         2000000000               1.97 ns/op 
BenchmarkFib1-8         2000000000               1.99 ns/op 
BenchmarkFib1-8         2000000000               1.96 ns/op 
BenchmarkFib1-8         2000000000               1.99 ns/op 
BenchmarkFib1-8         2000000000               2.01 ns/op 
BenchmarkFib1-8         2000000000               1.99 ns/op 
BenchmarkFib1-8         1000000000               2.00 ns/op
```
  

Fib(1) 的基准大约需要2纳秒，方差为 \+/\- 2%。

  

Go1.12中的新特性是 -benchtime 标志现在需要多次迭代，例如 -benchtime=20x，这将精确地运行你的代码 `benchtime`时间。

  
尝试以10x，20x，50x，100x和300x的-benchtime运行上面的fib工作台。你看到了什么？                                                                                                                         

  

> 如果你发现 `go test`   应用的默认值需要针对 particular package 进行调整, 我建议在  `Makefile`  中编写这些设置，以便每个想要运行基准测试的人都可以使用相同的设置执行此操作。

  

  

  

### 2.3. 将基准与 benchstat 进行比较

  

在上一节中，我建议不止运行一次基准测试，以获得更多的平均数据。对于任何基准测试，这都是一个很好的建议，因为我在本章开头提到了电源管理、后台流程和热管理的效果。

  
我要介绍一个叫做 [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) 的工具。

  
```
% go get golang.org/x/perf/cmd/benchstat
```
  

Benchstat 可以进行一组基准测试，并告诉它们它们有多稳定。这是一个 Fib(20) 的例子。

  
```
% go test -bench=Fib20 -count=10 ./examples/fib/ | tee old.txt 
goos: darwin 
goarch: amd64 
BenchmarkFib20-8           50000             38479 ns/op 
BenchmarkFib20-8           50000             38303 ns/op 
BenchmarkFib20-8           50000             38130 ns/op 
BenchmarkFib20-8           50000             38636 ns/op 
BenchmarkFib20-8           50000             38784 ns/op 
BenchmarkFib20-8           50000             38310 ns/op 
BenchmarkFib20-8           50000             38156 ns/op 
BenchmarkFib20-8           50000             38291 ns/op 
BenchmarkFib20-8           50000             38075 ns/op 
BenchmarkFib20-8           50000             38705 ns/op 
PASS 
ok      _/Users/dfc/devel/high-performance-go-workshop/examples/fib     23.125s 
% benchstat old.txt 
name     time/op 
Fib20-8  38.4µs ± 1%
```
  

  

`benchstat`告诉我们，平均值是38.8微秒，整个样本的变化率为±2%。这对电池很有好处。

  

*   第一次运行是最慢的一次，因为操作系统为了省电而关闭了CPU。

  

*   接下来的两次运行是最快的，因为操作系统认为这不是一个短暂的工作高峰，它提高了时钟速度，以便尽快完成工作，希望能够重新进入睡眠状态。

  

*   剩下的运行是操作系统和bios交换热量生产的功耗。

  

#### 2.3.1. 改进 `Fib`

  

确定两组基准测试之间的性能增量可能会非常繁琐且容易出错。Benchstat 可以帮我们解决这个问题。

  

>   
> 保存基准测试运行的输出是有用的，但是你也可以保存生成它的二进制文件。这允许你重新运行基准测试之前的迭代。为此，使用 `-c` 标志来保存测试二进制文件。我经常将这个二进制文件从 `.test` 重命名为 `.golden`。
> 
```
% go test -c 
% mv fib.test fib.golden
```
  

先前的 `Fib`函数对 fibonaci 序列中的第 0 个和第 1 个进行了硬编码。之后，代码递归地调用自己。今天稍后我们将讨论递归的成本，但目前假设它有成本，特别是当我们的算法使用指数时间的时候。

  

解决这个问题的简单方法是从斐波那契数列中硬编码另一个数，将每次重复调用的深度减少一个。

  
```golang
func Fib(n int) int {  
    switch n { 
        case 0: return 0 
        case 1: return 1 
        case 2: return 1 
        default: return Fib(n-1) + Fib(n-2) 
    } 
}
```
  

  

> 该文件还包含针对 Fib 的全面测试。如果没有通过验证当前行为的测试，请勿尝试改进基准。

  

为了比较我们的新版本，我们编译了一个新的测试二进制文件并对它们进行基准测试，然后使用 `benchstat`比较输出。

  
```
% go test -c 
% ./fib.golden -test.bench=. -test.count=10 > old.txt 
% ./fib.test -test.bench=. -test.count=10 > new.txt 
% benchstat old.txt new.txt
 name     old time/op  new time/op  delta 
 Fib20-8  44.3µs ± 6%  25.6µs ± 2%  -42.31%  (p=0.000 n=10+10)
```
  

比较基准时需要检查三件事

  

*   最新的和旧的方差 ± 1-2% 是好的，3-5% 是可以的，大于 5%，一些样品将被认为是不可靠的。在比较基准的时候要小心，如果一方有很高的方差，你可能看不到改进。

  

*   P值。P值低于 0.05 是好的，大于 0.05 意味着基准可能不具有统计学意义。

  

*   遗漏样本。benchstat 将报告它认为有效的新旧样本的数量，有时你可能只会发现报告了9个，即使你确实 -count=10。10% 或更低的拒绝率是可以的，高于10%可能表明你的设置是不稳定的，你可能比较的样本太少。

  

  

### 2.4. 避免基准化启动成本

  

有时你的基准测试运行的时候有设置成本。b.ResetTimer() 将被用来忽略设置过程中消耗的时间。

  
```golang
func BenchmarkExpensive(b *testing.B) {  
    boringAndExpensiveSetup() 
    b.ResetTimer() 
    for n := 0; n < b.N; n++ { 
        // function under test 
    } 
}
```
  

  

1.  重置基准计时器

  

如果每次循环迭代都有一些开销比较大的设置逻辑，请使用 b.StopTimer() 和 b.StartTimer() 暂停基准计时器。

  
```golang
func BenchmarkComplicated(b *testing.B) {  
    for n := 0; n < b.N; n++ { 
        b.StopTimer() 
        complicatedSetup() 
        b.StartTimer() 
        // function under test 
    } 
}
```
  

  

1.  暂停基准计时器
2.  恢复计时器

  

  

### 2.5. 基准配置

  

分配数量和大小与基准时间密切相关。你可以告诉 `testing`框架记录被测代码分配的数量。

  
```golang
func BenchmarkRead(b *testing.B) {  
    b.ReportAllocs() for n := 0; n < b.N; n++ { 
        // function under test 
    } 
}
```
  

这是使用 `bufio`软件包基准测试的示例。

  
```
% go test -run=^$ -bench=. bufio 
goos: darwin 
goarch: amd64 
pkg: bufio 
BenchmarkReaderCopyOptimal-8            20000000               103 ns/op 
BenchmarkReaderCopyUnoptimal-8          10000000               159 ns/op 
BenchmarkReaderCopyNoWriteTo-8            500000              3644 ns/op 
BenchmarkReaderWriteToOptimal-8          5000000               344 ns/op 
BenchmarkWriterCopyOptimal-8            20000000                98.6 ns/op 
BenchmarkWriterCopyUnoptimal-8          10000000               131 ns/op 
BenchmarkWriterCopyNoReadFrom-8           300000              3955 ns/op 
BenchmarkReaderEmpty-8                   2000000               789 ns/op            4224 B/op          3 allocs/op 
BenchmarkWriterEmpty-8                   2000000               683 ns/op            4096 B/op          1 allocs/op BenchmarkWriterFlush-8                  100000000               17.0 ns/op             0 B/op          0 allocs/op
```
  

  

> 你还可以使用 go test -benchmem 标志来强制测试框架报告所有基准运行的分配统计信息。
```
% go test -run=^$ -bench=. -benchmem bufio 
goos: darwin 
goarch: amd64 
pkg: bufio 
BenchmarkReaderCopyOptimal-8            20000000                93.5 ns/op            16 B/op          1 allocs/op 
BenchmarkReaderCopyUnoptimal-8          10000000               155 ns/op              32 B/op          2 allocs/op 
BenchmarkReaderCopyNoWriteTo-8            500000              3238 ns/op           32800 B/op          3 allocs/op 
BenchmarkReaderWriteToOptimal-8          5000000               335 ns/op              16 B/op          1 allocs/op 
BenchmarkWriterCopyOptimal-8            20000000                96.7 ns/op            16 B/op          1 allocs/op 
BenchmarkWriterCopyUnoptimal-8          10000000               124 ns/op              32 B/op          2 allocs/op 
BenchmarkWriterCopyNoReadFrom-8           500000              3219 ns/op           32800 B/op          3 allocs/op 
BenchmarkReaderEmpty-8                   2000000               748 ns/op            4224 B/op          3 allocs/op 
BenchmarkWriterEmpty-8                   2000000               662 ns/op            4096 B/op          1 allocs/op 
BenchmarkWriterFlush-8                  100000000               16.9 ns/op             0 B/op          0 allocs/op PASS ok      bufio   20.366s
```
  

### 2.6. 观察编译器优化

  

这个例子来自[issue 14813](https://github.com/golang/go/issues/14813#issue-140603392)

  
```golang
const m1 = 0x5555555555555555 
const m2 = 0x3333333333333333 
const m4 = 0x0f0f0f0f0f0f0f0f 
const h01 = 0x0101010101010101   
func popcnt(x uint64) uint64 { 
     x -= (x >> 1) & m1 
     x = (x & m2) + ((x >> 2) & m2) 
     x = (x + (x >> 4)) & m4 
     return (x * h01) >> 56 
}   
func BenchmarkPopcnt(b *testing.B) {  
    for i := 0; i < b.N; i++ { 
        popcnt(uint64(i)) 
    } 
}
```
  

你认为这个函数的基准测试速度有多快？让我们看看。

  
```
% go test -bench=. ./examples/popcnt/ 
goos: darwin 
goarch: amd64 
BenchmarkPopcnt-8       2000000000               0.30 ns/op 
PASS
```
  

0.3 纳米秒；基本上是一个时钟周期。即使假设 CPU 在每个时钟周期内可能有一些指令在运行，这个数字似乎也低得不合理。这是咋回事？

  

为了理解发生了什么，我们必须查看 benchmake下的函数 popcnt。popcnt 是一个leaf 函数-；它不调用任何其他函数——因此编译器可以内联它。

  

由于该函数是内联的，编译器现在可以看到它没有任何副作用。`popcnt` 不影响任何全局变量的状态。因此，该调用被消除。这是编译器看到的：

  
```golang
func BenchmarkPopcnt(b *testing.B) {  
    for i := 0; i < b.N; i++ { 
        // optimised away 
    } 
}
```
  

在我测试过的所有Go编译器版本中，仍然会生成循环。但是英特尔 CPU 确实擅长优化循环，尤其是空循环。

  

#### 2.6.1. 练习, 看编译

  

在我们继续之前，让我们看一下编译，以确认我们所看到的

  
```
% go test -gcflags=-S
```
  

使用 `gcflags =“-l -S”禁用内联，这如何影响编译的输出

  

> 优化是件好事
> 
>   
> 
> 我们要做的是，通过删除不必要的计算，使真正的代码快速运行的优化与删除没有可观察到的副作用的基准的优化是一样的。
> 
>   
> 
> 随着 Go编译器的改进，这种情况会变得更加常见。

  

  

#### 2.6.2. 修复基准

  

禁用内联使基准工作是不现实的；我们希望通过优化来构建我们的代码。

  

要修复此基准测试，我们必须确保编译器无法证明 `BenchmarkPopcnt`正文不会导致全局状态更改。

  
```golang
var Result uint64   
func BenchmarkPopcnt(b *testing.B) { 
     var r uint64 
     for i := 0; i < b.N; i++ { 
         r = popcnt(uint64(i)) 
     } 
     Result = r 
}
```
  

  

这是确保编译器无法优化循环正文的推荐方法。

  

首先，我们使用调用 `popcnt`的结果，将其存储在 r 中。其次，由于一旦基准测试结束，r 在基准级操作范围内本地声明，因此 r 的结果对程序的另一部分永远不会可见，因此作为最终操作，我们将 r 的值分配给公共变量 Result。

  

由于 `Result`是公共的，编译器无法证明导入此包的另一个包将无法看到结果的值随时间而变化，因此无法优化导致其分配的任何操作。

  

如果我们直接分配给 `Result`会怎样？这会影响基准时间吗？如果我们将 `popcnt`的结果赋给 _ 怎么办？

  

> 在我们之前的 `Fib`  基准测试中，我们没有采取这些预防措施，应该这样做吗？

  

  

### 2.7. 基准错误

  

`for`循环对基准的操作至关重要。

  

这里有两个不正确的基准，你能解释一下它们有什么问题吗？

  
```golang
func BenchmarkFibWrong(b *testing.B) {  
    Fib(b.N) 
}
```
  
```golang
func BenchmarkFibWrong2(b *testing.B) {  
    for n := 0; n < b.N; n++ { 
        Fib(n) 
    } 
}
```
  

  

运行这两个基准测试，你看到了啥？

  

  

### 2.8. 分析基准

  

`testing`包内置了对生成 CPU、内存和阻塞分析的支持。

  

*   `-cpuprofile=$FILE` 将 CPU 分析写入 $FILE。

  

*  ` -memprofile=$FILE`，将内存分析写入 $FILE，memprofilerate=N 将分析速率调整为 1/N。

  

*   `-blockprofile=$FILE`，将阻塞分析写入 $FILE。

  

使用这些命令中的任何一个也会保留二进制文件。

  
```
% go test -run=XXX -bench=. -cpuprofile=c.p bytes % 
go tool pprof c.p
```
 