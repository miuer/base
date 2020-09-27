package lqueue

import (
	"errors"

	"github.com/miuer/base/structure/list/slist"
)

var (
	errQueueEmpty = errors.New("QueueIsEmpty")
)

// Lqueue -
type Lqueue struct {
	queue *slist.Slist
	len   int
}

// Init -
func (lq *Lqueue) Init() *Lqueue {
	lq.queue = slist.NewSlist()
	lq.len = 0
	return lq
}

// NewLqueue -
func NewLqueue() *Lqueue {
	return new(Lqueue).Init()
}

func (lq *Lqueue) enQueue(value interface{}) *Lqueue {
	prev := lq.queue.FindPrevNode(nil)
	lq.queue.InsertAfter(value, prev)
	lq.len++

	return lq
}

// EnQueue -
func (lq *Lqueue) EnQueue(value interface{}) (*Lqueue, error) {
	return lq.enQueue(value), nil
}

// Top -
func (lq *Lqueue) Top() (interface{}, error) {
	if lq.len > 0 {
		top := lq.queue.Root().Next().Value
		return top, nil
	}

	return nil, errQueueEmpty
}

func (lq *Lqueue) deQueue() *Lqueue {
	if lq.len > 0 {
		lq.queue.Remove(lq.queue.Root())
		lq.len--
		return lq
	}

	return lq.Init()
}

// DeQueue -
func (lq *Lqueue) DeQueue() (interface{}, *Lqueue, error) {
	top, err := lq.Top()
	queue := lq.deQueue()

	return top, queue, err
}
