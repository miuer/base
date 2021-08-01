package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var sl atomic.Value
	sl.Store([]int{1, 3, 5, 7})
	store(sl)
	fmt.Printf("value: %+d \n", sl.Load())
}

func store(sl atomic.Value) {
	sl.Store([]int{2, 4, 6, 8})
}
