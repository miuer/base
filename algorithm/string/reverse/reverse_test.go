package reverse_test

import (
	"testing"

	"github.com/miuer/base/string/reverse"
)

func TestReverse(t *testing.T) {
	var str string = "abcde"

	t.Log(reverse.Reverse(str))
}
