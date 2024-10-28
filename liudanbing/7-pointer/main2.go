package main

import (
	"fmt"
	"time"
)

func changeValue0(p0 int) {
	p0 = 10
	fmt.Printf("type of p0 = %T\n", 1)
}

func changeValue1(p1 *int) {
	fmt.Println("p1 = ", p1) // 是传入的a1的地址
	fmt.Printf("Address of variable pointed by p1 = %p\n", p1) // 是传入的a1的地址
	fmt.Printf("Address of p1 itself, &p1 = %p\n", &p1) // 是p1自己的地址
	*p1 = 10
	fmt.Printf("type of p1 = %T\n", p1) // 是p1的类型
}

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	startTime := time.Now()

	var a int = 1
	changeValue0(a)
	fmt.Println("a = ", a)

	var a1 int = 1
	fmt.Printf("Address of a1 = %p\n", &a1)
	changeValue1(&a1)
	fmt.Println("a1 = ", a1)

	// swap
	var x int = 100
	var y int = 200
	swap(&x, &y)
	fmt.Println("x = ", x, ", y = ", y)

	// 一级指针
	fmt.Println("一级指针")
	var one_p *int
	one_p = &a // p等于a的地址
	fmt.Println("&a = ", &a)
	fmt.Println("one_p = ", one_p)
	
	// 二级指针
	fmt.Println("二级指针")
	var two_p **int
	two_p = &one_p
	fmt.Println("&one_p = ", &one_p)
	fmt.Println("two_p = ", two_p)

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("程序执行时间：", duration)
}