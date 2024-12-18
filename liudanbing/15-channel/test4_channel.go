// channel关闭后，是可以继续从channel中读取数据的，只不过读取到的数据是channel中剩余的数据。
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5)

	go func() {
		defer fmt.Println("子-goroutine结束")
		for i := 0; i < 5; i++ {
			c <- i
		}

		close(c)
	}()
	
	time.Sleep(10 * time.Second)

	for {
		// ok 为 true 时，表示channel没有关闭，为 false 时，表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main goroutine结束")
}