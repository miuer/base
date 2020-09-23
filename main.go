package main

import (
	"fmt"

	"github.com/miuer/base/data-structure/list"
)

func main() {
	slist := list.NewSlist()

	root := slist.GetRoot()

	n1 := slist.InsertAfter("test", root)

	n2 := slist.InsertBefore("demo", n1)

	slist.InsertBefore("root", root)

	fmt.Println(slist.GetRoot().Value)
	fmt.Println(slist.GetRoot().Next().Value)        // should print nil
	fmt.Println(slist.GetRoot().Next().Next().Value) // should print nil

	slist.Remove(n2)
	fmt.Println(slist.GetRoot().Value)
	fmt.Println(slist.GetRoot().Next().Value)

}
