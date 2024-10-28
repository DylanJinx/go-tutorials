package main

import "fmt"

// const 来定义枚举类型
const (
	BEIJING = 1
	SHANGHAI = 2
	SHENZHEN = 3
)

// 可以在const()添加一个关键字iota，每行的iota都会自动加1，第一行的iota默认为0
// iota 只能配合const()使用，iota只有在const进行累加效果
const (
	a = iota // 0
	b		 // iota = 1
	c		 // iota = 2
)

// 表达式中的iota自增值
const (
	a1 = iota * 2 // 0 * 2
	b1			 // 1 * 2
	c1			 // 2 * 2
)

// 更换表达式
const (
	a2, b2 = iota + 1, iota + 2 // iota = 0, a2 = 0 + 1, b2 = 0 + 2
	c2, d2                      // iota = 1, c2 = 1 + 1, d2 = 1 + 2
	e2, f2                      // iota = 2, e2 = 2 + 1, f2 = 2 + 2

	g2, h2 = iota * 2, iota * 3 // iota = 3, g2 = 3 * 2, h2 = 3 * 3
	i2, j2                      // iota = 4, i2 = 4 * 2, j2 = 4 * 3
)

func main() {
	const length int = 10
	fmt.Println("length = ", length)
	fmt.Printf("type of length = %T\n", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

	fmt.Println("a1 = ", a1)
	fmt.Println("b1 = ", b1)
	fmt.Println("c1 = ", c1)

	fmt.Println("a2 = ", a2, "b2 = ", b2)
	fmt.Println("c2 = ", c2, "d2 = ", d2)
	fmt.Println("e2 = ", e2, "f2 = ", f2)
	fmt.Println("g2 = ", g2, "h2 = ", h2)
	fmt.Println("i2 = ", i2, "j2 = ", j2)
}