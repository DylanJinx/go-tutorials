package main // 主包

// import "fmt" // 引入 fmt 包
// import "time"

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	
	time.Sleep(1 * time.Second)

	fmt.Println("Hello, World!")
}