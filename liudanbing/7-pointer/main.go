package main

import (
	"fmt"
)

func changeValue(p *int) {
	*p = 100
}

func swap1(a int, b int) {
	var temp int 
	temp = a
	a = b
	b = temp
}

func swap2(pa *int, pb *int) {
	var temp int 
	temp = *pa // temp = main::a
	*pa = *pb // main::a = main::b
	*pb = temp // main::b = temp
}

func main() {
	var a int = 1

	changeValue(&a)

	fmt.Println("a = ", a)

	var b int = 200
	swap1(a, b)
	fmt.Println("a = ", a, "b = ", b)

	swap2(&a, &b)
	fmt.Println("a = ", a, "b = ", b)

	var p *int = &a
	fmt.Println("p = ", p)
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Printf("type p = %T\n", p)
}