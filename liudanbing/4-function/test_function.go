package main

import "fmt"

// func 函数名(参数名 参数类型, ...) 返回值类型 {
func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 多返回值，但都是匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	d := 200

	return c, d
}

// 多返回值，但都是有名字的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----------foo3----------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 1000
	r2 = 2000

	return r1, r2
}

// 如果返回值是同一个类型的，可以一起写
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----------foo4----------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	r1 = 1000
	r2 = 2000

	return r1, r2
}

func main() {
	c := foo1("abc", 100)
	fmt.Println("c = ", c)

	e, f := foo2("def", 200)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)

	g, h := foo3("ghi", 300)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)

	i, j := foo4("jkl", 400)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
}