package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main goroutine结束")
	// 定义一个有缓冲的channel，传递int类型的数据
	c := make(chan int, 3)

	// len 代表channel中的当前元素个数，cap 代表channel的容量
	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		// defer fmt.Println("子-goroutine结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子-goroutine写入：", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()

	// 先让子goroutine向channel中写入数据
	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("main goroutine读取：", num)
	}

	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	time.Sleep(2 * time.Second) // 如果不sleep，可能会导致main goroutine结束，子goroutine还没结束
}

/* 输出一：
len(c) =  0 cap(c) =  3
子-goroutine写入： 0 len(c) =  1 cap(c) =  3
子-goroutine写入： 1 len(c) =  2 cap(c) =  3
子-goroutine写入： 2 len(c) =  3 cap(c) =  3
main goroutine读取： 0
main goroutine读取： 1
main goroutine读取： 2
main goroutine读取： 3
len(c) =  0 cap(c) =  3
子-goroutine写入： 3 len(c) =  0 cap(c) =  3
main goroutine结束
*/

/* 输出二：
len(c) =  0 cap(c) =  3
子-goroutine写入： 0 len(c) =  1 cap(c) =  3
子-goroutine写入： 1 len(c) =  2 cap(c) =  3
子-goroutine写入： 2 len(c) =  3 cap(c) =  3
main goroutine读取： 0
main goroutine读取： 1
main goroutine读取： 2
main goroutine读取： 3
len(c) =  0 cap(c) =  3
子-goroutine写入： 3 len(c) =  3 cap(c) =  3
main goroutine结束
*/
