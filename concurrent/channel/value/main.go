package main

import (
	"fmt"
	"time"
)

// ch 传递的值类型与函数一样，均是原值的副本
// 通过 strChan 和 mapIntChan 的对比就可以看出
// 对于复合类型 struct 需要根据其含有的类型进行
func main() {
	structChan()
}

// 以 string 为代表的值类型无法通过 channel 对原数据进行修改，说明 channel 的传值是通过拷贝副本实现的
func strChan() {
	syncChan := make(chan struct{}, 2)
	strChan := make(chan string, 1)

	go func() {
		var count int
		for {
			if str, ok := <-strChan; ok {
				count++
				str = "hello"
				fmt.Printf("receiver %d : %v\n", count, str)

			} else {
				break
			}
		}
		fmt.Println("stop")
		syncChan <- struct{}{}
	}()

	go func() {
		var str string
		for i := 0; i < 5; i++ {
			strChan <- str
			time.Sleep(time.Microsecond)
			fmt.Printf("sender %d : %v\n", i+1, str)
		}

		close(strChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}

// 以 pointer 为代表的引用类型，通过对 channel 传递值的修改对原数据产生影响
func pointerChan() {
	syncChan := make(chan struct{}, 2)
	strChan := make(chan *string, 1)

	go func() {
		var count int
		for {
			if str, ok := <-strChan; ok {
				count++
				temp := fmt.Sprintf("hello %d", count)
				str = &temp
				fmt.Printf("receiver %d : %v\n", count, *str)

			} else {
				break
			}
		}
		fmt.Println("stop")
		syncChan <- struct{}{}
	}()

	go func() {
		str := new(string)
		for i := 0; i < 5; i++ {
			strChan <- str
			time.Sleep(time.Microsecond)
			fmt.Printf("sender %d : %v\n", i+1, *str)
		}

		close(strChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}

type Counter struct {
	count int
	desc  *string
}

// strcut 是值类型的集合，内部的引用类型会被转化为值类型
func structChan() {
	syncChan := make(chan struct{}, 2)
	structChan := make(chan Counter, 1)

	go func() {
		var count int
		var hello string = "hello"
		for {
			if str, ok := <-structChan; ok {
				count++
				str.count = count
				str.desc = &hello
				fmt.Printf("receiver %d : %v\n", count, *str.desc)

			} else {
				break
			}
		}
		fmt.Println("stop")
		syncChan <- struct{}{}
	}()

	go func() {
		desc := ""
		var str Counter = Counter{
			count: 0,
			desc:  &desc,
		}
		for i := 0; i < 5; i++ {
			structChan <- str
			time.Sleep(time.Microsecond)
			fmt.Printf("sender %d : %v\n", i+1, *str.desc)
		}

		close(structChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}

func mapIntChan() {
	var mapIntChan = make(chan map[string]int, 1)

	syncChan := make(chan struct{}, 2)
	go func() {
		var count int
		for {
			if elem, ok := <-mapIntChan; ok {
				count++
				elem["count"]++
				fmt.Printf("receiver %d : %v\n", count, elem)

			} else {
				break
			}
		}
		fmt.Println("Stopped")
		syncChan <- struct{}{}
	}()
	go func() {
		intMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapIntChan <- intMap
			time.Sleep(time.Millisecond)
			fmt.Printf("sender %d : %v \n", i+1, intMap)
		}
		close(mapIntChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

// struct 值类型，通过 channel 传递无法修改原值，但可以将其转变为指针实现对原数据的更新
func mapStructMap() {
	var mapStructChan = make(chan map[string]Counter, 1)

	syncChan := make(chan struct{}, 2)
	var hello string = "hello world"
	go func() {
		for {
			if elem, ok := <-mapStructChan; ok {
				counter := elem["count"]
				counter.count++
				counter.desc = &hello
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() {
		countMap := map[string]Counter{
			"count": {},
		}
		for i := 0; i < 5; i++ {
			mapStructChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapStructChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func mapStrPoMap() {
	var mapStructChan = make(chan map[string]*Counter, 1)

	syncChan := make(chan struct{}, 2)
	var hello string = "hello world"
	go func() {
		for {
			if elem, ok := <-mapStructChan; ok {
				counter := elem["count"]
				counter.count++
				counter.desc = &hello
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() {
		countMap := map[string]*Counter{
			"count": {},
		}
		for i := 0; i < 5; i++ {
			mapStructChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap["count"].count)
		}
		close(mapStructChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

type Counters struct {
	c      Counter
	number int
}

func compositeMap() {
	var mapStructChan = make(chan map[string]*Counters, 1)

	syncChan := make(chan struct{}, 2)
	var hello string = "hello world"
	go func() {
		for {
			if elem, ok := <-mapStructChan; ok {
				counter := elem["count"]
				counter.c.count++
				counter.c.desc = &hello
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() {
		countMap := map[string]*Counters{
			"count": {},
		}
		for i := 0; i < 5; i++ {
			mapStructChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap["count"].c.count)
		}
		close(mapStructChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
