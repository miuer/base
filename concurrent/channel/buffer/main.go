package main

import (
	"fmt"
)

// 输出结果并不一定是”预期“的结果，主要是由于 Go 运行时的调度造成的这种假象
func main() {
	noBufChanFunc()
}

//	无缓冲通道需要接收方和发送方同时准备就绪才能执行，否则会在准备就绪的一侧堵塞
//	向无缓冲通道发送值的操作会被阻塞，直到至少有一个对应的接收操作就绪为止。接收操作会在对应发送操作之前完成
//	从无缓冲通道获取值的操作会被阻塞，直到至少有一个对应的发送操作就绪为止。发送操作会在对应接收操作之前完成
func noBufChanFunc() {
	var strChan = make(chan string)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go func() { // 接收
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second... [receiver]")
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	//	发送端尝试发送 "a" 时会阻塞，接收端接收 "a" 并打印信息，接收操作在发送操作之前完成
	//	接收端尝试获取值时阻塞，发送端发送 "b"并打印信息，发送操作在接收操作之前完成
	//	接收方和发送方交替上诉步骤直到通道关闭
	go func() { // 发送
		syncChan1 <- struct{}{}
		fmt.Println("Sent a sync signal. [sender]")

		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[sender]")

		}

		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}

//	缓冲通道队列是一个环形队列，一般接收方会从缓冲通道中获取值（复制两次），只有当缓冲队列为空时才会将值直接传递给接收方
//	经由通道传递的值至少会被复制一次（缓冲队列为空时），至多被复制两次，发送操作一定会在接收操作之前完成
//	向缓冲通道发送值的操作会在缓冲队列满时被阻塞，直到至少有一个对应的接收操作就绪为止
//	从缓冲通道接收值的操作会在缓冲队列为空时被阻塞，直到至少有一个对应的发送操作就绪为止
func bufChanFunc() {
	var strChan = make(chan string, 3)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go func() { // 接收
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second... [receiver]")
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	//	执行到 "c" 时 发送接收信号，并且尝试发送 "d", 由于缓冲队列已满，因此在发送 "d" 时阻塞
	//	接收端依次从队列中获取数据并执行打印操作，并且可以看到 "d" 的接收操作在发送操作之前完成
	go func() { // 发送
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync signal. [sender]")
			}
		}
		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}
