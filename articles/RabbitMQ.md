https://www.cnblogs.com/haolujun/p/9632835.html
https://mp.weixin.qq.com/s?__biz=MzAwNTMxMzg1MA==&mid=2654076958&idx=7&sn=1237f56165fbdc2751629af1049a64c8
对rabbitmq来说，产生死信的来源大致有如下几种：
- 消息被拒绝（basic.reject或basic.nack）并且requeue=false.
- 消息TTL过期
- 队列达到最大长度（队列满了，无法再添加数据到mq中）

[新来个技术总监，把 RabbitMQ 讲的那叫一个透彻，佩服！](https://blog.csdn.net/qq_46388795/article/details/125633587)