[Go 编译器内核：给 Go 新增一个语句 ](https://studygolang.com/articles/25101)

从Go官方工具链1.16开始，当GODEBUG环境变量包含inittrace=1时，Go运行时将会报告各个源代码文件中的init函数的执行时间和内存开辟消耗情况。比如对于下面的程序
`GODEBUG=inittrace=1 go run main.go`