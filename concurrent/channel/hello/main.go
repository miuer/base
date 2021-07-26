package main

import "fmt"

func main() {
	var ch chan string

	// 引用类型必须初始化，切片，字典，通道，函数，接口
	// 向一个未初始化的 channel 获取数据将造成当前 goroutinue 永久阻塞
	ch = make(chan string)
	defer close(ch)

	go func() {
		ch <- "hello world"
	}()

	str, ok := <-ch
	if ok {
		fmt.Println("the received data is: ", str)
	} else {
		fmt.Println("no data from channel was received")
	}
}
