package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x: // 尝试向c中写入数据，如果c可以写入数据，那么就会执行这个case；只要main中的子goroutine读取数据，那么这个case就会一直执行
			x = y	// 循环1: x = y = 1, y = 1 + 1 = 2; 循环2: x = y = 2, y = 2 + 2 = 4; 循环3: x = y = 4, y = 4 + 4 = 8
			y = x + y 
		case <-quit: // 尝试从quit中读取数据，如果quit中有数据，那么就会执行这个case
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c) // 不断尝试从channel中读取数据，如果channel中没有数据，就会阻塞
		}

		quit <- 0 // 当c 读出6个数据后，向quit中写入数据0，表示结束
	}()

	fibonacci(c, quit)

}