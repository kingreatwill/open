https://github.com/antirez/redis-doc

### 六种概率性结构
https://redis.ac.cn/docs/latest/develop/data-types/probabilistic/top-k/
- HyperLogLog 用于估计集合的个数
- Bloom filter 用于检查元素是否存在于集合中
- cuckoo filter 用于检查元素是否存在于集合中
- count-min sketch 用于估计集合中元素的频率
- top-k 用于估计集合中排名最高的 K 个元素
- t-digest 用于估计集合的百分位数