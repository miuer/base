package main

import (
	"github.com/miuer/base/data-structure/list"
)

func main() {
	slist := list.NewSlist()

	root := slist.GetRoot()

	n1 := slist.InsertAfter("test", root)

	slist.InsertBefore("demo", n1)

}
