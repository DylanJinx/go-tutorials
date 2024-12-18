package main

import "fmt"

func main() {
	// 定义一个有缓冲的channel，传递int类型的数据
	c := make(chan int, 3)

	// len 代表channel中的当前元素个数，cap 代表channel的容量
	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子-goroutine结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子-goroutine写入：", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("main goroutine读取：", num)
	}
}
