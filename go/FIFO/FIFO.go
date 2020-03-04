package FIFO

import (
	"AlgorithmAndDataStructure-learning/go/doubleLinkedList"
	"fmt"
	"sync"
)

type FIFOCache struct {
	dList          *doubleLinkedList.DList
	kMap           map[interface{}]*doubleLinkedList.DNode // 键值对映射关系， 存储key和对应的双向链表中的DNode的地址， 方便快速通过节点的值从而找到节点的地址
	PageFaultCount int                                     // 记录缺页中断的次数
}

// 初始化
func Init(capacity uint) *FIFOCache {
	return &FIFOCache{
		dList: &doubleLinkedList.DList{
			Capacity: capacity,
			Size:     0,
			Mutex:    new(sync.RWMutex),
		},
		kMap: make(map[interface{}]*doubleLinkedList.DNode),
	}
}

func (fifo *FIFOCache) GetPageFaultCount() int {
	return fifo.PageFaultCount
}

func (fifo *FIFOCache) Get(key interface{}) interface{} {
	fifo.dList.Mutex.RLock()
	defer fifo.dList.Mutex.RUnlock()

	if dNode, ok := fifo.kMap[key]; !ok {
		return dNode
	} else {
		return dNode.Data
	}
}

func (fifo *FIFOCache) Put(key interface{}) {
	if fifo.dList.Capacity == 0 {
		return
	}
	if _, ok := fifo.kMap[key]; !ok && fifo.dList.Size <= fifo.dList.Capacity {
		if fifo.dList.Size == fifo.dList.Capacity {
			dNode := fifo.dList.Head
			fifo.dList.HeadPop()
			delete(fifo.kMap, dNode.Data)
		}
		dNode := &doubleLinkedList.DNode{
			Data: key,
		}
		fifo.dList.Append(dNode)
		fifo.kMap[key] = dNode
		fifo.PageFaultCount++
	}
}

func (fifo *FIFOCache) PrintFIFOCache() {
	fmt.Printf("PageFaultCount: %v   Capacity: %v   ", fifo.PageFaultCount, fifo.dList.Capacity)
	fifo.dList.PrintDoubleLinkedList()
}
