package main

import "fmt"

// interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...", arg)
	fmt.Printf("Address of arg is %p\n", &arg)

	// interface{} 如何区分 此时引用的底层数据类型到底是什么？
	// interface{} 可以通过类型断言来判断底层数据类型是什么
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}
	fmt.Printf("address of book is %p\n", &book)
	myFunc(book) // myFunc is called... {Golang}

	fmt.Println("=====================================")
	a := 100
	fmt.Printf("address of a is %p\n", &a)
	myFunc(a) // myFunc is called... 100

	fmt.Println("=====================================")
	b := "abc"
	fmt.Printf("address of b is %p\n", &b)
	myFunc(b) // myFunc is called... abc

	fmt.Println("=====================================")
	c := 3.14
	fmt.Printf("address of c is %p\n", &c)
	myFunc(c) // myFunc is called... 3.14

}