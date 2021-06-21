package maxmin_test

import (
	"testing"

	"github.com/miuer/base/algorithm/array/maxmin"
)

func TestMaxmin(t *testing.T) {

	arr := []int{7, 3, 19, 40, 4, 7, 1, 5}

	max, min := maxmin.MiuerImpl(arr, 0, len(arr)-1)

	t.Log("max: ", max)
	t.Log("min: ", min)

}
