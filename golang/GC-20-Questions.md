# [Go GC 20 问](https://blog.csdn.net/qcrao/article/details/103856176)

本目录Go-Questions下有源码


## GODEBUG
https://go.dev/pkg/runtime/#hdr-Environment_Variables

`GCDEBUG=name1=val1,name2=val2`

```


GODEBUG=gctrace=1  #prints garbage collector events at each collection, summarizing the amount of memory collected and the length of the pause.

GODEBUG=gcpacertrace=1    # causes the garbage collector to print information about the internal state of the concurrent pacer.

GODEBUG=memprofilerate=X  # update the value of runtime.MemProfileRate

GODEBUG=inittrace=1  #prints a summary of execution time and memory allocation information for completed package initialization work.

GODEBUG=schedtrace=X  #prints scheduling events every X milliseconds.
```

https://golang.org/pkg/net/
```
GODEBUG=netdns=go    # DNS相关，force pure Go resolver
GODEBUG=netdns=cgo   # DNS相关，force cgo resolver
GODEBUG=netdns=1     # DNS相关，print its decisions，这个比较有用会告诉你合适发生了DNS解析，以及DNS解析的尝试顺序

```

https://golang.org/pkg/net/http/
```
GODEBUG=http2client=0  # disable HTTP/2 client support
GODEBUG=http2server=0  # disable HTTP/2 server support
GODEBUG=http2debug=1   # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2   # ... even more verbose, with frame dumps
```

## GOGC
GOGC 用于控制GC的处发频率， 其值默认为100, 

意为直到自上次垃圾回收后heap size已经增长了100%时GC才触发运行。即是GOGC=100意味着live heap size 每增长一倍，GC触发运行一次。

如设定GOGC=200, 则live heap size 自上次垃圾回收后，增长2倍时，GC触发运行， 总之，其值越大则GC触发运行频率越低， 反之则越高，

 如果GOGC=off 则关闭GC.

虽然go 1.5引入了低延迟的GC, 但是GOGC对GC运行频率的影响不变， 仍然是其值大于100,则越大GC运行频率越高，
反之则越低。

关闭GC：
方法一：设置环境变量 GOGC=off
方法二：运行时调用 debug.SetGCPercent(-1