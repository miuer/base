package heap

// Node -
type Node struct {
	index int64
	value interface{}
}

// BinaryHeap -
type BinaryHeap struct {
	heap []*Node
}

// BinaryHeaper -
type BinaryHeaper interface {
	Init() *BinaryHeap
	Insert(int64, interface{}) bool
	DeleteMin() (*Node, error)
}

// Init -
func (bh *BinaryHeap) Init() *BinaryHeap {
	bh.heap = append(bh.heap, &Node{0, nil})
	bh.heap = bh.heap[:1]
	return bh
}

// NewBinaryHeap -
func NewBinaryHeap() *BinaryHeap {
	return new(BinaryHeap).Init()
}

// Heap -
func (bh *BinaryHeap) Heap() []*Node {
	return bh.heap
}

// Insert -
func (bh *BinaryHeap) Insert(index int64, value interface{}) bool {
	if index < 1 {
		return false
	}

	node := &Node{index, value}

	bh.insert(node)
	return true
}

func (bh *BinaryHeap) insert(node *Node) {
	var index = len(bh.heap)
	bh.heap = append(bh.heap, &Node{})

	for ; index > 1 && node.index < bh.heap[index/2].index; index /= 2 {
		bh.heap[index] = bh.heap[index/2]
	}
	bh.heap[index] = node
}

// IsEmpty -
func (bh *BinaryHeap) IsEmpty() bool {
	return len(bh.heap) == 1
}

// DeleteMin -
func (bh *BinaryHeap) DeleteMin() (*Node, error) {
	node := bh.heap[1]
	bh.heap[1] = bh.heap[len(bh.heap)-1]
	bh.heap = bh.heap[:len(bh.heap)-1]
	bh.percolateDown(1)
	return node, nil
}

func (bh *BinaryHeap) percolateDown(index int) {

	tmp := bh.heap[index]

	for index*2 < len(bh.heap) {
		child := 2 * index

		// find the smallest index in the child
		if child != len(bh.heap)-1 && bh.heap[child+1].index < bh.heap[child].index {
			child++
		}

		// swap
		if bh.heap[child].index < tmp.index {
			bh.heap[index] = bh.heap[child]
		} else {
			break
		}

		// set new index
		index = child
	}

	bh.heap[index] = tmp
}

// Remove -
func (bh *BinaryHeap) Remove() *Node {

	return nil
}
