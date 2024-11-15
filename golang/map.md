深度解密Go语言之map
https://mp.weixin.qq.com/s?__biz=MjM5MDUwNTQwMQ==&mid=2257483772&idx=1&sn=a6462bc41ec70edf5d60df37a6d4e966&scene=19#wechat_redirect


### SwissTable
在go1.24中SwissTable作为map的默认实现

swisstable的时间复杂度和线性探测法的差不多，空间复杂度介于链表法和线性探测之间。
与原实现相比，Swiss Table在查询、插入和删除操作上均提升了20%至50%的性能，尤其是在处理大hashmap时表现尤为突出；迭代性能提升了10%；内存使用减少了0%至25%，并且不再消耗额外内存。

> https://www.cnblogs.com/apocelipes/p/17562468.html
> https://www.51cto.com/article/801694.html