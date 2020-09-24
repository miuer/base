/*
	单向链表解决了数组插入，删除的消耗
	但是因为单向链表中节点无前驱节点的信息，因此在尾部进行插入和删除所花费的支出和数组一样，
	而且因为无前驱节点的信息，因此在向前插入需要进行遍历，从而引入双向链表
	单向链表不存在索引，其查询效率不如数组
*/

package slist

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
	root *Node
	len  int
}

// Init -
func (sl *Slist) Init() *Slist {

	sl.root = &Node{
		Value: "root",
		slist: sl,
		next:  nil,
	}

	sl.len = 0
	return sl
}

// NewSlist -
func NewSlist() *Slist {
	return new(Slist).Init()
}

// Root -
func (sl *Slist) Root() *Node {
	return sl.root
}

// Len -
func (sl *Slist) Len() int {
	return sl.len
}

// Next -
func (node *Node) Next() *Node {
	if node.next != nil {
		return node.next
	}
	return nil
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

	if node == sl.root {
		n.next = node
		sl.root = n
		sl.len++
		return n, nil
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

	if node == sl.root {
		sl.root = sl.root.next
		return sl.root, nil
	}

	prev := sl.findPrevNode(node)
	return sl.remove(prev, node), nil
}

func (sl *Slist) findPrevNode(node *Node) *Node {
	if node.slist != sl {
		return nil
	}

	goal := sl.root

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
