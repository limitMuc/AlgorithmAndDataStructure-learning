进程运行过程中，如果发生缺页中断，而此时内存中有没有空闲的物理块是，为了能够把所缺的页面装入内存，系统必须从内存中选择一页调出到磁盘的对换区。计算机存储器空间的大小固定，无法容纳服务器上所有的文件，所以当有新的文件要被置换入缓存时，必须根据一定的原则来取代掉适当的文件。此原则即所谓缓存文件置换机制。

缓存文件置换方法有：

* 先进先出算法（FIFO）：最先进入的内容作为替换对象
* 最近最少使用算法（LFU）：最近最少使用的内容作为替换对象
* 最久未使用算法（LRU）：最久没有访问的内容作为替换对象
* 非最近使用算法（NMRU）：在最近没有使用的内容中随机选择一个作为替换对象

<br/>

### 缺页中断

在请求分页系统中，可以通过查询页表中的状态位来确定所要访问的页面是否存在于内存中。每当所要访问的页面不在内存时，会产生一次缺页中断，此时操作系统会根据页表中的外存地址在外存中找到所缺的一页，将其调入内存。

缺页本身是一种中断，与一般的中断一样，需要经过4个处理步骤：

1. 保护CPU现场
2. 分析中断原因
3. 转入缺页中断处理程序进行处理
4. 恢复CPU现场，继续执行

但是缺页中断时由于所要访问的页面不存在与内存时，有硬件所产生的一种特殊的中断，因此，与一般的中断存在区别：

1. 在指令执行期间产生和处理缺页中断信号
2. 一条指令在执行期间，可能产生多次缺页中断
3. 缺页中断返回时，执行产生中断的那一条指令，而一般的中断返回时，执行下一条指令

<br/>

### LRU算法

![](https://img-blog.csdnimg.cn/20200302220506244.png)

<br/>

实现： ```hash map（哈希表） + doubly linked list（双向链表）```

由于双向链表可以方便的进行节点的删除和插入，所以采用双向链表来实现LRU算法；哈希表的作用是什么呢？如果没有哈希表，我们要访问某个结点，就需要顺序地一个个找， 时间复杂度是O(n)，使用哈希表可以让我们在O(1)的时间快速找到想要访问的结点，所以在插入/删除数据和访问数据的时候都能达到O(1)的时间复杂度。

置换最近一段时间以来最长时间未访问过的页面。根据程序局部性原理，刚被访问的页面，可能马上又要被访问；而较长时间内没有被访问的页面，可能最近不会被访问。传统意义的LRU算法是为每一个Cache对象设置一个计数器，每次Cache命中则给计数器+1，而Cache用完，需要淘汰旧内容，放置新内容时，就查看所有的计数器，并将最少使用的内容替换掉。它的弊端很明显，如果Cache的数量少，问题不会很大， 但是如果Cache的空间过大，达到10W或者100W以上，一旦需要淘汰，则需要遍历所有计算器，其性能与资源消耗是巨大的，效率也就非常的慢了。

LRU算法普偏地适用于各种类型的程序，但是系统要时时刻刻对各页的访问历史情况加以记录和更新，开销太大，因此LRU算法必须要有硬件的支持。这种算法存在着问题：可能由于一次冷数据的批量查询而误淘汰大量热点的数据。

<br/>

### 参考
[https://studygolang.com/articles/23183?fr=sidebar](https://studygolang.com/articles/23183?fr=sidebar)

[https://blog.csdn.net/u011080472/article/details/51206332](https://blog.csdn.net/u011080472/article/details/51206332)

[https://blog.csdn.net/liuyonglun/article/details/103772802](https://blog.csdn.net/liuyonglun/article/details/103772802)

[https://www.cnblogs.com/mafeng/p/7346711.html](https://blog.csdn.net/u011080472/article/details/51206332)

<br/>

### 扩展
[https://www.cnblogs.com/mafeng/p/7346687.html](https://www.cnblogs.com/mafeng/p/7346687.html)