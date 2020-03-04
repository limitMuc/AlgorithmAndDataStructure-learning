package doubleLinkedList

import (
	"fmt"
	"sync"
)

type DNode struct {
	Data interface{}
	Prev *DNode
	Next *DNode
	Freq int // 该节点的访问频次, 非必要选项， 主要用于方便实现其他算法，如：LFU算法
}

func (dNode *DNode) GetData() interface{} {
	return dNode.Data
}

func (dNode *DNode) GetNext() *DNode {
	return dNode.Next
}

func (dNode *DNode) GetPrev() *DNode {
	return dNode.Prev
}

type DList struct {
	Capacity uint // 当前缓存容纳的容量( 0 指容量不作限制), 非必要选项， 主要用于方便实现其他算法，如：FIFO、LRU、LFU算法
	Size     uint
	Head     *DNode
	Tail     *DNode
	Mutex    *sync.RWMutex // 为了多协程读写安全，所以在链表中加了读写锁
}

// 初始化双向链表
func Init(capacity uint) *DList {
	return &DList{
		Capacity: capacity,
		Mutex:    new(sync.RWMutex),
	}
}

// 获取双向链表的长度
func (dList *DList) GetSize() uint {
	return dList.Size
}

// 获取双向链表的头部节点
func (dList *DList) GetHeadDNode() *DNode {
	return dList.Head
}

// 获取双向链表的尾部节点
func (dList *DList) GetTailDNode() *DNode {
	return dList.Tail
}

// 该节点是否是头部节点
func (dList *DList) isHead(dNode *DNode) bool {
	return dList.GetHeadDNode() == dNode
}

// 该节点是否是尾部节点
func (dList *DList) isTail(dNode *DNode) bool {
	return dList.GetTailDNode() == dNode
}

// 根据索引获取节点
func (dList *DList) GetDNodeByIndex(index uint) *DNode {
	if dList.Size == 0 || index > dList.Size-1 {
		return nil
	}
	if index == 0 {
		return dList.Head
	}
	if index == dList.Size-1 {
		return dList.Tail
	}

	var i uint
	dNode := dList.Head
	for i = 1; i <= index; i++ {
		dNode = dNode.Next
	}

	return dNode
}

// 从左边压入一个节点， 即从头部插入一个节点
func (dList *DList) HeadPush(dNode *DNode) *DNode {
	return dList.Insert4Head(dNode)
}

// 从左边弹出一个节点， 即从头部删除一个节点
func (dList *DList) HeadPop() *DNode {
	return dList.Delete4Head()
}

// 从右边压入一个节点， 即从尾部插入一个节点
func (dList *DList) TailPush(dNode *DNode) *DNode {
	return dList.Insert4Tail(dNode)
}

// 从右边弹出一个节点， 即从尾部删除一个节点
func (dList *DList) TailPop() *DNode {
	return dList.Delete4Tail()
}

// 从头部加入一个节点
func (dList *DList) Add2Head(dNode *DNode) *DNode {
	return dList.Insert4Head(dNode)
}

// 从尾部追加一个节点
func (dList *DList) Append(dNode *DNode) *DNode {
	return dList.Insert4Tail(dNode)
}

// 打印整个双向链表
func (dList *DList) PrintDoubleLinkedList() {
	if dList == nil || dList.Size == 0 {
		fmt.Println("this double linked list is nil or empty")
		return
	}

	dList.Mutex.RLock()
	defer dList.Mutex.RUnlock()

	dNode := dList.Head
	fmt.Printf("Size: %v   [", dList.Size)
	for dNode != nil {
		fmt.Printf(" %v", dNode.Data)
		dNode = dNode.Next
		if dNode == nil {
			fmt.Println(" ]")
		} else {
			fmt.Printf(" <=>")
		}
	}

}

// 头插法， 从头部插入
func (dList *DList) Insert4Head(dNode *DNode) *DNode {
	if dNode == nil {
		return nil
	}

	dList.Mutex.Lock()
	defer dList.Mutex.Unlock()

	// 空双向链表的时候
	if dList.GetSize() == 0 {
		dList.Head = dNode
		dList.Tail = dNode
		dNode.Prev = nil // 为防止出现插入的节点不是单一的节点，否则 dList.Head = dNode 将不能准确指向头部
		dNode.Next = nil // 为防止出现插入的节点不是单一的节点，否则 dList.Tail = dNode 将不能准确指向尾部
	} else { // 非空双向链表的时候
		dNode.Next = dList.Head
		dList.Head.Prev = dNode
		dList.Head = dNode
		dNode.Prev = nil // 为防止出现插入的节点不是单一的节点，否则 dList.Head = dNode 将不能准确指向头部
	}

	dList.Size++
	return dNode
}

// 尾插法， 从尾部插入
func (dList *DList) Insert4Tail(dNode *DNode) *DNode {
	if dNode == nil {
		return nil
	}

	dList.Mutex.Lock()
	defer dList.Mutex.Unlock()

	// 空双向链表的时候
	if dList.GetSize() == 0 {
		dList.Head = dNode
		dList.Tail = dNode
		dNode.Prev = nil // 为防止出现插入的节点不是单一的节点，否则 dList.Head = dNode 将不能准确指向头部
		dNode.Next = nil // 为防止出现插入的节点不是单一的节点，否则 dList.Tail = dNode 将不能准确指向尾部
	} else { // 非空双向链表的时候
		dList.Tail.Next = dNode
		dNode.Prev = dList.Tail
		dList.Tail = dNode
		dNode.Next = nil // 为防止出现插入的节点不是单一的节点，否则 dList.Tail = dNode 将不能准确指向尾部
	}

	dList.Size++
	return dNode
}

// 向双链表指定位置插入节点
func (dList *DList) Insert4Index(index uint, dNode *DNode) *DNode {
	if index > dList.Size || dNode == nil {
		return nil
	}

	if index == 0 {
		return dList.Insert4Head(dNode)
	}
	if index == dList.Size {
		return dList.Insert4Tail(dNode)
	}

	nextDNode := dList.GetDNodeByIndex(index)

	dList.Mutex.Lock()
	defer dList.Mutex.Unlock()

	dNode.Next = nextDNode
	dNode.Prev = nextDNode.Prev
	nextDNode.Prev = dNode
	nextDNode.Prev.Next = dNode
	dList.Size++

	return dNode
}

// 从头部删除节点
func (dList *DList) Delete4Head() *DNode {
	if dList.GetSize() == 0 {
		return nil
	}

	dList.Mutex.Lock()
	defer dList.Mutex.Unlock()

	headDNode := dList.GetHeadDNode()
	// 双向链表节点数为多个时
	if dList.Head.Next != nil {
		dList.Head = headDNode.Next
		dList.Head.Prev = nil
		headDNode.Next = nil
	} else { // 双向链表节点数为一个时
		dList.Head, dList.Tail = nil, nil
	}

	dList.Size--
	return headDNode
}

// 从尾部删除节点
func (dList *DList) Delete4Tail() *DNode {
	if dList.GetSize() == 0 {
		return nil
	}

	dList.Mutex.Lock()
	defer dList.Mutex.Unlock()

	tailDNode := dList.GetTailDNode()
	// 双向链表节点数为多个时
	if dList.Tail.Prev != nil {
		dList.Tail = tailDNode.Prev
		dList.Tail.Next = nil
		tailDNode.Prev = nil
	} else { // 双向链表节点数为一个时
		dList.Head, dList.Tail = nil, nil
	}

	dList.Size--
	return tailDNode
}

// 向双链表删除指定位置节点
func (dList *DList) Delete4Index(index uint) *DNode {
	if dList.GetSize() == 0 || index > dList.Size-1 {
		return nil
	}

	if index == 0 {
		headDNode := dList.Head
		dList.Delete4Head()
		return headDNode
	}
	if index == dList.Size-1 {
		tailDNode := dList.Tail
		dList.Delete4Tail()
		return tailDNode
	}

	dNode := dList.GetDNodeByIndex(index)

	dList.Mutex.Lock()
	defer dList.Mutex.Unlock()

	dNode.Prev.Next = dNode.Next
	dNode.Next.Prev = dNode.Prev
	dList.Size--

	return dNode
}

// 在双链表删除指定节点
func (dList *DList) Remove(dNode *DNode) *DNode {
	if dNode == nil {
		return nil
	}
	if dList.isHead(dNode) {
		return dList.HeadPop()
	} else if dList.isTail(dNode) {
		return dList.TailPop()
	} else {
		dList.Mutex.Lock()
		dNode.Prev.Next = dNode.Next
		dNode.Next.Prev = dNode.Prev
		dNode.Next = nil
		dNode.Prev = nil
		dList.Size--
		dList.Mutex.Unlock()
	}
	return dNode
}

// 在双向链表里从头部开始查找符合节点值的节点地址， 一旦找到就会返回而不会再继续往下寻找
func (dList *DList) Search(data interface{}) *DNode {
	if dList.Size == 0 {
		return nil
	}

	dNode := dList.Head
	for ; dNode != nil; dNode = dNode.Next {
		if dNode.Data == data {
			break
		}
	}
	return dNode
}