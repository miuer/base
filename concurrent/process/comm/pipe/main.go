package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

/*
	基于 pipe 实现 cmd 子进程之间的通信
	父子之间的通信在基于 cmd 创建进程时已实现
	匿名管道适用于父子进程以及同祖先的子进程
*/
func main() {
	cmd1 := exec.Command("ls", "-l")
	cmd2 := exec.Command("grep", "main")

	var outputBuf1 bytes.Buffer
	var outputBuf2 bytes.Buffer

	cmd1.Stdout = &outputBuf1

	// 将前一个进程的输出作为后一个进程的输入进行关联
	cmd2.Stdin = &outputBuf1
	cmd2.Stdout = &outputBuf2

	if err := cmd1.Start(); err != nil {
		fmt.Println(err)
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Println(err)
		return
	}

	if err := cmd2.Start(); err != nil {
		fmt.Println(err)
		return
	}

	if err := cmd2.Wait(); err != nil {
		fmt.Println(err)
		return
	}

	// output
	fmt.Printf("%s", outputBuf2.Bytes())
}
