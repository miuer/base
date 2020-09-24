package slist_test

import (
	"testing"

	"github.com/miuer/base/structure/list/slist"
)

func TestNew(t *testing.T) {
	sl := slist.NewSlist()

	if sl.Root().Value != "root" {
		t.Error("create new slist faild")
	}
}

func TestInsert(t *testing.T) {
	sl := slist.NewSlist()

	root := sl.Root()

	tail, _ := sl.InsertAfter("tail", root)
	sl.InsertAfter("after root", root)

	sl.InsertBefore("before root", root)
	sl.InsertBefore("before tail", tail)

	n := sl.Root()
	for {
		if n.Next() != nil {
			t.Log(n.Value)
			n = n.Next()
			continue
		}

		t.Log(n.Value)
		break
	}
}

func TestRemove(t *testing.T) {
	sl := slist.NewSlist()

	root := sl.Root()

	tail, _ := sl.InsertAfter("tail", root)
	after, _ := sl.InsertAfter("after root", root)

	sl.InsertBefore("before root", root)
	sl.InsertBefore("before tail", tail)

	sl.Remove(sl.Root())
	sl.Remove(tail)
	sl.Remove(after)

	n := sl.Root()
	for {
		if n.Next() != nil {
			t.Log(n.Value)
			n = n.Next()
			continue
		}

		t.Log(n.Value)
		break
	}
}
