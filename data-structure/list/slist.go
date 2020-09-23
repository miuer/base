package list

import (
	"errors"
)

var (
	errSlistNotMatch = errors.New("SlistNotMatch")
)

// Node -
type Node struct {
	Value interface{}
	next  *Node
	slist *Slist
}

// Slist -
type Slist struct {
	root Node
	len  int
}

// Slister -
type Slister interface {
}

// Init -
func (sl *Slist) Init() *Slist {
	sl.root.next = nil
	sl.root.Value = nil
	sl.root.slist = sl
	sl.len = 1
	return sl
}

// NewSlist -
func NewSlist() *Slist {
	return new(Slist).Init()
}

// GetRoot -
func (sl *Slist) GetRoot() *Node {
	return &sl.root
}

// Len -
func (sl *Slist) Len() int {
	return sl.len
}

// Next -
func (node *Node) Next() (*Node, error) {
	if node.next != nil {
		return node.next, nil
	}

	return nil, errors.New("no next node")
}

func (sl *Slist) insert(newNode, at *Node) *Node {
	newNode.next = at.next
	at.next = newNode
	sl.len++

	return newNode
}

// InsertAfter -
func (sl *Slist) InsertAfter(value interface{}, node *Node) (*Node, error) {
	if node.slist != sl {
		return nil, errSlistNotMatch
	}

	n := &Node{
		Value: value,
		slist: sl,
	}

	return sl.insert(n, node), nil
}

// InsertBefore - not support inserting in the header
func (sl *Slist) InsertBefore(value interface{}, node *Node) (*Node, error) {
	if node.slist != sl {
		return nil, errSlistNotMatch
	}

	n := &Node{
		Value: value,
		slist: sl,
	}

	if node == &sl.root {
		//todo
	}

	prev := sl.findPrevNode(node)

	return sl.insert(n, prev), nil

}

func (sl *Slist) remove(prev, node *Node) *Node {
	prev.next = node.next
	node.Value = nil
	node.next = nil
	node.slist = nil
	sl.len--

	return node
}

// Remove -
func (sl *Slist) Remove(node *Node) (*Node, error) {
	if node.slist != sl {
		return nil, errSlistNotMatch
	}

	if node == &sl.root {
		sl.root = *sl.root.next
		return &sl.root, nil
	}

	prev := sl.findPrevNode(node)
	return sl.remove(prev, node), nil
}

func (sl *Slist) findPrevNode(node *Node) *Node {
	if node.slist != sl {
		return nil
	}

	goal := &sl.root

	for {
		// find goal
		if goal.next == node {
			return goal

			// not find goal
		} else if goal.next == nil {
			return nil
		}
		goal = goal.next
	}
}
