package tree

import (
	"errors"
	"fmt"

	"github.com/miuer/base/structure/stack/astack"
)

var (
	errNodeExists   = errors.New("NodeAlreadyExists")
	errNodeNULL     = errors.New("NodeIsNil")
	errNodePosition = errors.New("NodePositionError")
)

// Node -
type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

// BinaryTree -
type BinaryTree struct {
	root     *Node
	traverse []*Node
}

// Init -
func (bt *BinaryTree) Init() *BinaryTree {
	bt.root = &Node{}
	return bt
}

// NewBinaryTree -
func NewBinaryTree() *BinaryTree {
	return new(BinaryTree).Init()
}

// Root -
func (bt *BinaryTree) Root() *Node {
	return bt.root
}

// Print -
func (bt *BinaryTree) Print() {
	for _, v := range bt.traverse {
		fmt.Printf("%v ", v.Value)
	}
}

func newChild(value interface{}) *Node {
	return &Node{Value: value}
}

func (n *Node) insert(child *Node, position int) error {
	if n == nil {
		return errNodeNULL
	}

	switch position {
	case 0:
		if n.Left == nil {
			n.Left = child
			return nil
		}
		left := n.Left
		child.Left = left
		n.Left = child
		return nil
	case 1:
		if n.Right == nil {
			n.Right = child
			return nil
		}
		right := n.Right
		child.Right = right
		n.Right = child
		return nil
	default:
		return errNodePosition
	}
}

// Insert -
func (n *Node) Insert(value interface{}, position int) (*Node, error) {
	child := newChild(value)
	err := n.insert(child, position)
	return child, err
}

func (n *Node) remove() error {
	/*
		1. root
		2. others
	*/

	return nil
}

func (bt *BinaryTree) preorder(node *Node) {
	if node == nil {
		return
	}

	bt.traverse = append(bt.traverse, node)
	bt.preorder(node.Left)
	bt.preorder(node.Right)
}

// Preorder -
func (bt *BinaryTree) Preorder() {
	bt.traverse = bt.traverse[:0]
	bt.preorder(bt.Root())
}

func (bt *BinaryTree) preorderNonRecursive(node *Node) {
	if node == nil {
		return
	}

	stack := astack.NewAstack()
	stack.Push(node)

	for !stack.IsEmpty() {
		top, _ := stack.Top()
		node = top.(*Node)

		bt.traverse = append(bt.traverse, node)
		stack.Pop()

		if node.Right != nil {
			stack.Push(node.Right)
		}

		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

// PreorderNonRecursive -
func (bt *BinaryTree) PreorderNonRecursive() {
	bt.traverse = bt.traverse[:0]
	bt.preorderNonRecursive(bt.Root())
}

func (bt *BinaryTree) inorder(node *Node) {
	if node == nil {
		return
	}

	bt.inorder(node.Left)
	bt.traverse = append(bt.traverse, node)
	bt.inorder(node.Right)
}

// Inorder -
func (bt *BinaryTree) Inorder() {
	bt.traverse = bt.traverse[:0]
	bt.inorder(bt.Root())
}

func (bt *BinaryTree) inorderNonRecursive(node *Node) {
	stack := astack.NewAstack()
	for node != nil || !stack.IsEmpty() {
		if node != nil {
			stack.Push(node)
			node = node.Left
		} else {
			top, _ := stack.Top()
			node = top.(*Node)
			bt.traverse = append(bt.traverse, node)
			stack.Pop()
			node = node.Right
		}
	}
}

// InorderNonRecursive -
func (bt *BinaryTree) InorderNonRecursive() {
	bt.traverse = bt.traverse[:0]
	bt.inorderNonRecursive(bt.Root())
}

func (bt *BinaryTree) postorder(node *Node) {
	if node == nil {
		return
	}

	bt.postorder(node.Left)
	bt.postorder(node.Right)
	bt.traverse = append(bt.traverse, node)

}

// Postorder -
func (bt *BinaryTree) Postorder() {
	bt.traverse = bt.traverse[:0]
	bt.postorder(bt.Root())
}

func (bt *BinaryTree) postorderNonRecursive(node *Node) {
	stack := astack.NewAstack()
	stack.Push(node)

	for !stack.IsEmpty() {
		top, _ := stack.Top()
		node = top.(*Node)
		bt.traverse = append([]*Node{node}, bt.traverse...)
		stack.Pop()

		if node.Left != nil {
			stack.Push(node.Left)
		}
		if node.Right != nil {
			stack.Push(node.Right)
		}

	}
}

// PostorderNonRecursive -
func (bt *BinaryTree) PostorderNonRecursive() {
	bt.traverse = bt.traverse[:0]
	bt.postorderNonRecursive(bt.Root())
}
