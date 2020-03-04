package LFU

import (
	"AlgorithmAndDataStructure-learning/go/doubleLinkedList"
	"fmt"
	"testing"
)

func TestGetMapMinKey(t *testing.T) {
	freqMap := make(map[int]*doubleLinkedList.DList)
	freqMap[100] = nil
	freqMap[10] = nil
	freqMap[33] = nil
	freqMap[5] = nil
	fmt.Println(freqMap)
	fmt.Println(GetMapMinKey(freqMap))
}

func TestInit(t *testing.T) {
	lfu := Init(3)
	if lfu.dList.Capacity == 3 {
		t.Log("LFU Cache init success")
	} else {
		t.Error("LFU Cache init failed")
	}
}

func TestPut(t *testing.T) {
	lfuCache := Init(3)

	lfuCache.Put(1)
	lfuCache.PrintLFUCache()
	lfuCache.Put(2)
	lfuCache.PrintLFUCache()
	lfuCache.Put(3)
	lfuCache.PrintLFUCache()
	lfuCache.Put(4)
	lfuCache.PrintLFUCache()
	lfuCache.Put(1)
	lfuCache.PrintLFUCache()
	lfuCache.Put(2)
	lfuCache.PrintLFUCache()
	lfuCache.Put(5)
	lfuCache.PrintLFUCache()
	lfuCache.Put(1)
	lfuCache.PrintLFUCache()
	lfuCache.Put(2)
	lfuCache.PrintLFUCache()
	lfuCache.Put(3)
	lfuCache.PrintLFUCache()
	lfuCache.Put(4)
	lfuCache.PrintLFUCache()
	lfuCache.Put(5)
	lfuCache.PrintLFUCache()
}

func TestGet(t *testing.T) {
	lfuCache := Init(3)

	lfuCache.Put(1)
	lfuCache.Put(2)
	lfuCache.Put(3)
	lfuCache.PrintLFUCache()

	t.Log(lfuCache.Get(2))
	lfuCache.PrintLFUCache()
}