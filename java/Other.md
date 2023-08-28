[万字长文：细谈Linux、JDK、Netty中NIO与零拷贝](https://www.toutiao.com/article/6939752055871848973/)

[VM 内存布局详解，图文并茂，写得太好了！](https://mp.weixin.qq.com/s/m9833QYw8Yb_URUx97Kz4Q)

[DynamicTp动态线程池](https://mp.weixin.qq.com/s/OUD5Orj1lLTaNaOAVfaOIw)

[Mybatis Plus 多租](https://github.com/osinn/druid-multi-tenant-starter)

### 为什么jvm启动后一段时间内内存占用越来越多
在gc时，由于会复制存活对象到堆的空闲部分，如果正好复制到了以前未使用过的区域，就又会触发Linux进行内存分配，故一段时间内内存占用会越来越多，直到堆的所有区域都被touch到。

而通过添加JVM参数`-XX:+AlwaysPreTouch`，可以让JVM为堆申请虚拟内存后，立即把堆全部touch一遍，使得堆区域全都被分配物理内存，而由于Java进程主要活动在堆内，故后续内存就不会有很大变化了

此参数则会使应用运行得更稳定。

> 参考: [一次Java内存占用高的排查案例，解释了我对内存问题的所有疑问](https://www.cnblogs.com/codelogs/p/17659370.html)