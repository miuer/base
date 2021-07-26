package main

import (
	"fmt"
	"time"
)

func main() {
	sliceSleep()
}

// for 使用多个 G 对同一函数进行封装，在 for 循环结束之后并发执行 G
// case 1 打印的结果多多为 'e', 其他仅有少数是其他字符
// case 2 Sleep 让程序串行执行，结果为 'a, b, c, d, e'
// case 3 通过副本传递变量，达到预期效果
// case 4 ... ？ ...
func sliceSleep() {
	names := []string{"a", "b", "c", "d", "e"}

	// // case 1 --- most of them are "e"
	// for _, name := range names {
	// 	go func() {
	// 		fmt.Println(name)
	// 	}()
	// }
	// time.Sleep(time.Millisecond)

	// // case 1 --- a, b, c, d, e
	// for _, name := range names {
	// 	go func() {
	// 		fmt.Println(name)
	// 	}()
	// 	time.Sleep(time.Millisecond)
	// }

	// case 3 --- random "abcde"
	for _, name := range names {
		go func(str string) {
			fmt.Println(str)
		}(name)
	}
	time.Sleep(time.Millisecond)

	// // case 4 --- random "abcde"
	// for _, name := range names {
	// 	go fmt.Println(name)
	// }
	// time.Sleep(time.Millisecond)

}
