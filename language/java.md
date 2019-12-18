# OOM
一、确认是不是内存本身就分配过小
```
jmap -heap 10765
```
二、找到最耗内存的对象
```
jmap -histo:live 10765 | more
```
三、确认是否是资源耗尽
工具：
- pstree
- netstat
查看进程创建的线程数，以及网络连接数，如果资源耗尽，也可能出现OOM。

这里介绍另一种方法，通过
/proc/${PID}/fd

/proc/${PID}/task
可以分别查看句柄详情和线程数。

只要
ll /proc/${PID}/fd | wc -l

ll /proc/${PID}/task | wc -l （效果等同pstree -p | wc -l）
就能知道进程打开的句柄数和线程数。


分布式事务
https://mp.weixin.qq.com/s?__biz=MzA3ODQ0Mzg2OA==&mid=2649050820&idx=1&sn=dd307577d387f0393a5c43e955ec81d0
seata


多module（非java9 的module）
https://segmentfault.com/a/1190000011367492?utm_source=tag-newest