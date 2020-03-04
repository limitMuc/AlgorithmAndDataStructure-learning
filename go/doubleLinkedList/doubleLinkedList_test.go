package doubleLinkedList

import (
	"testing"
)

func TestInit(t *testing.T) {
	dList := Init(0)
	if dList.Size == 0 {
		t.Log("double list init success")
	} else {
		t.Error("double list init failed")
	}
}

func TestPushAndPop(t *testing.T) {
	dList := Init(0)

	dNode := dList.HeadPush(nil)
	if dNode != nil {
		t.Error("double list HeadPush nil failed")
	} else {
		t.Log("double list HeadPush nil success")
	}

	dNode = dList.HeadPush(&DNode{Data: 2})
	if dNode != nil {
		t.Log("double list HeadPush success")
	} else {
		t.Log("double list HeadPush failed")
	}

	dNode = dList.HeadPush(&DNode{Data: 1})
	if dNode != nil {
		t.Log("double list HeadPush success")
	} else {
		t.Log("double list HeadPush failed")
	}

	dList.PrintDoubleLinkedList()

	dNode = dList.TailPush(nil)
	if dNode != nil {
		t.Error("double list TailPush nil failed")
	} else {
		t.Log("double list TailPush nil success")
	}

	dNode = dList.TailPush(&DNode{Data: 3})
	if dNode != nil {
		t.Log("double list TailPush success")
	} else {
		t.Log("double list TailPush failed")
	}

	dNode = dList.TailPush(&DNode{Data: 4})
	if dNode != nil {
		t.Log("double list TailPush success")
	} else {
		t.Log("double list TailPush failed")
	}

	dList.PrintDoubleLinkedList()

	headDNode := dList.HeadPop()
	if headDNode != nil {
		t.Log("double list HeadPop success")
	} else {
		t.Error("double list HeadPop failed")
	}
	dList.PrintDoubleLinkedList()

	tailDNode := dList.TailPop()
	if tailDNode != nil {
		t.Log("double list TailPop success")
	} else {
		t.Error("double list TailPop failed")
	}
	dList.PrintDoubleLinkedList()
}

func TestGetDNodeByIndex(t *testing.T) {
	dList := Init(0)

	dNode := dList.GetDNodeByIndex(5)
	if dNode != nil {
		t.Error("double list GetDNodeByIndex failed")
	} else {
		t.Log("double list GetDNodeByIndex success")
	}

	dNode = dList.Append(&DNode{Data: 1})
	if dNode != nil {
		t.Log("double list Append success")
	} else {
		t.Log("double list Append failed")
	}

	dNode = dList.Append(&DNode{Data: 2})
	if dNode != nil {
		t.Log("double list Append success")
	} else {
		t.Log("double list Append failed")
	}

	dNode = dList.Append(&DNode{Data: 3})
	if dNode != nil {
		t.Log("double list Append success")
	} else {
		t.Log("double list Append failed")
	}

	dList.PrintDoubleLinkedList()

	dNode = dList.GetDNodeByIndex(1)
	if dNode != nil {
		t.Logf("dNode: %v\n", dNode.Data)
	} else {
		t.Error("double list GetDNodeByIndex failed")
	}
}

func TestRemove(t *testing.T) {
	dList := Init(0)

	dList.Append(&DNode{Data: 1})
	dList.Append(&DNode{Data: 2})
	dList.Append(&DNode{Data: 3})
	dList.PrintDoubleLinkedList()

	dList.Remove(dList.Head.Next)
	dList.PrintDoubleLinkedList()
}

func TestDelete4Head(t *testing.T) {
	dList := Init(0)

	dList.Append(&DNode{Data: 1})
	dList.PrintDoubleLinkedList()

	dList.Delete4Head()
}