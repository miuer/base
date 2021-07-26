package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	ticker := time.NewTicker(time.Second)
	done := make(chan struct{})
	go func(chan struct{}) {
	Loop:
		for range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			case _, ok := <-done:
				if ok {
					break Loop
				}
			}
		}
		fmt.Println("End. [sender]")
	}(done)

	var sum int
	for e := range intChan {
		fmt.Printf("Received: %v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("Got: %v\n", sum)
			break
		}
	}

	done <- struct{}{}

	fmt.Println("End. [receiver]")
}
