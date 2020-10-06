package tree_test

import (
	"fmt"
	"testing"

	tree "github.com/miuer/base/structure/tree/binary-tree"
)

func TestInsert(t *testing.T) {
	bt := tree.NewBinaryTree()

	root := bt.Root()
	root.Value = "+"

	leftChild, _ := root.Insert("+", 0)
	rightChild, _ := root.Insert("*", 1)

	leftChild.Insert("a", 0)

	leftRightChild, _ := leftChild.Insert("*", 1)
	leftRightChild.Insert("b", 0)
	leftRightChild.Insert("c", 1)

	rightLeftChild, _ := rightChild.Insert("+", 0)
	rightChild.Insert("g", 1)

	rightLeftLeftChild, _ := rightLeftChild.Insert("*", 0)
	rightLeftChild.Insert("f", 1)

	rightLeftLeftChild.Insert("d", 0)
	rightLeftLeftChild.Insert("e", 1)

	bt.Inorder()
	bt.Print()

	fmt.Println()

	bt.Level()
	bt.Print()

	fmt.Println(bt.Height(bt.Root()))

}
