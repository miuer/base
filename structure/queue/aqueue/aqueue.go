package aqueue

import (
	"errors"
)

var (
	errQueueEmpty = errors.New("QueueIsEmpty")
)

// Aqueue -
type Aqueue struct {
	queue []interface{}
}

// Init -
func (aq *Aqueue) Init() *Aqueue {
	aq.queue = aq.queue[:0]

	return aq
}

// NewAqueue -
func NewAqueue() *Aqueue {
	return new(Aqueue).Init()
}

func (aq *Aqueue) enQueue(value interface{}) *Aqueue {
	aq.queue = append(aq.queue, value)
	return aq
}

// EnQueue -
func (aq *Aqueue) EnQueue(value interface{}) (*Aqueue, error) {
	return aq.enQueue(value), nil
}

// Top -
func (aq *Aqueue) Top() (interface{}, error) {
	if len(aq.queue) > 0 {
		return aq.queue[0], nil
	}

	return nil, errQueueEmpty
}

func (aq *Aqueue) deQueue() *Aqueue {
	if len(aq.queue) > 0 {
		aq.queue = aq.queue[1:]
		return aq
	}

	return aq.Init()
}

// DeQueue -
func (aq *Aqueue) DeQueue() (interface{}, *Aqueue, error) {
	top, err := aq.Top()
	queue := aq.deQueue()

	return top, queue, err
}

// IsEmpty -
func (aq *Aqueue) IsEmpty() bool {
	if len(aq.queue) == 0 {
		return true
	}

	return false
}
