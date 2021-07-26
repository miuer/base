package main

import "fmt"

// channel 的关闭应该在发送端实现，应该避免在接收端关闭通道，往往和 for 循环和 time 协作完成
// channel 的关闭遵循谁创建谁关闭
func main() {

}

func closeChanFunc() {
	var strChan = make(chan string, 3)

	// 谁创建谁关闭
	syncChan1 := make(chan struct{}, 1)
	defer close(syncChan1)
	syncChan2 := make(chan struct{}, 2)
	defer close(syncChan2)

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
	go func() { // 发送
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync signal. [sender]")
			}
		}

		// 通道的关闭应该在发送端关闭
		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}

func elegantClose() {

}

// channel 关闭两次会 panic

// 对已经关闭的 channel 进行写操作会 panic
/*	runtime_chan.go_chansend()
	if c.closed != 0 {
 		unlock(&c.lock)
		panic(plainError("send on closed channel"))
	}
*/

// 使用 close 关闭未初始化的 chanel 会 panic

// 读已关闭的chan会一直能读到值
/* runtime_chan.go_chanrecv()
if c.closed != 0 && c.qcount == 0 {
	if raceenabled {
		raceacquire(c.raceaddr())
	}
	unlock(&c.lock)
	if ep != nil {
		typedmemclr(c.elemtype, ep)
	}
	return true, false
}
*/
