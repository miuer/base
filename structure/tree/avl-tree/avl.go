package tree

import (
	"fmt"
)

// Node -
type Node struct {
	height int
	index  int
	Value  interface{}

	left  *Node
	right *Node
}

// AVLTree -
type AVLTree struct {
	root *Node
}

// Init -
func (avl *AVLTree) Init() *AVLTree {
	avl.root = nil
	return avl
}

// NewAVLTree -
func NewAVLTree() *AVLTree {
	return new(AVLTree).Init()
}

// Root -
func (avl *AVLTree) Root() *Node {
	return avl.root
}

// Height -
func (node *Node) Height() int {
	if node == nil {
		return -1
	}

	return node.height
}

// left left
// α -> right
// α.left -> root
func (avl *AVLTree) rotateWithLeftChild(k *Node) *Node {
	n := k.left
	k.left = n.right
	n.right = k

	k.height = max(k.left.Height(), k.right.Height()) + 1
	n.height = max(n.left.Height(), k.height) + 1

	return n
}

// right right
// α -> left
// α.right -> root
func (avl *AVLTree) rotateWithRightChild(k *Node) *Node {
	n := k.right
	k.right = n.left
	n.left = k

	k.height = max(k.left.Height(), k.right.Height()) + 1
	n.height = max(n.right.Height(), k.height) + 1

	return n
}

// left right
// α -> right
// hightest -> root
// mid -> left

// α.left ->rotateWithRightChild
// α -> rotateWithLeftChild
func (avl *AVLTree) doubleWithLeftChild(k *Node) *Node {
	k.left = avl.rotateWithRightChild(k.left)
	return avl.rotateWithLeftChild(k)
}

// right left
// α -> left
// hightest -> root
// mid -> right

// α.right ->rotateWithLeftChild
// α -> rotateWithRightChild
func (avl *AVLTree) doubleWithRightChild(k *Node) *Node {
	k.right = avl.rotateWithLeftChild(k.right)
	return avl.rotateWithRightChild(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Insert -
func (avl *AVLTree) Insert(index int, value interface{}) {
	node := &Node{
		index: index,
		Value: value,
	}

	avl.root = avl.insert(avl.Root(), node)
}

func (avl *AVLTree) insert(node, newNode *Node) *Node {
	if node == nil {
		node = newNode
	}

	if newNode.index < node.index {
		node.left = avl.insert(node.left, newNode)
		if node.left.Height()-node.right.Height() == 2 {
			if newNode.index < node.left.index {
				node = avl.rotateWithLeftChild(node)
			} else {
				node = avl.doubleWithLeftChild(node)
			}
		}
	} else if node.index < newNode.index {
		node.right = avl.insert(node.right, newNode)
		if node.right.Height()-node.left.Height() == 2 {
			if node.right.index < newNode.index {
				node = avl.rotateWithRightChild(node)
			} else {
				node = avl.doubleWithRightChild(node)
			}
		}
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1

	return node
}

// Print -
func (avl *AVLTree) Print(node *Node) {
	if node == nil {
		return
	}

	avl.Print(node.left)
	fmt.Printf("%d ", node.index)
	avl.Print(node.right)
}
