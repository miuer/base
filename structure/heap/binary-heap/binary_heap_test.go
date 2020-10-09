package heap_test

import (
	"testing"

	heap "github.com/miuer/base/structure/heap/binary-heap"
)

func TestInsert(t *testing.T) {

	var arr []uint64

	arr = append(arr, 13, 14, 16, 19, 21, 19, 68, 65, 26, 32, 31)

	bh := heap.NewBinaryHeap()

	for _, v := range arr {
		bh.Insert(v, v)
	}

	for _, v := range bh.Heap() {
		t.Logf("%d\n", v)
	}

	t.Logf("\n")

	bh.DeleteMin()

	for _, v := range bh.Heap() {
		t.Logf("%d\n", v)
	}
}
