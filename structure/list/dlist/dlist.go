package dlist

import "errors"

var (
	errDlistNotMatch = errors.New("DlistNotMatch")
)

// Node -
type Node struct {
	next, prev *Node
	Value      interface{}
}

// Dlist -
type Dlist struct {
	head, tail *Node
	len        int
}

// Init -
func (dl *Dlist) Init() *Dlist {
	dl.len = 0
	dl.head = &Node{Value: "head"}
	dl.tail = &Node{Value: "tail"}
	dl.head.next = dl.tail
	dl.head.prev = nil
	dl.tail.prev = dl.head
	dl.tail.next = nil

	return dl
}

// NewDlist -
func NewDlist() *Dlist {
	return new(Dlist).Init()
}

// Len -
func (dl *Dlist) Len() int {
	return dl.len
}

// Head -
func (dl *Dlist) Head() *Node {
	return dl.head
}

// Tail -
func (dl *Dlist) Tail() *Node {
	return dl.tail
}

// Next -
func (n *Node) Next() *Node {
	if node := n.next; node != nil {
		return node
	}

	return nil
}

// Prev -
func (n *Node) Prev() *Node {
	if node := n.prev; node != nil {
		return node
	}

	return nil
}

func (dl *Dlist) insert(newNode, at *Node) *Node {
	newNode.next = at.next
	newNode.prev = at
	at.next.prev = newNode
	at.next = newNode
	dl.len++

	return newNode
}

// InsertAfter -
func (dl *Dlist) InsertAfter(value interface{}, node *Node) (*Node, error) {

	n := &Node{
		Value: value,
	}

	if node.next == nil {
		// 尾部插入
		node.next = n
		n.prev = node
		n.next = nil
		dl.tail = n
		dl.len++
		return n, nil
	}

	return dl.insert(n, node), nil
}

// InsertBefore -
func (dl *Dlist) InsertBefore(value interface{}, node *Node) (*Node, error) {

	n := &Node{
		Value: value,
	}

	if node.prev == nil {
		// 头部插入
		node.prev = n
		n.next = node
		n.prev = nil
		dl.head = n
		dl.len++
		return n, nil
	}

	return dl.insert(n, node.prev), nil
}

func (dl *Dlist) remove(node *Node) *Node {
	node.prev.next = node.next
	node.next.prev = node.prev

	node = &Node{}

	dl.len--
	return node
}

// Remove -
func (dl *Dlist) Remove(node *Node) (*Node, error) {

	if node.prev == nil {
		node.next.prev = nil
		dl.head = node.next
		node = &Node{}
		dl.len--
		return dl.head, nil
	}

	if node.next == nil {
		node.prev.next = nil
		dl.tail = node.prev
		node = &Node{}
		dl.len--
		return dl.tail, nil
	}

	return dl.remove(node), nil
}

func (dl *Dlist) swap(node *Node) *Node {
	prev := node.Prev()
	next := node.Next()

	prev.next = next
	node.next = next.Next()
	node.prev = next
	next.prev = prev
	next.next = node
	next.Next().prev = node

	return node
}

// Swap -
func (dl *Dlist) Swap(node *Node) (*Node, error) {
	switch {
	case true:

	default:
		dl.swap(node)
	}

	return dl.swap(node), nil
}
