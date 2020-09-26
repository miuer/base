package astack_test

import (
	"testing"

	"github.com/miuer/base/structure/stack/astack"
)

func TestAstack(t *testing.T) {
	as := astack.NewAstack()

	as.Push("test push 1")

	as.Push("test push 2")

	top, err := as.Top()
	if err != nil {
		t.Log(err)
	}

	t.Log(top)

	as.Pop()
	as.Pop()

	top, err = as.Top()
	if err != nil {
		t.Log(err)
	}

	t.Log(top)

}
