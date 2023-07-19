# Bcache
Bcache是Linux内核块设备层cache，支持多块HDD使用同一块SSD作为缓存盘。它让SSD作为HDD的缓存成为了可能。
Bcache是从Linux-3.10开始正式并入内核主线的，因此，要使用Bcache，需要将内核升级到3.10及以上版本才行。

[官方文档](https://www.kernel.org/doc/html/latest/admin-guide/bcache.html)

## Bcache缓存策略
Bcache支持三种缓存策略，分别是：writeback、writethrough、writearoud，默认使用writethrough，缓存策略可动态修改。

writeback 回写策略：回写策略默认是关闭的，如果开启此策略，则所有的数据将先写入缓存盘，然后等待系统将数据回写入后端数据盘中。
writethrough 写通策略：默认的就是写通策略，此模式下，数据将会同时写入缓存盘和后端数据盘。
writearoud ：选择此策略，数据将直接写入后端磁盘。

> 修改缓存策略`echo writeback > /sys/block/bcache0/bcache/cache_mode`

## bcache-tools的安装

要使用Bcache，必须安装bcache-tools工具包，由于CentOS 7的源中没有bcache-tools，因此，需要手动下载源码包进行编译。源码在这：[bcache-tools](https://github.com/g2p/bcache-tools.git)，下载之后，需要安装libblkid-devel依赖包方可进行编译，通过以下命令即可安装：

`yum install -y libblkid-devel`
`apt-get install -y pkg-config libblkid-dev git gcc`
安装libblkid-devel包成功之后，直接编译bcache-tools安装即可。
```
# git clone https://evilpiepirate.org/git/bcache-tools.git
# cd bcache-tools
# make
# make install
```

## 创建缓存盘

`make-bcache -C /dev/nvme0n1 -B /dev/sdc`

bcache-tools新版本:https://git.kernel.org/pub/scm/linux/kernel/git/colyli/bcache-tools.git/
bcache-tools已经合入最新的linux内核版本了, 3.11以上版本
`bcache make -B /dev/sda /dev/sdb -C /dev/sdc`


-B是后端盘, -C前端盘也就是缓存盘
## 测试
```
# 随机小文件写
sudo fio -ioengine=libaio -name=test2 -filename=/dev/bcache1  -direct=1 -bs=4KB -iodepth=128 -rw=randwrite -runtime=300

# 顺序大文件写
sudo fio -ioengine=libaio -name=test2 -filename=/dev/bcache1  -direct=1 -bs=128KB -iodepth=128 -rw=write -runtime=300
```

默认情况下, bcache不会缓存顺序IO和大文件, 所以顺序写性能应该没什么提升. 可以打开顺序IO缓存 
```
[root@localhost opt]# echo 0 > /sys/block/bcache0/bcache/sequential_cutoff
```


## 参考
[Linux下SSD缓存加速之bcache使用](https://www.cnblogs.com/zhangxinglong/p/14247606.html)
