LRU-K中的K代表最近使用的次数，因此LRU可以认为是LRU-1。LRU-K的主要目的是为了解决LRU算法“缓存污染”的问题，其核心思想是将“最近使用过1次”的判断标准扩展为“最近使用过K次”。

通过两个队列记录数据，一个是新数据队列，另一个是达到K次访问的队列，新数据队列可以通过FIFO方式处理也可以根据LRU方式处理，当到达K次后移动到K次队列。

https://www.iteye.com/blog/flychao88-1977653

https://blog.csdn.net/jake_li/article/details/50659868