package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		for i := 0; i < 2; i++ {
			c <- i
		}

		close(c)
	}()
	
	// 如果c中有数据，那么range就会给出数据，如果c中没有数据，那么range就会阻塞
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main goroutine结束")
}