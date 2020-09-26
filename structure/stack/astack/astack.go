package astack

import (
	"errors"
)

var (
	errStackEmpty = errors.New("StackIsEmpty")
)

// Lstatck -
type Lstatck struct {
	array []interface{}
	top   int
}

// Init -
func (l *Lstatck) Init() *Lstatck {
	l.array = l.array[:0]
	l.top = -1

	return l
}

// NewAstack -
func NewAstack() *Lstatck {
	return new(Lstatck).Init()
}

// Push -
func (l *Lstatck) Push(value interface{}) (*Lstatck, error) {
	return l.push(value), nil
}

func (l *Lstatck) push(value interface{}) *Lstatck {
	l.array = append(l.array, value)
	l.top++
	return l
}

// Pop -
func (l *Lstatck) Pop() (interface{}, *Lstatck, error) {
	top, err := l.Top()
	ls := l.pop()

	return top, ls, err
}

func (l *Lstatck) pop() *Lstatck {
	l.array = l.array[:l.top]
	l.top--
	return l
}

// Top -
func (l *Lstatck) Top() (interface{}, error) {
	if l.top > -1 {
		return l.array[l.top], nil
	}

	return nil, errStackEmpty
}
