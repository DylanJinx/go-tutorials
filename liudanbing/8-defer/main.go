package main

import (
	"fmt"
	"time"
)

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func defer_call() {
	defer func1()
	defer func2()
	defer func3()
}

// defer表示一个函数在执行最后、在结束之前的一种机制，它会在它所在的函数体在结束之前执行，类似于C++中的析构函数
func main1() {
	startTime := time.Now()

	// 写入defer关键字
	defer fmt.Println("main end1")
	// defer是采用压栈的方式，先进后出
	defer fmt.Println("main end2")

	defer_call() // C B A

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("main() duration = ", duration)
}