package dlist_test

import (
	"testing"

	"github.com/miuer/base/structure/list/dlist"
)

func TestNew(t *testing.T) {
	dl := dlist.NewDlist()

	t.Log(dl.Head().Value)
	t.Log(dl.Tail().Value)
}

func TestInsert(t *testing.T) {
	dl := dlist.NewDlist()

	dl.InsertAfter("after head", dl.Head())
	dl.InsertBefore("before head", dl.Head())
	dl.InsertBefore("before tail", dl.Tail())
	dl.InsertAfter("after tail", dl.Tail())

	n := dl.Head()
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
	dl := dlist.NewDlist()

	afterHead, _ := dl.InsertAfter("after head", dl.Head())
	dl.InsertBefore("before head", dl.Head())
	dl.InsertBefore("before tail", dl.Tail())
	dl.InsertAfter("after tail", dl.Tail())

	dl.Remove(afterHead)
	dl.Remove(dl.Head())
	dl.Remove(dl.Tail())

	n := dl.Head()
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

func TestSwap(t *testing.T) {
	dl := dlist.NewDlist()

	afterHead, _ := dl.InsertAfter("after head", dl.Head())
	dl.InsertBefore("before head", dl.Head())
	dl.InsertBefore("before tail", dl.Tail())
	dl.InsertAfter("after tail", dl.Tail())

	n := dl.Head()
	for {
		if n.Next() != nil {
			t.Log(n.Value)
			n = n.Next()
			continue
		}

		t.Log(n.Value)
		break
	}

	t.Logf("\n")

	dl.Swap(afterHead)

	n = dl.Head()
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
