package main

import "fmt"

type Modifiable interface {
    modify()
}

type Book struct {
    Title string
}

// 指针接收者方法
func (b *Book) modify() {
    fmt.Println("Book modified:", b.Title)
}

func main() {
    b := Book{Title: "Go Programming"}

    // 通过值类型调用指针接收者方法，编译错误
    var m Modifiable = b // 错误: Book 类型不能自动实现指针接收者方法 正确写法: var m Modifiable = &b
	m.modify()

    // 必须通过指针类型调用
    // var m2 Modifiable = &b
    // m2.modify() // 输出: Book modified: Go Programming
}
