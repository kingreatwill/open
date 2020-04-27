# go语言原本
https://github.com/changkun/go-under-the-hood
记录一些可能会用到的信息

## 汇编代码
https://changkun.de/golang/zh-cn/part1basic/ch01proc/asm/
对于一段 Go 程序，我们可以通过下面的命令来获得编译后的汇编代码：
```
go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o main.out main.go
go tool objdump -s "main.main" main.out > main.S
# or
go tool compile -S main.go
# or
go build -gcflags -S main.go
```
FUNCDATA 和 PCDATA 指令包含了由垃圾回收器使用的信息，他们由编译器引入。

