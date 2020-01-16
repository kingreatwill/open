[深度解密Go语言之 pprof](https://qcrao.com/2019/11/10/dive-into-go-pprof/)

[深度解密Go语言之 pprof](https://www.cnblogs.com/qcrao-2018/p/11832732.html)


堆的信息
```
go tool pprof -http :9090 http://ip:port/debug/pprof/heap
```

pprof可以比较两个时间点的分配的内存的差值

1. 首先确保你已经配置了 pprof 的 http 路径， 可以访问`http://ip:port/debug/pprof/`查看(如果你没有修改默认的 pprof 路径)
2. 导出**时间点 1**的堆的 profile: `curl -s http://127.0.0.1:8080/debug/pprof/heap > base.heap`, 我们把它作为基准点
3. 喝杯茶，等待一段时间后导出**时间点 2**的堆的 profile: `curl -s http://127.0.0.1:8080/debug/pprof/heap > current.heap`
4. 现在你就可以比较这两个时间点的堆的差异了: `go tool pprof --base base.heap current.heap`

5. 使用web命令会生成一个 SVG 文件，可能你需要使用浏览器打开它。
或者你直接使用命令打开 web 界面: go tool pprof --http :9090 --base base.heap current.heap。


通过比较差值，就容易看到哪些地方产生的内存"残留"的比较多，没有被内存释放，极有可能是内存泄漏的点。
