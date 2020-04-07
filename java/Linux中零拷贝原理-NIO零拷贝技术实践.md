<!--toc-->
[TOC]
**本文导读：**

- **什么是零拷贝**
    
- **传统 IO 数据拷贝原理**
    
- **什么是 DMA**
    
- **sendfile 数据零拷贝原理**
    
- **mmap 数据零拷贝原理**
    
- **Java 中 NIO 零拷贝实现**
    
- ****Java IO 与 NIO 实战案例分析****
    

**什么是零拷贝**

关于零拷贝，WIKI 上给出的定义如下：

「Zero-copy」 describes computer operations in which the CPU does not perform the task of copying data from one memory area to another. This is frequently used to save CPU cycles and memory bandwidth when transmitting a file over a network.

所谓「零拷贝」描述的是计算机操作系统当中，CPU不执行将数据从一个内存区域，拷贝到另外一个内存区域的任务。通过网络传输文件时，这样通常可以节省 CPU 周期和内存带宽。

从描述中已经了解到**零拷贝技术**给我们带来的好处：

1、节省了 CPU 周期，空出的 CPU 可以完成更多其他的任务

2、减少了内存区域之间数据拷贝，节省内存带宽

3、减少用户态和内核态之间数据拷贝，提升数据传输效率

4、应用零拷贝技术，减少用户态和内核态之间的上下文切换

**传统 IO 数据拷贝原理**

在正式分析零拷贝机制原理之前，我们先来看下传统 IO 在数据拷贝的基本原理，从数据拷贝 (I/O 拷贝) 的次数以及上下文切换的次数进行对比分析。

**传统 IO:**

![640?wx_fmt=png](img/Linux中零拷贝原理-NIO零拷贝技术实践-4.octet-stream)

1、JVM 进程内发起 read() 系统调用，操作系统由用户态空间切换到内核态空间（**第一次上下文切换**）2、通过 DMA 引擎建数据从磁盘拷贝到内核态空间的输入的 socket 缓冲区中（**第一次拷贝**）3、将内核态空间缓冲区的数据原封不动的拷贝到用户态空间的缓存区中（第二次拷贝），同时内核态空间切换到用户态空间（**第二次上下文切换**），read() 系统调用结束4、JVM 进程内业务逻辑代码执行5、JVM 进程内发起 write() 系统调用6、操作系统由用户态空间切换到内核态空间（**第三次上下文切换**），将用户态空间的缓存区数据原封不动的拷贝到内核态空间输出的 socket 缓存区中（**第三次拷贝**）7、write() 系统调用返回，操作系统由内核态空间切换到用户态空间（第四次上下文切换），通过 DMA 引擎将数据从内核态空间的 socket 缓存区数据拷贝到协议引擎中（**第四次拷贝**）

传统 IO 方式，一共在用户态空间与内核态空间之间发生了 4 次上下文的切换，4 次数据的拷贝过程，其中包括 2 次 DMA 拷贝和 2 次 I/O 拷贝（内核态与用户应用程序之间发生的拷贝）。

内核空间缓冲区的一大用处是为了减少磁盘I/O操作，因为它会从磁盘中预读更多的数据到缓冲区中。而使用 BufferedInputStream 的用处是减少 「系统调用」。

**什么是DMA**

**DMA（Direct Memory Access）**—直接内存访问 ：DMA是允许外设组件将 I/O 数据直接传送到主存储器中并且传输不需要 CPU 的参与，以此将 CPU 解放出来去完成其他的事情。

**sendfile 数据零拷贝原理**

**sendfile 数据零拷贝:**显然，在传统 IO 中，用户态空间与内核态空间之间的复制是完全不必要的，因为用户态空间仅仅起到了一种数据转存媒介的作用，除此之外没有做任何事情。Linux 提供了 sendfile() 用来减少我们前面提到的数据拷贝和的上下文切换次数。

**如下图所示：**

![640?wx_fmt=png](img/Linux中零拷贝原理-NIO零拷贝技术实践-5.octet-stream)

1、发起 sendfile() 系统调用，操作系统由用户态空间切换到内核态空间（**第一次上下文切换**）2、通过 DMA 引擎将数据从磁盘拷贝到内核态空间的输入的 socket 缓冲区中（**第一次拷贝**）3、将数据从内核空间拷贝到与之关联的 socket 缓冲区（**第二次拷贝**）4、将 socket 缓冲区的数据拷贝到协议引擎中（**第三次拷贝**）5、sendfile() 系统调用结束，操作系统由用户态空间切换到内核态空间（**第二次上下文切换**）

根据以上过程，一共有 2 次的上下文切换，3 次的 I/O 拷贝。我们看到从用户空间到内核空间并没有出现数据拷贝，从操作系统角度来看，这个就是零拷贝。内核空间出现了复制的原因: 通常的硬件在通过DMA访问时期望的是连续的内存空间。

**支持 scatter-gather 特性的 sendfile 数据零拷贝：**

![640?wx_fmt=png](img/Linux中零拷贝原理-NIO零拷贝技术实践-6.octet-stream)

这次相比 sendfile() 数据零拷贝，减少了一次从内核空间到与之相关的 socket 缓冲区的数据拷贝。

**基本流程:**1、发起 sendfile() 系统调用，操作系统由用户态空间切换到内核态空间（**第一次上下文切换**）

2、通过 DMA 引擎将数据从磁盘拷贝到内核态空间的输入的 socket 缓冲区中（**第一次拷贝**）

3、将描述符信息会拷贝到相应的 socket 缓冲区当中，该描述符包含了两方面的信息： 

**_a)_**kernel buffer的内存地址；

_**b)**_kernel buffer的偏移量。

4、DMA gather copy 根据 socket 缓冲区中描述符提供的位置和偏移量信息直接将内核空间缓冲区中的数据拷贝到协议引擎上(**第二次拷贝**)，这样就避免了最后一次 I/O 数据拷贝。

5、sendfile() 系统调用结束，操作系统由用户态空间切换到内核态空间（**第二次上下文切换**）

**下面这个图更进一步理解：**

![640?wx_fmt=png](img/Linux中零拷贝原理-NIO零拷贝技术实践-7.octet-stream)

Linux/Unix 操作系统下可以通过下面命令查看是否支持 scatter-gather 特性。

```javascript
ethtool -k eth0 | grep scatter-gatherscatter-gather: on
```

许多的 web server 都已经支持了零拷贝技术，比如 Apache、Tomcat。

sendfile 零拷贝消除了所有内核空间缓冲区与用户空间缓冲区之间的数据拷贝过程，因此 sendfile 零拷贝 I/O 的实现是完成在内核空间中完成的，这对于应用程序来说就无法对数据进行操作了。

**mmap 数据零拷贝原理**

如果需要对数据做操作，Linux 提供了mmap 零拷贝来实现。

**mmap 零拷贝：**

![640?wx_fmt=png](img/Linux中零拷贝原理-NIO零拷贝技术实践-8.octet-stream)

通过上图看到，一共发生了 4 次的上下文切换，3 次的 I/O 拷贝，包括 2 次 DMA 拷贝和 1 次的 I/O 拷贝，相比于传统 IO 减少了一次 I/O 拷贝。使用 mmap() 读取文件时，只会发生第一次从磁盘数据拷贝到 OS 文件系统缓冲区的操作。

**1）在什么场景下使用 mmap() 去访问文件会更高效？**

> 对文件执行随机访问时，如果使用 read() 或 write()，则意味着较低的 cache 命中率。这种情况下使用 mmap() 通常将更高效。
> 多个进程同时访问同一个文件时（无论是顺序访问还是随机访问），如果使用mmap()，那么操作系统缓冲区的文件内容可以在多个进程之间共享，从操作系统角度来看，使用 mmap() 可以大大节省内存。

**2）什么场景下没有使用 mmap() 的必要？**

> 访问小文件时，直接使用 read() 或 write() 将更加高效。
> 单个进程对文件执行顺序访问时 (sequential access)，使用 mmap() 几乎不会带来性能上的提升。譬如说，使用 read() 顺序读取文件时，文件系统会使用 read-ahead 的方式提前将文件内容缓存到文件系统的缓冲区，因此使用 read() 将很大程度上可以命中缓存。

**Java 中 NIO 零拷贝实现**

Java NIO 中的通道（Channel）相当于操作系统的内核空间（kernel space）的缓冲区，而缓冲区（Buffer）对应的相当于操作系统的用户空间（user space）中的用户缓冲区（user buffer）。

- **通道（Channel）**是全双工的（双向传输），它既可能是读缓冲区（read buffer），也可能是网络缓冲区（socket buffer）。
    
- **缓冲区（Buffer）**分为堆内存（HeapBuffer）和堆外内存（DirectBuffer），这是通过 malloc() 分配出来的用户态内存。
    

Java NIO 引入了用于通道的缓冲区的 ByteBuffer。 

**ByteBuffer有三个主要的实现：**

**1、HeapByteBuffer**

调用 ByteBuffer.allocate() 方法时使用到 HeapByteBuffer。这个缓存区域是在 JVM 进程的堆上分配的，可以获得如GC支持和缓存优化的优势。 

但它不是页面对齐的，这意味着若需通过JNI与本地代码交谈，JVM将不得不复制到对齐的缓冲区空间。

**2、DirectByteBuffer**

调用 ByteBuffer.allocateDirect() 方法时使用。 JVM 会使用 malloc() 在堆空间之外分配内存空间。 由于它的内存空间不由 JVM 管理，所以你的内存空间是页面对齐的，不受GC影响。但需要自己管理这个内存，注意分配和释放内存来防止内存泄漏。

**3、MappedByteBuffer**

调用 FileChannel.map() 时使用。与DirectByteBuffer类似，这也是 JVM 堆外部分配内存空间。它基本上作为操作系统 mmap() 系统调用的包装函数，以便代码直接操作映射的物理内存数据。

**Java IO 与 NIO 实战案例分析**

下面我们通过代码示例来对比下传统 IO 与使用了零拷贝技术的 NIO 之间的差异。我们通过服务端开启 socket 监听，然后客户端连接的服务端进行数据的传输，数据传输文件大小为 237M。

零拷贝技术的 NIO，这里咱们通过刚刚介绍的 HeapByteBuffer 来实战对比一下。

1、构建传统IO的socket服务端，监听8898端口。

```javascript
public class OldIOServer {	
    public static void main(String[] args) throws Exception {	
         try (ServerSocket serverSocket = new ServerSocket(8898)) {	
             while (true) {	
                 Socket socket = serverSocket.accept();	
                 DataInputStream inputStream = new DataInputStream(socket.getInputStream());	

	
                 byte[] bytes = new byte[4096];	
                  // 从socket中读取字节数据	
                 while (true) {	
                     // 读取的字节数大小，-1则表示数据已被读完	
                     int readCount = inputStream.read(bytes, 0, bytes.length);	
                     if (-1 == readCount) {	
                       break;	
                     }	
                  }	
              }	
          }	
     }	
}
```

2、构建传统 IO 的客户端，连接服务端的 8898 端口，并从磁盘读取 237M 的数据文件向服务端 socket 中发起写请求。

```javascript
public class OldIOClient {	

	
    public static void main(String[] args) throws Exception {	
        Socket socket = new Socket();	
        socket.connect(new InetSocketAddress("localhost", 8898)); // 连接服务端socket 8899端口	

	
        // 设置一个大的文件, 237M	
        try (FileInputStream fileInputStream = new FileInputStream(new File("/Users/david/Downloads/jdk-8u144-macosx-x64.dmg"));	
             // 定义一个输出流	
             DataOutputStream dataOutputStream = new DataOutputStream(socket.getOutputStream());) {	

	
            // 读取文件数据	
            // 定义byte缓存	
            byte[] buffer = new byte[4096];	
            int readCount; // 每一次读取的字节数	
            int total = 0; // 读取的总字节数	

	
            long startTime = System.currentTimeMillis();	

	
            while ((readCount = fileInputStream.read(buffer)) > 0) {	
                total += readCount; //累加字节数	
                dataOutputStream.write(buffer); // 写入到输出流中	
            }	

	
            System.out.println("发送的总字节数：" + total + ", 耗时：" + (System.currentTimeMillis() - startTime));	
        }	
    }	
}	

```

运行结果：发送的总字节数：237607747，耗时：450 （400~600毫秒之间）接下来，我们通过使用 JDK 提供的 NIO 的方式实现数据传输与上述传统 IO 做对比。

1、构建基于 NIO 的服务端，监听 8899 端口。

```javascript
public class NewIOServer {	
public static void main(String[] args) throws Exception {	
       ServerSocketChannel serverSocketChannel = ServerSocketChannel.open();	
       serverSocketChannel.bind(new InetSocketAddress(8899));	

	
       ByteBuffer byteBuffer = ByteBuffer.allocate(4096);	

	
       while (true) {	
           SocketChannel socketChannel = serverSocketChannel.accept();	
           socketChannel.configureBlocking(false); // 这里设置为阻塞模式	
           int readCount = socketChannel.read(byteBuffer);	

	
            while (-1 != readCount) {	
               readCount = socketChannel.read(byteBuffer);	

	
               // 这里一定要调用下rewind方法，将position重置为0开始位置	
               byteBuffer.rewind();	
           }	
       }	
   }	
}
```

2、构建基于 NIO 的客户端，连接NIO的服务端 8899 端口，通过 

FileChannel.transferTo  传输 237M 的数据文件。

```javascript
public class NewIOClient {	
public static void main(String[] args) throws Exception {	
       SocketChannel socketChannel = SocketChannel.open();	
       socketChannel.connect(new InetSocketAddress("localhost", 8899));	
       socketChannel.configureBlocking(true);	

	
       String fileName = "/Users/david/Downloads/jdk-8u144-macosx-x64.dmg";	

	
       FileInputStream fileInputStream = new FileInputStream(fileName);	
       FileChannel fileChannel = fileInputStream.getChannel();	

	
       long startTime = System.currentTimeMillis();	
       long transferCount = fileChannel.transferTo(0, fileChannel.size(), socketChannel); // 目标channel	
       System.out.println("发送的总字节数：" + transferCount + ",耗时：" + (System.currentTimeMillis() - startTime));	
       fileChannel.close();	
   }	
}
```

运行结果：发送的总字节数：237607747,耗时：161（100到300毫秒之间）结合运行结果，基于 NIO 零拷贝技术要比传统 IO 传输效率高 3倍多。所以，后续当设计大文件数据传输时可以优先采用类似 NIO 的方式实现。

这里我们使用了 FileChannel，其中调用的 transferTo() 方法将数据从 FileChannel传输到其他的 channel 中，如果操作系统底层支持的话 transferTo、transferFrom 会使用相关的零拷贝技术来实现数据的传输。所以，这里是否使用零拷贝必须依赖于底层的系统实现。

**FileChannel.transferTo 方法：**

```javascript
public abstract long transferTo(long position,	
long count,	
WritableByteChannel target) throws IOException
```

将字节从此通道的文件传输到给定的可写入字节通道。

试图读取从此通道的文件中给定 position 处开始的 count 个字节，并将其写入目标通道。

此方法的调用不一定传输所有请求的字节；

是否传输取决于通道的性质和状态。

如果此通道的文件从给定的 position 处开始所包含的字节数小于 count 个字节，或者如果目标通道是非阻塞的并且其输出缓冲区中的自由空间少于 count 个字节，则所传输的字节数要小于请求的字节数。

此方法不修改此通道的位置。

如果给定的位置大于该文件的当前大小，则不传输任何字节。

如果目标通道中有该位置，则从该位置开始写入各字节，然后将该位置增加写入的字节数。

与从此通道读取并将内容写入目标通道的简单循环语句相比，此方法可能高效得多。

很多操作系统可将字节直接从文件系统缓存传输到目标通道，而无需实际复制各字节。

参数：

position - 文件中的位置，从此位置开始传输；

必须为非负数

count - 要传输的最大字节数；

必须为非负数

target - 目标通道

返回：实际已传输的字节数，可能为零

**FileChannel.transferFrom 方法：**

```javascript
public abstract long transferFrom(ReadableByteChannel src,	
long position,	
long count) throws IOException
```

将字节从给定的可读取字节通道传输到此通道的文件中。

试着从源通道中最多读取 count 个字节，并将其写入到此通道的文件中从给定 position 处开始的位置。

此方法的调用不一定传输所有请求的字节；

是否传输取决于通道的性质和状态。

如果源通道的剩余空间小于 count 个字节，或者如果源通道是非阻塞的并且其输入缓冲区中直接可用的空间小于 count 个字节，则所传输的字节数要小于请求的字节数。

此方法不修改此通道的位置。

如果给定的位置大于该文件的当前大小，则不传输任何字节。

如果该位置在源通道中，则从该位置开始读取各字节，然后将该位置增加读取的字节数。

与从源通道读

取并将内容写入此通道的简单循环语句相比，此方法可能高效得多。

很多操作系统可将字节直接从源通道传输到文件系统缓存，而无需实际复制各字节。

参数：

src - 源通道

position - 文件中的位置，从此位置开始传输；

必须为非负数

count - 要传输的最大字节数；

必须为非负数

返回：实际已传输的字节数，可能为零

**发生相应的异常的情况：**

异常抛出：

IllegalArgumentException - 如果关于参数的前提不成立
NonReadableChannelException - 如果不允许从此通道进行读取操作
NonWritableChannelException - 如果目标通道不允许进行写入操作
ClosedChannelException - 如果此通道或目标通道已关闭
AsynchronousCloseException - 如果正在进行传输时另一个线程关闭了任一通道
ClosedByInterruptException - 如果正在进行传输时另一个线程中断了当前线程，因此关闭了两个通道并将当前线程设置为中断
IOException - 如果发生其他 I/O 错误

**参考资料：**http://xcorpion.tech/2016/09/10/It-s-all-about-buffers-zero-copy-mmap-and-Java-NIO/http://www.jianshu.com/p/e76e3580e356http://www.linuxjournal.com/node/6345http://senlinzhan.github.io/2017/03/25/%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B%E4%B8%AD%E7%9A%84zerocpoy%E6%8A%80%E6%9C%AF/JDK 官方文档


[原文](https://blog.csdn.net/david_lua/article/details/102735550)