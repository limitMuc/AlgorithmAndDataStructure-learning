计算机存储器空间的大小固定，无法容纳服务器上所有的文件，所以当有新的文件要被置换入缓存时，必须根据一定的原则来取代掉适当的文件。此原则即所谓缓存文件置换机制。

缓存文件置换方法有：

* 先进先出算法（FIFO）：最先进入的内容作为替换对象
* 最近最少使用算法（LFU）：最近最少使用的内容作为替换对象
* 最久未使用算法（LRU）：最久没有访问的内容作为替换对象
* 非最近使用算法（NMRU）：在最近没有使用的内容中随机选择一个作为替换对象

<br/>

![](https://img-blog.csdnimg.cn/20200302220506244.png)

<br/>

由于双向链表可以方便的进行节点的删除和插入，所以采用双向链表来实现LRU算法

置换最近一段时间以来最长时间未访问过的页面。根据程序局部性原理，刚被访问的页面，可能马上又要被访问；而较长时间内没有被访问的页面，可能最近不会被访问。传统意义的LRU算法是为每一个Cache对象设置一个计数器，每次Cache命中则给计数器+1，而Cache用完，需要淘汰旧内容，放置新内容时，就查看所有的计数器，并将最少使用的内容替换掉。它的弊端很明显，如果Cache的数量少，问题不会很大， 但是如果Cache的空间过大，达到10W或者100W以上，一旦需要淘汰，则需要遍历所有计算器，其性能与资源消耗是巨大的，效率也就非常的慢了。

LRU算法普偏地适用于各种类型的程序，但是系统要时时刻刻对各页的访问历史情况加以记录和更新，开销太大，因此LRU算法必须要有硬件的支持。这种算法存在着问题：可能由于一次冷数据的批量查询而误淘汰大量热点的数据

<br/>

### 参考
[https://studygolang.com/articles/23183?fr=sidebar](https://studygolang.com/articles/23183?fr=sidebar)