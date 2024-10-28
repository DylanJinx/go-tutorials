package main

import "fmt"

// 声明一种新的数据类型 myInt 为 int 的别名
type myInt int

// 定义一个结构体
type Book struct {
	title string
	auth string
}

func changeBook(book Book) {
	// 传递一个book的副本
	book.auth = "li4"
}

func changeBook2(book *Book) {
	// 传递一个book的指针
	book.auth = "li4"
}

func main() {
	var a myInt = 10
	fmt.Println("a = ", a)
	fmt.Printf("a type is %T\n", a) // a type is main.myInt

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("book1 = %v\n", book1) // book1 = {Golang zhang3}

	// 传递结构体的副本
	changeBook(book1)
	fmt.Printf("book1 = %v\n", book1) // book1 = {Golang zhang3}

	// 传递结构体的指针
	changeBook2(&book1)
	fmt.Printf("book1 = %v\n", book1) // book1 = {Golang li4}
}