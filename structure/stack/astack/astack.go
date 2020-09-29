package astack

import (
	"errors"
)

var (
	errStackEmpty = errors.New("StackIsEmpty")
)

// Astack -
type Astack struct {
	array []interface{}
	top   int
}

// Init -
func (as *Astack) Init() *Astack {
	as.array = as.array[:0]
	as.top = -1

	return as
}

// NewAstack -
func NewAstack() *Astack {
	return new(Astack).Init()
}

// Push -
func (as *Astack) Push(value interface{}) (*Astack, error) {
	return as.push(value), nil
}

func (as *Astack) push(value interface{}) *Astack {
	as.array = append(as.array, value)
	as.top++
	return as
}

// Pop -
func (as *Astack) Pop() (interface{}, *Astack, error) {
	top, err := as.Top()
	astack := as.pop()

	return top, astack, err
}

func (as *Astack) pop() *Astack {
	as.array = as.array[:as.top]
	as.top--
	return as
}

// Top -
func (as *Astack) Top() (interface{}, error) {
	if as.top > -1 {
		return as.array[as.top], nil
	}

	return nil, errStackEmpty
}

// IsEmpty -
func (as *Astack) IsEmpty() bool {
	return as.top == -1
}
