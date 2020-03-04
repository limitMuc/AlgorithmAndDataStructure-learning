package FIFO

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	fifoCache := Init(3)
	if fifoCache.dList.Capacity == 3 {
		t.Log("FIFO Cache init success")
	} else {
		t.Error("FIFO Cache init failed")
	}
}

func TestGet(t *testing.T) {
	fifoCache := Init(3)

	fifoCache.Put(1)
	fifoCache.Put(2)
	fifoCache.Put(3)

	fmt.Println(fifoCache.Get(5))

}

func TestPut(t *testing.T) {
	fifoCache := Init(3)

	fifoCache.Put(1)
	fifoCache.Put(2)
	fifoCache.Put(3)
	fifoCache.Put(4)
	fifoCache.Put(1)
	fifoCache.Put(2)
	fifoCache.Put(5)
	fifoCache.Put(1)
	fifoCache.Put(2)
	fifoCache.Put(3)
	fifoCache.Put(4)
	fifoCache.Put(5)
	fifoCache.Put(6)
	fifoCache.PrintFIFOCache()

	fmt.Println(fifoCache.Get(3))
}