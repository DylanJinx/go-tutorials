package main

import "fmt"

// 主goroutine是main
func main() {
	// 定义一个无缓冲的channel，传递int类型的数据
	c := make(chan int)

	go func() {
		defer fmt.Println("子-goroutine结束") // 永远是在 num := <- c 之后执行
		fmt.Println("子-goroutine开始...")

		c <- 666 // 将666发送到channel c
	}()
	
	// 之前想用flag := go func()...来获得返回值,但是不行,因为子goroutine是异步的,所以需要用channel来实现
	num := <- c // 从channel c接收数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}

