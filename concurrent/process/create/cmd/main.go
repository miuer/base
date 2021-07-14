package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [command]\n", os.Args[0])
		os.Exit(1)
	}

	// 从 PATH 获取外部系统命令或二进制可执行文件的绝对路径
	cmdName := os.Args[1]
	if filepath.Base(os.Args[1]) == os.Args[1] {
		if lp, err := exec.LookPath(os.Args[1]); err != nil {
			fmt.Println("look path error:", err)
			os.Exit(1)
		} else {
			cmdName = lp
		}
	}

	cmd := exec.Command(cmdName)
	outputPipe, _ := cmd.StdoutPipe()
	defer outputPipe.Close()
	outputbuf := bufio.NewReader(outputPipe)

	var output bytes.Buffer

	err := cmd.Start()
	if err != nil {
		fmt.Println("start cmd error:", err)
		os.Exit(2)
	}

	go func() {
		var buffer []byte = make([]byte, 4096)
		for {
			n, err := outputbuf.Read(buffer)
			if err != nil {
				if err == io.EOF {
					// fmt.Printf("pipe has Closed\n")
					break
				} else {
					fmt.Println("Read content failed")
				}
			}
			if n > 0 {
				output.Write(buffer[:n])
			}
		}
	}()

	processState, err := cmd.Process.Wait()
	if err != nil {
		fmt.Println("wait error:", err)
		os.Exit(3)
	}

	fmt.Println(output.String())
	fmt.Println("主进程 pid: ", os.Getpid())
	fmt.Println("子进程 pid: ", processState.Pid())
	fmt.Println("子进程 exitCode: ", processState.ExitCode())

	// 挂起主进程
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
