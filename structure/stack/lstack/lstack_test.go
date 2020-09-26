package lstack_test

import (
	"testing"

	"github.com/miuer/base/structure/stack/lstack"
)

func TestLstack(t *testing.T) {
	ls := lstack.NewLstack()

	ls.Push("test push 1")
	ls.Push("test push 2")

	top, err := ls.Top()
	if err != nil {
		t.Log(err)
	}

	t.Log(top)

	ls.Pop()
	ls.Pop()

	top, err = ls.Top()
	if err != nil {
		t.Log(err)
	}

	t.Log(top)
}
