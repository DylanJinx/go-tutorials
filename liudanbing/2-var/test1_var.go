package main
/*
	四种变量的声明方式
*/

import "fmt"

// 全局变量声明，方法一、方法二、方法三是可以的
var g1 int
var g2 int = 100
var g3 = 100

// g4 := 100 // 编译错误，全局变量不能使用 := 方式声明，:= 只能用于函数内部
// v := "dd"


func main() {
	// 方法一：指定变量类型，声明后若不赋值，使用默认值，int默认值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var b2 string = "b2"
	fmt.Printf("b2 = %s, type of b2 = %T\n", b2, b2)

	// 方法三：在初始化的时候，可以省去数据类型，自动匹配类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var c2 = "c2"
	fmt.Printf("c2 = %s, type of c2 = %T\n", c2, c2)

	// 方法四：省略var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
	d := 100 // 既初始化一个值，同时自动匹配类型
	fmt.Println("e = ", d)
	fmt.Printf("type of d = %T\n", d)

	d2 := "d2"
	fmt.Printf("d2 = %s, type of d2 = %T\n", d2, d2)

	d3 := 3.14
	fmt.Println("d3 = ", d3)
	fmt.Printf("d3 = %f, type of d3 = %T\n", d3, d3)

	// 打印全局变量
	fmt.Println("g1 = ", g1)
	fmt.Println("g2 = ", g2)
	fmt.Println("g3 = ", g3)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)

	var kk, ll = 100, "ll"
	fmt.Println("kk = ", kk, "ll = ", ll)
	fmt.Printf("type of kk = %T, type of ll = %T\n", kk, ll)

	// 多行的多个变量声明
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj)
}





