package LRU

import (
	"AlgorithmAndDataStructure-learning/go/doubleLinkedList"
	"fmt"
	"sync"
)

type LRUCache struct {
	dList          *doubleLinkedList.DList
	kMap           map[interface{}]*doubleLinkedList.DNode // 键值对映射关系， 存储key和对应的双向链表中的DNode的地址， 方便快速通过节点的值从而找到节点的地址
	PageFaultCount int                                     // 记录缺页中断的次数
}

func Init(capacity uint) *LRUCache {
	return &LRUCache{
		dList: &doubleLinkedList.DList{
			Capacity: capacity,
			Size:     0,
			Mutex:    new(sync.RWMutex),
		},
		kMap: make(map[interface{}]*doubleLinkedList.DNode),
	}
}

func (lru *LRUCache) GetPageFaultCount() int {
	return lru.PageFaultCount
}

func (lru *LRUCache) Get(key interface{}) interface{} {
	if dNode, ok := lru.kMap[key]; !ok {
		return dNode
	} else {
		lru.dList.Remove(dNode)
		lru.dList.TailPush(&doubleLinkedList.DNode{Data: dNode.Data})
		lru.kMap[dNode.Data] = lru.dList.Tail
		return dNode.Data
	}
}

func (lru *LRUCache) Put(key interface{}) {
	if lru.dList.Capacity == 0 {
		return
	}

	if lru.dList.Size <= lru.dList.Capacity {
		if _, ok := lru.kMap[key]; ok {
			lru.Get(key)
		} else {
			if lru.dList.Size == lru.dList.Capacity {
				dNode := lru.dList.Head
				lru.dList.Remove(dNode)
				delete(lru.kMap, dNode.Data)
			}
			dNode := &doubleLinkedList.DNode{
				Data: key,
			}
			lru.dList.Append(dNode)
			lru.kMap[key] = dNode
			lru.PageFaultCount++
		}
	}
}

func (lru *LRUCache) PrintLRUCache() {
	fmt.Printf("PageFaultCount: %v   Capacity: %v   ", lru.PageFaultCount, lru.dList.Capacity)
	lru.dList.PrintDoubleLinkedList()
}
