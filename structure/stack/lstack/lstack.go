package lstack

import (
	"errors"

	"github.com/miuer/base/structure/list/slist"
)

var (
	errStackEmpty = errors.New("StackIsEmpty")
)

// Lstack -
type Lstack struct {
	sl  *slist.Slist
	top int
}

// Init -
func (ls *Lstack) Init() *Lstack {
	ls.sl = slist.NewSlist()
	ls.top = -1
	return ls
}

// NewLstack -
func NewLstack() *Lstack {
	return new(Lstack).Init()
}

func (ls *Lstack) push(value interface{}) *Lstack {
	root := ls.sl.Root()
	ls.sl.InsertBefore(value, root)
	ls.top++

	return ls
}

// Push -
func (ls *Lstack) Push(value interface{}) (*Lstack, error) {
	ls = ls.push(value)
	return ls, nil
}

// Top -
func (ls *Lstack) Top() (interface{}, error) {
	if ls.top > -1 {
		ls.top--
		return ls.sl.Root().Value, nil
	}

	return nil, errStackEmpty
}

func (ls *Lstack) pop() *Lstack {
	ls.sl.Remove(ls.sl.Root())
	return ls
}

// Pop -
func (ls *Lstack) Pop() (interface{}, *Lstack, error) {
	top, err := ls.Top()
	ls = ls.pop()
	return top, ls, err
}

// IsEmpty -
func (ls *Lstack) IsEmpty() bool {
	return ls.top == -1
}
