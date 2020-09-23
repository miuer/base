package list

import "log"

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
func (node *Node) Next() *Node {
	// 最后一个节点会报错
	return node.next
}

func (sl *Slist) insert(newNode, at *Node) *Node {
	newNode.next = at.next
	at.next = newNode
	sl.len++

	return newNode
}

// InsertAfter -
func (sl *Slist) InsertAfter(value interface{}, node *Node) *Node {
	if node.slist != sl {
		return nil
	}

	n := &Node{
		Value: value,
		slist: sl,
	}

	return sl.insert(n, node)
}

// InsertBefore -
func (sl *Slist) InsertBefore(value interface{}, node *Node) *Node {
	if node.slist != sl {
		return nil
	}

	n := &Node{
		Value: value,
		slist: sl,
	}

	if node == &sl.root {
		//todo finish this bug!
		n.next = node

		log.Println(n.next.Value) // print nil
		// ???
		sl.root = *n
		log.Println(n.next.next.Value) // print nil

		return n
	}

	prev := sl.findPrevNode(node)

	sl.insert(n, prev)

	return n
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
func (sl *Slist) Remove(node *Node) *Node {
	if node.slist != sl {
		return nil
	}

	if node == &sl.root {
		// root 会不会成为垃圾，一直占据内存
		sl.root = *sl.root.next
		return &sl.root
	}

	prev := sl.findPrevNode(node)
	return sl.remove(prev, node)
}

func (sl *Slist) findPrevNode(node *Node) *Node {
	// 不属于同一链表，直接返回空值
	if node.slist != sl {
		return nil
	}

	goal := &sl.root

	for {
		// 找到目标节点
		if goal.next == node {
			return goal

			// 无目标节点
		} else if goal.next == nil {
			return nil
		}
		goal = goal.next
	}
}
