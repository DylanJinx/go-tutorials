package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("子-goroutine结束")
		for i := 0; i < 5; i++ {
			c <- i
		}

		close(c)
		/*
		如果不关闭channel，会导致死锁，输出如下：
		0
		1
		2
		3
		4
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan receive]:
		main.main()
				/Volumes/T7 Shield/Code_apps/MacBook_github/Go_Projects/tutorials/liudanbing/15-channel/test3_channel.go:18 +0xc0
		exit status 2
		*/
	}()

	// 如果不关闭channel，那么会导致for循环一直阻塞
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

/*
两种情况的输出：
➜  15-channel git:(main) ✗ go run test3_channel.go
0
1
2
3
4
main goroutine结束
➜  15-channel git:(main) ✗ go run test3_channel.go
0
1
2
3
子-goroutine结束
4
main goroutine结束
*/
