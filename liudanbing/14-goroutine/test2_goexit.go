package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 用 go 创建承载一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		// 如果直接写一个匿名函数，那么只有定义，并没有调用；想要调用，需要加上()，括号里面是调用函数的参数
		func() {
			defer fmt.Println("B.defer")

			// 退出当前goroutine
			runtime.Goexit() // 此时当前func会出栈，再外层的func也会出栈，所以两个defer会执行，但是A B不会打印
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()
	
	// 如何获得当前的返回值？
	// 不能flag := ...来获得
	// 因为这里的子goroutine是和主goroutine是异步的（并行的）
	// 所以需要用channel来实现
	go func(a int, b int) bool {
		fmt.Println(a, b)
		return true
	}(10, 20)

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}
}