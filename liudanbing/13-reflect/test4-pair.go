package main

import (
	"fmt"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book")
}

func main() {
	// b: pair<static type: *Book, value: 指向Book实例的指针>
	b := &Book{}
	// b: Book{} // b pair<static type: Book, value: Book实例>

	// r: pair<static type: , value: >
	var r Reader
	// r: pair<concrete type: *Book, value: 指向Book实例的指针>
	r = b
	r.ReadBook() // Read a book

	// w: pair<static type: , value: >
	var w Writer
	// w: pair<concrete type: *Book, value: 指向Book实例的指针>
	w = r.(Writer) // 为什么这里可以断言成功？因为 r 的值是 *Book 类型，而 *Book 类型实现了 Writer 接口
	w.WriteBook() // Write a book

}