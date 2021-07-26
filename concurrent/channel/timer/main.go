package main

import (
	"fmt"
	"log"
	"time"
)

// 适用于超时触发器
func main() {
	intChan := make(chan int, 1)
	go func() {
		for i := 1; i < 5; i++ {

			d, err := time.ParseDuration(fmt.Sprintf("%dms", i))
			if err != nil {
				log.Fatalln("parse duration error: ", err)
			}
			time.Sleep(d)
			intChan <- i
		}
		close(intChan)
	}()

	timeout := 2 * time.Millisecond
	var timer *time.Timer
Loop:
	for {
		// 在 for 循环中应当避免使用注释中的形式声明计时器
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}
		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Printf("Received: %v\n", e)

		// case <-time.After(2 * time.Millisecond):
		// case <-time.NewTimer(2 * time.Millisecond).C:
		case <-timer.C:
			fmt.Println("Timeout!")
			break Loop
		}
	}

}
