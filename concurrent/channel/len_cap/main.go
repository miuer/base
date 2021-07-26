package main

import "fmt"

func main() {
	bufChanFunc()
}

// 缓冲通道的长度大小为 channel 中含有的元素个数，容量大小是 make 初始化时给定的值
func bufChanFunc() {
	bufChan := make(chan int, 5)

	for i := 0; i < 3; i++ {
		bufChan <- i
	}

	fmt.Println("buffer channel len: ", len(bufChan))

	fmt.Println("buffer channel cap: ", cap(bufChan))

	close(bufChan)
}

// 无缓冲通道的长度和容量都为 0
func noBufChanFunc() {
	noBufChan := make(chan int)

	fmt.Println("no buffer channel len: ", len(noBufChan))

	fmt.Println("no buffer channel cap: ", cap(noBufChan))

	close(noBufChan)
}
