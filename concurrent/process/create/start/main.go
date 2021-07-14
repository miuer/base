package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

/*
	1.命令行参数捕获和重构
	2.创建子进程执行外部命令
	3.输出进程信息进行观察
	4.挂起程序便于验证
	5.创建测试程序hello
*/
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

	procAttr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("look path error:", err)
		os.Exit(1)
	}

	start := time.Now()

	// StartProcess -- 创建子进程
	// name -- 外部系统命令或二进制可执行文件
	process, err := os.StartProcess(cmdName, []string{cwd}, procAttr)
	if err != nil {
		fmt.Println("start process error:", err)
		os.Exit(2)
	}

	// Wait -- 捕获子进程
	processState, err := process.Wait()
	if err != nil {
		fmt.Println("wait error:", err)
		os.Exit(3)
	}

	// output
	fmt.Println()
	fmt.Println("real", time.Now().Sub(start))
	fmt.Println("user", processState.UserTime())
	fmt.Println("system", processState.SystemTime())

	fmt.Println("主进程 pid: ", os.Getpid())
	fmt.Println("子进程 pid: ", processState.Pid())
	fmt.Println("子进程 exitCode: ", processState.ExitCode())

	// 挂起主进程
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
