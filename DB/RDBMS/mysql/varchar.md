

### MySQL中Varchar(50)和varchar(500)区别是什么？
https://www.toutiao.com/article/7391749141468086821/

[Mysql 建表总长度限制（所有字段）学习总结 以及 所有字段长度计算方式](https://blog.csdn.net/weixin_44131922/article/details/122131467)

当字符集为 UTF-8 时 char 和 varchar 后面的数字 例如 varchar(20) 我们计算时就需要 乘3，varchar(20) 占用字节长度就是 20 * 3 也就是 60字节
当字符集为 GBK 时 char 和 varchar 后面的数字 例如 varchar(20) 我们计算时就需要 乘2，varchar(20) 占用字节长度就是 20 * 2 也就是 40字节

总结:

- Mysql是存在表字段总长度限制的,每行允许最多 65535 字节 长度
- 占用空间一样
- 全表有排序(order by)的情况下, 两种性能差异巨大；
原因: 当我们对该字段进行排序操作的时候，Mysql会根据该字段的设计的长度进行内存预估，如果设计过大的可变长度，会导致内存预估的值超出sort_buffer_size的大小，导致mysql采用磁盘临时文件排序,最终影响查询性能。