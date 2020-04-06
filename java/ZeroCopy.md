# 零拷贝(zero copy)
## Linux系统内存管理知识补充
Linux系统是虚拟内存系统，虚拟内存并不是真正的物理内存，而是虚拟的连续内存地址空间。虚拟内存又分为内核空间和用户空间，内核空间是内核程序运行的地方，用户空间是用户进程代码运行的地方，只有内核才能直接访问物理内存并为用户空间映射物理内存(MMU)。内核会为每个进程分配独立的连续的虚拟内存空间，并且在需要的时候映射物理内存，为了完成内存映射，内核为每个进程都维护了一张页表，记录虚拟地址与物理地址的映射关系，这个页表就是存在于MMU中；用户进程访问内存的时候，通过页表把虚拟内存地址转换为物理内存地址进而访问数据；其实对于用户进程而言，虚拟内存就是内存一般的存在(当作内存看待就好)。这样的设计可以把用户程序和系统程序分开，互不影响；内核可以对所有的用户程序进行管理，比如限制内存滥用等。

虚拟内存的最小单位是页，通常是4KB大小，所以虚拟内存会有很多很多的页组成，当然也有大页，顾名思义就是大的虚拟内存空间，比如12KB，2MB。虚拟内存和物理内存的映射都是等空间的，映射的物理内存是多大的，那么占用的虚拟内存差不多也是多大，都是4KB的整数倍。比如映射了一个1KB的内存空间，那么也是占用一页4KB虚拟内存。

用户进程在处于用户态时，只能访问用户空间；只有进入内核态后，才可以访问内核空间。虽然每个进程的地址空间都包含了内核空间，但这些内核空间映射的物理内存都是相同的，所以当进程切换到内核态后可以快速的访问内核空间数据。

内核其实就是一段特殊的代码程序，运行于内核空间，控制着计算机的CPU、IO、内存等，提供了一系列的系统接口供外部调用，通常叫做系统调用。只有线程或者进程处于内核态的时候才能进行系统调用，如果处于用户态的话，是需要转换为内核态才能访问。其实就是权限不同，内核态(Ring0)拥有比用户态(Ring3)更高的权限，拥有着访问系统硬件资源的权限。

一般用户线程或者进程是不需要切换到内核态运行的，除非：

1. 系统调用，其实系统调用本身就是中断，但是软件中断，跟硬中断不同。

2. 异常：如果当前进程运行在用户态，如果这个时候发生了异常事件，就会触发切换。
例如：缺页异常。

3. 外设中断：当外设完成用户的请求时，会向CPU发送中断信号。

比如读取硬盘数据，除了IO属于系统操作需要切换为内核态来获取权限的原因外还要一原因是：

为了减少磁盘的IO操作，为了提高性能而考虑的，因为我们的程序访问一般都带有局部性，也就是所谓的局部性原理，即我们访问了文件的某一段数据，那么接下去很可能还会访问接下去的一段数据，由于磁盘IO操作的速度比直接访问内存慢了好几个数量级，所以OS根据局部性原理会在一次 read()系统调用过程中预读更多的文件数据缓存在内核IO缓冲区中，当继续访问的文件数据在缓冲区中时便直接拷贝数据到进程私有空间，避免了再次的低效率磁盘IO操作。

## 传统IO发送文件

1. 用户程序调用read，进入内核态，上下文切换由用户空间切换为内核空间，由DMA(Direct Memory Access)加载文件数据到内核空间。

2. CPU把数据从内核空间复制到用户空间，转换为用户态，上下文由内核空间切换为用户空间。

3. 用户程序调用write，再次进入内核态，CPU把数据从用户空间复制到socket关联的内核空间。

4. 最后通过DMA 将内核模式下的socket缓冲区中的数据复制到网卡设备中传送，进而返回用户空间进入用户态。

## sendfile零拷贝(<Linux 2.4)

1. 用户程序调用read，进入内核态，上下文切换由用户空间切换为内核空间，由DMA(Direct Memory Access)加载文件数据到内核空间，第一步和传统IO相同。

2. 在内核态下，CPU把数据从内核空间复制到socket关联的内核空间。

3. 最后通过DMA 将内核模式下的socket缓冲区中的数据复制到网卡设备中传送，进而返回用户空间进入用户态，最后一步也是和传统IO相同。

与传统IO相比，缺少了把数据从内核空间复制到用户空间，再由用户空间复制到内核空间，比原来缺少了一次CPU复制(复制3次，CPU参与复制一次)，少了两次上下文切换(两次)。

**从内核空间角度来看，其实已经是“ZERO COPY”了，因为没有往用户空间复制的操作。**

## sendfile零拷贝(>=Linux 2.4)

1. 用户程序调用read，进入内核态，上下文切换由用户空间切换为内核空间，由DMA(Direct Memory Access)加载文件数据到内核空间，第一步和传统IO相同。

2. 在内核态下，描述符(包含了数据的位置和长度等信息)追加到socket关联的缓冲区中，并没有进行数据的拷贝。

3. 最后DMA根据提供的位置和偏移量信息直接将内核空间缓冲区中的数据拷贝到协议引擎上进而返回用户空间进入用户态。

**这次优化点在于没有CPU参与复制，两次DMA数据复制，不过还是两次上下文切换。**

## 通过mmap实现的零拷贝(常用来处理大文件)

当进行mmap系统调用的时候，将文件的内容的全部或一部分直接映射到进程的地址空间(虚拟内存)，映射完成后，进程可以像访问普通内存一样做其他的操作，mmap并不分配物理地址空间，它只是占有进程的虚拟地址空间。

当进程访问内核中的缓冲区时候，并没有实际拷贝数据，这时MMU在地址映射表中是无法找到与ptr相对应的物理地址的，也就是MMU失败，就会触发缺页中断。内核将文件的这一页数据读入到内核高速缓冲区中，并更新用户进程的页表，使页表指向内核缓冲中的这一页，实现了用户空间和内核空间数据的直接交换，可以看待为内核空间和用户空间共享的一段物理内存。

## Java调用零拷贝
```java
FileInputStream input = new FileInputStream("1.txt");
FileChannel channel = input.getChannel();
FileOutputStream out = new FileOutputStream("2.txt");
channel.transferTo(0, channel.size(), out.getChannel());
```
上面这种方式其实调用的是Linux系统的sendfile系统指令，无论什么语言代码实现的零拷贝其实调用的都是操作系统本身提供的系统指令，只是做了封装而已。
```java
FileInputStream input = new FileInputStream("1.txt");
FileChannel channel = input.getChannel();
MappedByteBuffer mappedBuffer = channel.map(MapMode.READ_ONLY, 0, channel.size());
System.out.println(Charset.forName("utf-8").decode(mappedBuffer).toString());
```
上面这种方式其实调用的是Linux系统的mmap系统指令；在读取大文件的时候用这种方法映射大文件的一部分到内存空间，比较方便快捷。
```java
//mmap写数据
Instant now = Instant.now();
RandomAccessFile outFile = new RandomAccessFile("1.txt","rw");  
FileChannel channel = outFile.getChannel();
long size = 1024*1024*60;
MappedByteBuffer mappedBuffer = channel.map(MapMode.READ_WRITE, 0, size);
for(int i=0;i<1000000;i++) {
    mappedBuffer.put("11111111111111111111111111111111111111111111111111111111111\n".getBytes());
}
System.out.println(ChronoUnit.MILLIS.between(now, Instant.now())); // 118
//fileOutputStream写数据
Instant nowStream = Instant.now();
FileOutputStream outStream = new FileOutputStream("2.txt");
for(int i=0;i<1000000;i++) {
    outStream.write("11111111111111111111111111111111111111111111111111111111111\n".getBytes());
}
System.out.println(ChronoUnit.MILLIS.between(nowStream, Instant.now())); // 9130
```
通过上面的测试可以看出在频繁的写入文件操作上mmap占有很多大的优势，数量级的优势。但是把上例的循环次数改为50的话，mmap就不占优势了，因为在映射的时候需要新开辟内存空间，这个耗时相对于极少量的写操作而言显得占比重就大了。


# 压缩20M文件从30秒到1秒的优化过程
https://www.jianshu.com/p/25b328753017
https://github.com/modouxiansheng/Doraemon