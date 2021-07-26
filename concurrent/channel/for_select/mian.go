package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	forSelectChanFunc()
}

func forChanFunc() {
	var strChan = make(chan string, 2)

	syncChan1 := make(chan struct{}, 1)
	defer close(syncChan1)
	syncChan2 := make(chan struct{}, 2)
	defer close(syncChan2)

	go func() { // 接收
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second... [receiver]")
		// for {
		// 	if elem, ok := <-strChan; ok {
		// 		fmt.Println("Received:", elem, "[receiver]")
		// 	} else {
		// 		break
		// 	}
		// }

		// for ... range 依次通道缓冲队列中的数据
		for elem := range strChan {
			fmt.Println("Received:", elem, "[receiver]")
		}

		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()
	go func() { // 发送
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[sender]")
			if elem == "b" {
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

// select ... case ... 遵循从左到右，自上而下的顺序对 case 语句的表达式进行计算
// 当有多个 case 满足执行条件时，会使用一个伪随机的算法选择其中的一个进行执行(不包括 default)
// 当所有的case都不满足时，程序会在此阻塞，直到有一个 case 可用为止（应使用超时或者default）
func selectChanFunc() {
	ch := make(chan int, 1)
	defer close(ch)

	rand.Seed(time.Now().UnixNano())
	ch <- rand.Intn(5)

	//  当 number 的值为 2 时，打印 "end" 并使用 break 跳出 select
	select {
	case number := <-ch:
		if number == 2 {
			fmt.Println("end")
			break
		}
		fmt.Println("the number received is: ", number)
	default:
		fmt.Println("non number")
	}

}

func forSelectChanFunc() {
	ch := make(chan int, 5)
	defer close(ch)

	for elem := range []int{0, 1, 2, 3, 4} {
		ch <- elem
	}

	// 当 for 和 select 一起使用的时候，break 跳出的 select，配合 label 可以跳出 for 循环
Loop:
	for {
		select {
		case number := <-ch:
			if number == 2 {
				fmt.Println("end")
				break Loop
			}
			fmt.Println("the number received is: ", number)
		default:
			fmt.Println("non number")
		}
	}
}
