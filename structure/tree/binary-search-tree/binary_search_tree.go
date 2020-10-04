package tree

import (
	"errors"
	"fmt"
)

var (
	errIndexLessThanOne = errors.New("NodeIndexMustBeGreaterThanZero")
)

// Node -
type Node struct {
	// index > 0
	index int
	Value interface{}
	left  *Node
	right *Node
}

// BST -
type BST struct {
	root *Node
}

// Init -
func (bst *BST) Init() *BST {
	bst.root = &Node{0, nil, nil, nil}

	return bst
}

// NewBinarySearchTree -
func NewBinarySearchTree() *BST {
	return new(BST).Init()
}

// Root -
func (bst *BST) Root() *Node {
	return bst.root
}

// Left -
func (node *Node) Left() *Node {
	return node.left
}

// Right -
func (node *Node) Right() *Node {
	return node.right
}

func (bst *BST) contains(node *Node, index int) bool {
	if node == nil {
		return false
	}

	switch {
	case index < node.index:
		return bst.contains(node.Left(), index)
	case node.index < index:
		return bst.contains(node.right, index)
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

	if node.Left() == nil {
		return node
	}

	return bst.findMin(node.Left())
}

func (bst *BST) findMinNonRecursive(node *Node) *Node {
	if node != nil {
		for node.left != nil {
			node = node.Left()
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

	if node.Right() == nil {
		return node
	}

	return bst.findMax(node.Right())
}

func (bst *BST) findMaxNonRecursive(node *Node) *Node {
	if node != nil {
		for node.right != nil {
			node = node.right
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

	node := &Node{index: index, Value: value, left: nil, right: nil}

	if bst.root.index == 0 {
		bst.root = node
		return true
	}

	return bst.insert(bst.Root(), node)
}

func (bst *BST) insert(node, newNode *Node) bool {
	if newNode.index < node.index {
		if node.left == nil {
			node.left = newNode
			return true
		}
		bst.insert(node.Left(), newNode)
	} else {
		if node.right == nil {
			node.right = newNode
			return true
		}
		bst.insert(node.right, newNode)
	}

	return true
}

// Remove -
func (bst *BST) Remove(index int) bool {
	if node := bst.remove(bst.root, index); node != nil {
		return true
	}

	return false
}

func (bst *BST) remove(node *Node, index int) *Node {
	if node == nil {
		return nil
	}

	if index < node.index {
		node.left = bst.remove(node.left, index)
		return node
	} else if node.index < index {
		node.right = bst.remove(node.right, index)
		return node
	}

	if node.left != nil && node.right != nil {
		min := bst.findMin(node.Right())
		node.Value = min.Value
		node.index = min.index

		node.right = bst.removeMin(node.right)
		return node
	}

	if node.left != nil {
		return node.left
	}

	return node.right
}

func (bst *BST) removeMin(node *Node) *Node {
	if node.left == nil {
		var right *Node
		right = node.right
		node.right = nil
		return right
	}

	node.left = bst.removeMin(node.left)

	return node
}

// Print -
func (bst *BST) Print(node *Node) {
	if node == nil {
		return
	}

	bst.Print(node.left)
	fmt.Printf("%d ", node.index)
	bst.Print(node.right)

}
