
## go 1.12

- [踩坑记：go服务内存暴涨](https://blog.csdn.net/felix021/article/details/106066765)
> ## Runtime 

>Go 1.12 significantly improves the performance of sweeping when a large fraction of the heap remains live. This reduces allocation latency immediately following a garbage collection.

> (中间省略2段不太相关的内容)

> On Linux, the runtime now uses MADV_FREE to release unused memory. This is more efficient but may result in higher reported RSS. The kernel will reclaim the unused data when it is needed. 
>golang.org/doc/go1.12

翻译一下：

> 在堆内存大部分活跃的情况下，go 1.12 可以显著提高清理性能，降低 [紧随某次gc的内存分配] 的延迟。
> 在Linux上，Go Runtime现在使用 MADV_FREE 来释放未使用的内存。这样效率更高，但是可能导致更高的 RSS；内核会在需要时回收这些内存。

> 在 Linux Kernel 4.5 之前，只支持 MADV_DONTNEED（go 1.11 及以前默认)