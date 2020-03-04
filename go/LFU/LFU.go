package LFU

import (
	"AlgorithmAndDataStructure-learning/go/doubleLinkedList"
	"fmt"
	"sort"
	"sync"
)

type LFUCache struct {
	dList          *doubleLinkedList.DList
	kMap           map[interface{}]*doubleLinkedList.DNode // 键值对映射关系， 存储key和对应的双向链表中的DNode的地址， 方便快速通过节点的值从而找到节点的地址
	freqMap        map[int]*doubleLinkedList.DList         // 键值对映射关系， 存储频次和对应的双向链表的地址
	PageFaultCount int                                     // 记录缺页中断的次数
}

func Init(capacity uint) *LFUCache {
	return &LFUCache{
		dList: &doubleLinkedList.DList{
			Capacity: capacity,
			Size:     0,
			Mutex:    new(sync.RWMutex),
		},
		kMap:    make(map[interface{}]*doubleLinkedList.DNode),
		freqMap: make(map[int]*doubleLinkedList.DList),
	}
}

func (lfu *LFUCache) GetPageFaultCount() int {
	return lfu.PageFaultCount
}

// 更新该节点的访问频次
func (lfu *LFUCache) UpdateDNodeFreq(dNode *doubleLinkedList.DNode) {
	if dNode == nil {
		return
	}

	// 将访问节点在访问频次map中的双向链表里移除
	dNode = lfu.freqMap[dNode.Freq].Remove(dNode)
	// 如果访问频次map中的双向链表没有节点了， 就可以删除这个频次了
	if lfu.freqMap[dNode.Freq].Size == 0 {
		delete(lfu.freqMap, dNode.Freq)
	}

	dNode.Freq++
	// 如果没有对应访问频次map的话， 就新增一个访问频次map的对应双向链表
	if _, ok := lfu.freqMap[dNode.Freq]; !ok {
		lfu.freqMap[dNode.Freq] = doubleLinkedList.Init(0)
	}
	lfu.freqMap[dNode.Freq].Append(dNode)
}

// 根据map的key进行升序排序
func GetMapMinKey(freqMap map[int]*doubleLinkedList.DList) (int, *doubleLinkedList.DList) {
	var keySlice []int
	for freq, _ := range freqMap {
		keySlice = append(keySlice, freq)
	}
	sort.Ints(keySlice)
	return keySlice[0], freqMap[keySlice[0]]
}

func (lfu *LFUCache) Get(key interface{}) interface{} {
	if dNode, ok := lfu.kMap[key]; !ok {
		return dNode
	} else {
		lfu.UpdateDNodeFreq(lfu.freqMap[lfu.kMap[key].Freq].Search(key))
		lfu.kMap[key].Freq++
		return dNode.Data
	}
}

func (lfu *LFUCache) Put(key interface{}) {
	if lfu.dList.Capacity == 0 {
		return
	}

	if lfu.dList.Size <= lfu.dList.Capacity {
		// 如果已经存在在双向链表中， 即缓存命中， 那么就只要更新该节点的访问频次即可
		if _, ok := lfu.kMap[key]; ok {
			lfu.UpdateDNodeFreq(lfu.freqMap[lfu.kMap[key].Freq].Search(key))
			lfu.kMap[key].Freq++
		} else { // 缓存没有命中

			if lfu.dList.Size == lfu.dList.Capacity {
				minFreq, _ := GetMapMinKey(lfu.freqMap)
				freqDListDNode := lfu.freqMap[minFreq].HeadPop()
				lfu.dList.Remove(lfu.kMap[freqDListDNode.Data])
				delete(lfu.kMap, freqDListDNode.Data)
			}
			dNode := &doubleLinkedList.DNode{
				Data: key,
				Freq: 1,
			}
			if _, isExist := lfu.freqMap[dNode.Freq]; !isExist {
				lfu.freqMap[dNode.Freq] = doubleLinkedList.Init(0)
			}

			lfu.dList.Append(dNode)

			lfu.PageFaultCount++
			lfu.kMap[key] = dNode

			lfu.freqMap[dNode.Freq].Append(&doubleLinkedList.DNode{
				Data: key,
				Freq: 1,
			})
		}
	}
}

func (lfu *LFUCache) PrintLFUCache() {
	fmt.Printf("PageFaultCount: %v   Capacity: %v   ", lfu.PageFaultCount, lfu.dList.Capacity)
	lfu.dList.PrintDoubleLinkedList()

	for freq, freqDList := range lfu.freqMap {
		fmt.Printf("freq: %v   ==>    ", freq)
		freqDList.PrintDoubleLinkedList()
	}
	fmt.Printf("\n***************************************分割线***************************************\n\n\n\n\n")
}
