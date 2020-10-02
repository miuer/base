package tree

import (
	"errors"
)

var (
	errIndexLessThanOne = errors.New("NodeIndexMustBeGreaterThanZero")
)

// Node -
type Node struct {
	// index > 0
	index int
	Value interface{}
	Left  *Node
	Right *Node
}

// BST -
type BST struct {
	root *Node
}

// Init -
func (bst *BST) Init() *BST {
	bst.root = &Node{}
	return bst
}

// NewBinarySearchTree -
func NewBinarySearchTree() *BST {
	return new(BST).Init()
}

func (bst *BST) contains(node *Node, index int) bool {
	if node == nil {
		return false
	}

	switch {
	case index < node.index:
		return bst.contains(node.Left, index)
	case node.index < index:
		return bst.contains(node.Right, index)
	default:
		return true
	}
}

// IsContains -
func (bst *BST) IsContains(index int) bool {
	return bst.contains(bst.root, index)
}

func (bst *BST) findMin(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Left == nil {
		return node
	}

	return bst.findMin(node.Left)
}

func (bst *BST) findMinNonRecursive(node *Node) *Node {
	if node != nil {
		for node.Left != nil {
			node = node.Left
		}
	}

	return node
}

// FindMin -
func (bst *BST) FindMin() *Node {
	return bst.findMin(bst.root)
}

func (bst *BST) findMax(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Right == nil {
		return node
	}

	return bst.findMax(node.Right)
}

func (bst *BST) findMaxNonRecursive(node *Node) *Node {
	if node != nil {
		for node.Right != nil {
			node = node.Right
		}
	}

	return node
}

// FindMax -
func (bst *BST) FindMax() *Node {
	return bst.findMax(bst.root)
}

// Insert -
func (bst *BST) Insert(index int, value interface{}) bool {
	if index < 1 {
		return false
	}

	node := &Node{index: index, Value: value}

	if bst.root.index == 0 {
		bst.root = node
		return true
	}

	return bst.insert(bst.root, node)
}

func (bst *BST) insert(node, newNode *Node) bool {
	if newNode.index < node.index {
		if node.Left == nil {
			node.Left = newNode
		}
		bst.insert(node.Left, newNode)
	} else {
		if node.Right == nil {
			node.Right = newNode
		}
		bst.insert(node.Right, newNode)
	}

	return true
}
