package main

import (
	"fmt"
	"time"
)

// 主goroutine是main
func main() {
	defer fmt.Println("main goroutine 结束...")

	go func() {
		defer fmt.Println("子-goroutine结束")
		fmt.Println("子-goroutine开始...")
	}()
	
	// 如果不加这个，main 函数在启动子 goroutine 后立即结束，导致程序在子 goroutine 有机会执行之前就已经退出了。打印结果是：main goroutine 结束...
	time.Sleep(1 * time.Second) // 1秒后main goroutine结束，打印结果是：子-goroutine开始... 子-goroutine结束 main goroutine 结束...

}

