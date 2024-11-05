package main

// import "fmt"

type Book struct {
    Title string
}

// 指针接收者方法
func (b *Book) ModifyTitle(newTitle string) {
    b.Title = newTitle
}

func main() {
    var b Book = Book{"Go Programming"}  // 值类型的 Book
    var p *Book = &b                    // 指针类型的 Book

    b.ModifyTitle("Updated Go")  // 编译错误：值类型不能调用指针接收者方法
    p.ModifyTitle("Updated Go")  // 正常：指针类型调用指针接收者方法
}
