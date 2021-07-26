package main

import (
	"fmt"
	"time"
)

func main() {
	baseSleep()
}

func baseSleep() {
	go fmt.Println("hello world")

	time.Sleep(time.Millisecond)

}

// func printSleep() {
// 	var name string = "miuer"

// 	go func() {
// 		fmt.Println(name)
// 	}()

// 	time.Sleep(time.Millisecond)
// 	name = "richard"

// }
