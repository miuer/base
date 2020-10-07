package tree_test

import (
	"fmt"
	"testing"

	avl "github.com/miuer/base/structure/tree/avl-tree"
)

func TestAVLInsert(t *testing.T) {
	tree := avl.NewAVLTree()

	var arr = []int{2, 3, 7, 10, 10, 10, 10, 23, 9, 102, 109, 111, 112, 113}

	fmt.Println(tree.Root().Height())

	for i := 0; i < len(arr); i++ {
		tree.Insert(arr[i], arr[i])
	}

	tree.Print(tree.Root())

	fmt.Println()

	fmt.Println(tree.Root().Height())

}
