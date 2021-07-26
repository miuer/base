package main

import (
	"fmt"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go consumer(strChan, syncChan1, syncChan2)
	go producer(strChan, syncChan1, syncChan2)

	<-syncChan2
	<-syncChan2
}

//	strChan <-chan 消费者获取发送通道中的数据
//	syncChan1 <-chan 接收生产者传来的同步信号
//	syncChan2 chan<- 向主进程传递同步信号
func consumer(strChan <-chan string,
	syncChan1 <-chan struct{},
	syncChan2 chan<- struct{}) {
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
}

//	strChan chan<- 生产者向接收通道传递数据
//	syncChan1 chan<- 向消费者传递同步信号
//	syncChan2 chan<- 向主进程传递同步信号
func producer(strChan chan<- string,
	syncChan1 chan<- struct{},
	syncChan2 chan<- struct{}) {
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
}

// 在语法级别规避了对参数 c 的错误操作
// 适用于接口声明
// func notify(c chan<- os.Signal, sig ...os.Signal) {}

// 对返回的通道类型进行约束，只允许从中获取数据，而不能向其发送数据
// 适用于函数和结构体
// func Notify(sig ...os.Signal) <-chan os.Signal
