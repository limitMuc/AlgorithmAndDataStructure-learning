package LRU

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	lru := Init(3)

	if lru.dList.Capacity == 3 {
		t.Log("LRU Cache init success")
	} else {
		t.Error("LRU Cache init failed")
	}
}

func TestPut(t *testing.T) {
	lruCache := Init(3)

	lruCache.Put(2)
	lruCache.Put(3)
	lruCache.Put(2)
	lruCache.Put(1)
	lruCache.Put(5)
	lruCache.Put(2)
	lruCache.Put(4)
	lruCache.Put(5)
	lruCache.Put(3)
	lruCache.Put(2)
	lruCache.Put(5)
	lruCache.Put(2)
	lruCache.PrintLRUCache()
}

func TestGet(t *testing.T) {
	lruCache := Init(3)

	lruCache.Put(2)
	lruCache.Put(3)
	lruCache.Put(2)

	fmt.Println(lruCache.Get(5))
	lruCache.PrintLRUCache()
}