# CLR的Garbage Collection

https://www.cnblogs.com/dacc123/p/10980718.html

前言
--

在阅读这篇文章：[Announcing Net Core 3 Preview3](https://devblogs.microsoft.com/dotnet/announcing-net-core-3-preview-3/)的时候，我看到了这样一个特性：

> Docker and cgroup memory Limits

> We concluded that the primary fix is to set a GC heap maximum significantly lower than the overall memory limit as a default behavior. In retrospect, this choice seems like an obvious requirement of our implementation. We also found that Java has taken a similar approach, introduced in Java 9 and updated in Java 10.

大概的意思呢就是在 .NET Core 3.0 版本中，我们已经通过修改 GC 堆内存的最大值，来避免这样一个情况：在 docker 容器中运行的 .NET Core 程序，因为 docker 容器内存限制而被 docker 杀死。

恰好，我在 docker swarm 集群中跑的一个程序，总是被 docker 杀死，大都是因为内存超出了限制。那么升级到 .NET Core 3.0 是不是会起作用呢？这篇文章将浅显的了解 .NET Core 3.0 的 `Garbage Collection` 机制，以及 Linux 的 `Cgroups` 内核功能。最后再写一组 实验程序 去真实的了解 .NET Core 3.0 带来的 GC 变化。

GC
--

### CLR

.NET 程序是运行在 CLR : Common Language Runtime 之上。CLR 就像 JAVA 中的 JVM 虚拟机。CLR 包括了 JIT 编译器，GC 垃圾回收器，CIL CLI 语言标准。

那么 .NET Core 呢？它运行在 `CoreCLR 上`，是属于 .NET Core 的 Runtime。二者大体我觉得应该差不多吧。所以我介绍 CLR 中的一些概念，这样才可以更好的理解 GC

*   我们的程序都是在操作虚拟内存地址，从来不直接操作内存地址，即使是 Native Code。
    
*   一个进程会被分配一个独立的虚拟内存空间，我们定义的和管理的对象都在这些空间之中。  
    虚拟内存空间中的内存 有三种状态：空闲 (可以随时分配对象)，预定 (被某个进程预定，尚且不能分配对象)，提交（从物理内存中分配了地址到该虚拟内存，这个时候才可以分配对象）
    
*   CLR 初始化GC 后，GC 就在上面说的虚拟内存空间中分配内存，用来让它管理和分配对象，被分配的内存叫做 `Managed Heap` **管理堆**，每个进程都有一个管理堆内存，进程中的线程共享一个管理堆内存
    
*   CLR 中还有一块堆内存叫做`LOH` Large Object Heap 。它也是隶属于 GC 管理，但是它很特别，只分配大于 85000byte 的对象，所以叫做大对象，为什么要这么做呢？很显然大对象太难管理了，GC 回收大对象将很耗时，所以没办法，只有给这些 “大象” 另选一出房子，GC 这个“管理员” 很少管 “大象”。
    

那么什么时候对象会被分配到堆内存中呢？

所有引用类型的对象，以及作为类属性的值类型对象，都会分配在堆中。大于 85000byte 的对象扔到 “大象房” 里。

堆内存中的对象越少，GC 干的事情越少，你的程序就越快，因为 GC 在干事的时候，程序中的其他线程都必须毕恭毕敬的站着不动（挂起），等 GC 说：我已经清理好了。然后大家才开始继续忙碌。所以 GC 一直都是在干帮线程擦屁股的事情。

所以没有 GC 的编程语言更快，但是也更容易产生废物。

### GC Generation

那么 GC 在收拾垃圾的过程中到底做了什么呢？首先要了解 CLR 的 GC 有一个`Generation` \*\*代 \*\* 的概念 GC 通过将对象分为三代，优化对象管理。GC 中的代分为三代：

*   `Generation 0` 零代或者叫做初代，初代中都是一些短命的对象，shorter object，它们通常会被很快清除。当 new 一个新对象的时候，该对象都会分配在 Generation 0 中。**只有一段连续的内存**
    
*   `Generation 1` 一代，一代中的对象也是短命对象，它相当于 shorter object 和 longer object 之间的缓冲区。**只有一段连续的内存**
    
*   `Generation 2` 二代，二代中的对象都是长寿对象，他们都是从零代和一代中选拔而来，一旦进入二代，那就意味着你很安全。之前说的 LOH 就属于二代，static 定义的对象也是直接分配在二代中。**包含多段连续的内存**。
    

零代和一代 占用的内存因为他们都是短暂对象，所以叫做短暂内存块。 那么他们占用的内存大小是多大？32位和63位的系统是不一样的，不同的GC类型也是不一样的。

WorkStation GC:

32 位操作系统 16MB ，64位 操作系统 256M

Server GC：

32 w位操作系统 65MB，64 位操作系统 4GB！

### GC 回收过程

当 管理堆内存中使用到达一定的阈值的时候，这个阈值是GC 决定的，或者系统内存不够用的时候，或者调用 `GC.Collect()` 的时候，GC 都会立刻可以开始回收，没有商量的余地。于是所有线程都会被挂起（也并不都是这样）

GC 会在 Generation 0 中开始巡查，如果是 死对象，就把他们的内存释放，如果是 活的对象，那么就标记这些对象。接着把这些活的对象升级到下一代：移动到下一代 Generation 1 中。

同理 在 Generation 1 中也是如此，释放死对象，升级活对象。

三个 Generation 中，Generation 0 被 GC 清理的最频繁，Generation 1 其次，Generation 2 被 GC 访问的最少。因为要清理 Generation 2 的消耗太大了。

GC 在每一个 Generation 进行清理都要进行三个步骤：

*   标记： GC 循环遍历每一个对象，给它们标记是 死对象 还是 活对象
    
*   重新分配：重新分配活对象的引用
    
*   清理：将死对象释放，将活对象移动到下一代中
    

### WorkStation GC 和 Server GC

GC 有两种形式：`WorkStation GC`和 `Server GC`

[默认的.NET](http://xn--hxyp46bm5n.NET) 程序都是 WorkStation GC ，那么 WorkStation GC 和 Server GC 有什么区别呢。

上面已经提到一个区别，那就是 Server GC 的 Generation 内存更大，64位操作系统 Generation 0 的大小居然有4G ，这意味着啥？在不调用`GC.Collect` 的情况下，4G 塞满GC 才会去回收。那样性能可是有很大的提升。但是一旦回收了，4GB 的“垃圾” 也够GC 喝一壶的了。

还有一个很大的区别就是，Server GC 拥有专门用来处理 GC的线程，而WorkStation GC 的处理线程就是你的应用程序线程。WorkStation 形式下，GC 开始，所有应用程序线程挂起，GC选择最后一个应用程序线程用来跑GC，直到GC 完成。所有线程恢复。

而ServerGC 形式下： 有几核 CPU ，那么就有几个专有的线程来处理 GC。每个线程都一个堆进行GC ，不同的堆的对象可以相互引用。

所以在GC 的过程中，Server GC 比 WorkStation GC 更快。但是有专有线程，并不代表可以并行GC 哦。

上面两个区别，决定了 Server GC 用于对付高吞吐量的程序，而WorkStation GC 用于一般的客户端程序足以。

[如果你的.NET](http://xn--6qqv0wpxlfqo.NET) 程序正在疲于应付 高并发，不妨开启 Server GC : [https://docs.microsoft.com/en-us/dotnet/framework/configure-apps/file-schema/runtime/gcserver-element](https://docs.microsoft.com/en-us/dotnet/framework/configure-apps/file-schema/runtime/gcserver-element)

### Concurrent GC 和 Non-Concurrent GC

GC 有两种模式：`Concurrent` 和 `Non-Concurrent`，也就是并行 GC 和 不并行 GC 。无论是 Server GC 还是 Concurrent GC 都可以开启 Concurrent GC 模式或者关闭 Concurrent GC 模式。

Concurrent GC 当然是为了解决上述 GC 过程中所有线程挂起等待 GC 完成的问题。因为工作线程挂起将会影响 用户交互的流畅性和响应速度。

Concurrent 并行实际上 只发生在Generation 2 中，因为 Generation 0 和 Generation1 的处理是在太快了，相当于工作线程没有阻塞。

在 GC 处理 Generation 2 中的第一步，也就是标记过程中，工作线程是可以同步进行的，工作线程仍然可以在 Generation 0 和 Generation 1 中分配对象。

所以并行 GC 可以减少工作进程因为GC 需要挂起的时间。但是与此同时，在标记的过程中工作进程也可以继续分配对象，所以GC占用的内存可能更多。

而Non-Concurrent GC 就更好理解了。

.NET 默认开启了 Concurrent 模式，可以在 [https://docs.microsoft.com/en-us/dotnet/framework/configure-apps/file-schema/runtime/gcconcurrent-element](https://docs.microsoft.com/en-us/dotnet/framework/configure-apps/file-schema/runtime/gcconcurrent-element) 进行配置

### Background GC

又来了一种新的 GC 模式： `Background GC` 。那么 Background GC 和 Concurrent GC 的区别是什么呢？在阅读很多资料后，终于搞清楚了，因为英语水平不好。以下内容比较重要。

首先：Background GC 和 Concurrent GC 都是为了减少 因为 GC 而挂起工作线程的时间，从而提升用户交互体验，程序响应速度。

其次：Background GC 和 Concurrent GC 一样，都是使用一个专有的GC 线程，并且都是在 Generation 2 中起作用。

最后：Background GC 是 Concurrent GC 的增强版，[在.NET](http://xn--3ds.NET) 4.0 之前都是默认使用 Concurrent GC 而 .NET 4.0+ 之后使用Background GC 代替了 Concurrent GC。

那么 Background GC 比 Concurrent GC 多了什么呢：

之前说到 Concurrent GC 在 Generation 2 中进行清理时，工作线程仍然可以在 Generation 0/1 中进行分配对象，但是这是有限制的，当 Generation 0/1 中的内存片段 Segment 用完的时候，就不能再分配了，知道 Concurrent GC 完成。而 Background GC 没有这个限制，为啥呢？因为 Background GC 在 Generation 2 中进行清理时，允许了 Generation 0/1 进行清理，也就说是当 Generation 0/1 的 Segment 用完的时候， GC 可以去清理它们，这个GC 称作 `Foreground GC` ( 前台GC ) ，Foreground GC 清理完之后，工作线程就可以继续分配对象了。

所以 Background GC 比 Concurrent GC 减少了更多 工作线程暂停的时间。

GC 的简单概念就到这里了以上是阅读大量英文资料的精短总结，如果有写错的地方还请斧正。

作为最后一句总结GC的话：并不是使用了 Background GC 和 Concurrent GC 的程序运行速度就快，它们只是提升了用户交互的速度。因为 专有的GC 线程会对CPU 造成拖累，此外GC 的同时，工作线程分配对象 和正常的时候分配对象 是不一样的，它会对性能造成拖累。

.NET Core 3.0 的变化
-----------------

*   堆内存的大小进行了限制：max (20mb , 75% of memory limit on the container)
    
*   ServerGC 模式下 默认的Segment 最小是16mb, 一个堆 就是 一个segment。这样的好处可以举例来说明，比如32核服务器，运行一个内存限制32 mb的程序，那么在Server GC 模式下，会分配32个Heap,每个Heap 大小是1mb。但是现在，只需要分配2个Heap,每个Heap 大小16mb。
    
*   其他的就不太了解了。
    

实际体验
----

从开头的 介绍 [ASP.NET](http://ASP.NET) Core 3.0 文章中了解到 ，在 Docker 中，对容器的资源限制是通过 cgroup 实现的。cgroup 是 Linux 内核特性，它可以限制 进程组的 资源占用。当容器使用的内存超出docker的限制，docker 就会将改容器杀死。在之前 .NET Core 版本中，经常出现 .NET Core 应用程序消耗内存超过了docker 的 内存限制，从而导致被杀死。[而在.NET](http://xn--3ds946g.NET) Core 3.0 中这个问题被解决了。

为此我做了一个实验。

这是一段代码：

    using System;
    using System.Collections.Generic;
    using System.Threading;
    
    namespace ConsoleApp1
    {
        class Program
        {
            static void Main(string[] args)
            {
                    if (GCSettings.IsServerGC == true)
                        Console.WriteLine("Server GC");
                else
                        Console.WriteLine("GC WorkStationGC");
                byte[] buffer;
                for (int i = 0; i <= 100; i++)
                {
                    buffer = new byte[ 1024 * 1024];
    
    
                    Console.WriteLine($"allocate number {i+1} objet ");
    
                    var num = GC.CollectionCount(0);
                    var usedMemory = GC.GetTotalMemory(false) /1024 /1024;
    
                    Console.WriteLine($"heap use {usedMemory} mb");
                    Console.WriteLine($"GC occurs {num} times");
                    Thread.Sleep(TimeSpan.FromSeconds(5));
               }
            }
        }
    }
    
    

这段代码是在 for 循环 分配对象。`buffer = new byte[1024 * 1024]` 占用了 1M 的内存  
这段代码分别在 .NET Core 2.2 和 .NET Core 3.0 运行，完全相同的代码。运行的内存限制是 9mb

.NET Core 2.2 运行的结果是：

    GC WorkStationGC
    allocate number 1 objet 
    heap use 1 mb
    GC occurs 0 times
    allocate number 2 objet 
    heap use 2 mb
    GC occurs 0 times
    allocate number 3 objet 
    heap use 3 mb
    GC occurs 0 times
    allocate number 4 objet 
    heap use 1 mb
    GC occurs 1 times
    allocate number 5 objet 
    heap use 2 mb
    GC occurs 1 times
    allocate number 6 objet 
    heap use 3 mb
    GC occurs 1 times
    allocate number 7 objet 
    heap use 4 mb
    GC occurs 2 times
    allocate number 8 objet 
    heap use 5 mb
    GC occurs 3 times
    allocate number 9 objet 
    heap use 6 mb
    GC occurs 4 times
    allocate number 10 objet 
    heap use 7 mb
    GC occurs 5 times
    allocate number 11 objet 
    heap use 8 mb
    GC occurs 6 times
    allocate number 12 objet 
    heap use 9 mb
    
    Exit
    

[首先.NET](http://xn--44qr67o.NET) Core 2.2默认使用 WorkStation GC ，当heap使用内存到达9mb时，程序就被docker 杀死了。

[在.NET](http://xn--3ds.NET) Core 3.0 中

    GC WorkStationGC
    allocate number 1 objet 
    heap use 1 mb
    GC occurs 0 times
    allocate number 2 objet 
    heap use 2 mb
    GC occurs 0 times
    allocate number 3 objet 
    heap use 3 mb
    GC occurs 0 times
    allocate number 4 objet 
    heap use 1 mb
    GC occurs 1 times
    allocate number 5 objet 
    heap use 2 mb
    GC occurs 1 times
    allocate number 6 objet 
    heap use 3 mb
    GC occurs 1 times
    allocate number 7 objet 
    heap use 1 mb
    GC occurs 2 times
    allocate number 8 objet 
    heap use 2 mb
    GC occurs 2 times
    allocate number 9 objet 
    heap use 3 mb
    GC occurs 2 times
    ....
    

运行一直正常没问题。

二者的区别就是 .NET Core 2.2 GC 之后，堆内存没有减少。为什么会发生这样的现象呢？

一下是我的推测，没有具体跟踪GC的运行情况  
首先定义的占用 1Mb 的对象，由于大于 85kb 都存放在LOH 中，Large Object Heap，前面提到过。 GC 是很少会处理LOH 的对象的, 除非是 GC heap真的不够用了（一个GC heap包括 Large Object Heap 和 Small Object Heap）[由于.NET](http://xn--6kqu55g.NET) Core 3.0 对GC heap大小做了限制，所以当heap不够用的时候，它会清理LOH，[但是.NET](http://xn--gqq717c.NET) Core 2.2 下认为heap还有很多，所以它不清理LOH ，导致程序被docker杀死。

我也试过将分配的对象大小设置小于 85kb, .NET Core 3.0 [和.NET](http://xn--0tr.NET) Core2.2 在内存限制小于10mb都可以正常运行，这应该是和 GC 在 Generation 0 中的频繁清理的机制有关，因为清理几乎不消耗时间，不像 Generation 2, 所以在没有限制GC heap的情况也可以运行。

我将上述代码 发布到了 StackOverFlow 和Github 进行提问，

[https://stackoverflow.com/questions/56578084/why-doesnt-heap-memory-used-go-down-after-a-gc-in-clr](https://stackoverflow.com/questions/56578084/why-doesnt-heap-memory-used-go-down-after-a-gc-in-clr)

[https://github.com/dotnet/coreclr/issues/25148](https://github.com/dotnet/coreclr/issues/25148)

有兴趣可以探讨一下。

总结
--

.NET Core 3.0 的改动还是很大滴，以及应该根据自己具体的应用场景去配置GC ，让GC 发挥最好的作用，充分利用Microsoft 给我们的权限。比如启用Server GC 对于高吞吐量的程序有帮助，比如禁用 Concurrent GC 实际上对一个高密度计算的程序是有性能提升的。

参考文章

*   [https://docs.microsoft.com/en-us/dotnet/standard/garbage-collection/fundamentals](https://docs.microsoft.com/en-us/dotnet/standard/garbage-collection/fundamentals)
*   [https://devblogs.microsoft.com/premier-developer/understanding-different-gc-modes-with-concurrency-visualizer/](https://devblogs.microsoft.com/premier-developer/understanding-different-gc-modes-with-concurrency-visualizer/)
*   [https://devblogs.microsoft.com/dotnet/running-with-server-gc-in-a-small-container-scenario-part-1-hard-limit-for-the-gc-heap/](https://devblogs.microsoft.com/dotnet/running-with-server-gc-in-a-small-container-scenario-part-1-hard-limit-for-the-gc-heap/)
*   [https://devblogs.microsoft.com/dotnet/running-with-server-gc-in-a-small-container-scenario-part-0/](https://devblogs.microsoft.com/dotnet/running-with-server-gc-in-a-small-container-scenario-part-0/)
*   [https://devblogs.microsoft.com/dotnet/announcing-net-core-3-preview-3/](https://devblogs.microsoft.com/dotnet/announcing-net-core-3-preview-3/)

更新=

[对于.NET](http://xn--6kqx04a.NET) Core 3.0 GC的变化，我有针对Github 上作者的Merge Request 做出了以下总结：

.NET Core3.0 对GC 改动的 [Merge Request](https://github.com/dotnet/coreclr/pull/22180/commits/fdbbfc9b68a3b57f361eb1709ee2b9277840ff0c#diff-295f0ed467af7d7d972f659a633bf8b9)

代码就不看了，一是看不懂，二是根本没发现对内存的限制，只是添加了获取容器是否设置内存限制的代码，和HeapHardLimit的宏定义，那就意味着，`GCHeadHardLimit` 只是一个阈值而已。由次可见，`GCHeapHardLimit` 属于GC的一个小部件。

其中有一段很重要的总结，[是.NET](http://xn--0iv.NET) Core 3.0 GC的主要变化

        // + we never need to acquire new segments. This simplies the perf
        // calculations by a lot.
        //
        // + we now need a different definition of "end of seg" because we
        // need to make sure the total does not exceed the limit.
        //
        // + if we detect that we exceed the commit limit in the allocator we
        // wouldn't want to treat that as a normal commit failure because that
        // would mean we always do full compacting GCs.
    

*   首先就是，在有内存限制的 Docker 容器中，GC不需要去问虚拟内存要新的`Segments`，因为初始化CLR的时候，把`heap` 和`Segment`都分配好了。在`Server GC` 模式下，一个核心 CPU 对应一个进程，对应一个`heap`， 而一个`segment` 大小 就是 `limit / number of heaps` 。  
    所以程序启动时，如果分配CPU 是一核，那么就会分配一个`heap` ，一个`heap` 中只有一个`segment` ，大小就是 `limit` ,GC 也不会再去问CLR要内存了。请注意这里的 `limit` 和 `GCHeapHardLimit` 不是同一个，这里的`limit` 应该就是容器内存限制。所以GC 堆大小是多少？就是容器的内存限制`limit`
    
*   特殊的判断segment结束标志，以判断是否超过`GCHeapHardLimit`
    
*   如果发现，在 `segment` 中分配内存的时候超出了`GCHeadHardLimit` ，那么不会把这次分配看做失败的，所以就不会发生GC。结合上面两点的铺垫我们可以发现：
    
    1.  首先从上述代码我们可以发现`GCHeapHardLimit` 只是一个数字而已。它就是一个阈值。
        
    2.  其次 GC堆的大小： **请注意，GC堆大小不是 HeapHardLimit 而是 容器内存限制 limit**。GC 分配对象的时候，如果溢出了这个`GCHeapHardLimit`数字，GC 也会睁一只眼闭一只眼，否则只要溢出，它就要去整个`heap`中 GC 一遍。所以 `GCHeadHardLimit` 不是 GC堆申请的`segment`的大小，而是 GC 会管住自己的手脚，不能碰的东西咱尽量不要去碰，要是真碰了，也只有那么一次。
        

如果你的程序使用内存超出了`GCHeapHardLimit`阈值，segment 中还是有空余的，但是 GC 就是不用，它就是等着报`OutOfMemoryException`错误，而且docker根本杀不死你。

但是这并不代表`GCHeapHardLimit`的设置是不合理的，如果你的程序自己不能合理管理对象，或者你太抠门了，那么神仙也乏术。

但是人家说了！`GCHeapHardLimit` 是可以修改的！

     // Users can specify a hard limit for the GC heap via GCHeapHardLimit or
        // a percentage of the physical memory this process is allowed to use via
        // GCHeapHardLimitPercent. This is the maximum commit size the GC heap 
        // can consume.
        //
        // The way the hard limit is decided is:
        // 
        // If the GCHeapHardLimit config is specified that's the value we use;
        // else if the GCHeapHardLimitPercent config is specified we use that 
        // value;
        // else if the process is running inside a container with a memory limit,
        // the hard limit is 
        // max (20mb, 75% of the memory limit on the container).
    

如果你觉得`GCHeapHardLimit` 太气人了,那么就手动修改它的数值吧。




# Garbage Collection



.NET's garbage collector manages the allocation and release of memory for your application. Each time you create a new object, the common language runtime allocates memory for the object from the managed heap. As long as address space is available in the managed heap, the runtime continues to allocate space for new objects. However, memory is not infinite. Eventually the garbage collector must perform a collection in order to free some memory. The garbage collector's optimizing engine determines the best time to perform a collection, based upon the allocations being made. When the garbage collector performs a collection, it checks for objects in the managed heap that are no longer being used by the application and performs the necessary operations to reclaim their memory.

## Related Topics


Title | Description
---|---
[Fundamentals of Garbage Collection](https://docs.microsoft.com/en-us/dotnet/standard/garbage-collection/fundamentals) | Describes how garbage collection works, how objects are allocated on the managed heap, and other core concepts.
[Garbage Collection and Performance](https://docs.microsoft.com/en-us/dotnet/standard/garbage-collection/performance) |Describes the performance checks you can use to diagnose garbage collection and performance issues.
[Induced Collections](https://docs.microsoft.com/en-us/dotnet/standard/garbage-collection/induced)| Describes how to make a garbage collection occur.
[Latency Modes](https://docs.microsoft.com/en-us/dotnet/standard/garbage-collection/latency)|Describes the modes that determine the intrusiveness of garbage collection.
[Optimization for Shared Web Hosting](https://docs.microsoft.com/en-us/dotnet/standard/garbage-collection/optimization-for-shared-web-hosting) |Describes how to optimize garbage collection on servers shared by several small Web sites.
[Garbage Collection Notifications](https://docs.microsoft.com/en-us/dotnet/standard/garbage-collection/notifications) |Describes how to determine when a full garbage collection is approaching and when it has completed.
[Application Domain Resource Monitoring](https://docs.microsoft.com/en-us/dotnet/standard/garbage-collection/app-domain-resource-monitoring) | Describes how to monitor CPU and memory usage by an application domain.
[Weak References](https://docs.microsoft.com/en-us/dotnet/standard/garbage-collection/weak-references)| Describes features that permit the garbage collector to collect an object while still allowing the application to access that object.

## Reference

[System.GC](/en-us/dotnet/api/system.gc)

[System.GCCollectionMode](/en-us/dotnet/api/system.gccollectionmode)

[System.GCNotificationStatus](/en-us/dotnet/api/system.gcnotificationstatus)

[System.Runtime.GCLatencyMode](/en-us/dotnet/api/system.runtime.gclatencymode)

[System.Runtime.GCSettings](/en-us/dotnet/api/system.runtime.gcsettings)

[GCSettings.LargeObjectHeapCompactionMode](/en-us/dotnet/api/system.runtime.gcsettings.largeobjectheapcompactionmode)

[Object.Finalize](/en-us/dotnet/api/system.object.finalize)

[System.IDisposable](/en-us/dotnet/api/system.idisposable)
