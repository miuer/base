package tree_test

import (
	"fmt"
	"testing"

	bst "github.com/miuer/base/structure/tree/binary-search-tree"
)

func TestInsert(t *testing.T) {
	tree := bst.NewBinarySearchTree()

	tree.Insert(8, "8")
	tree.Insert(3, "3")
	tree.Insert(1, "1")
	tree.Insert(6, "6")
	tree.Insert(4, "4")
	tree.Insert(7, "7")
	tree.Insert(10, "10")

	tree.Insert(13, "13")
	tree.Insert(14, "14")

	min := tree.FindMin()
	t.Log(min.Value)

	max := tree.FindMax()
	t.Log(max.Value)

	tree.Print(tree.Root())

	fmt.Println()
	tree.Remove(4)

	tree.Print(tree.Root())
}
