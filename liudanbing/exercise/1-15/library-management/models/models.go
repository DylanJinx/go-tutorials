package models

import "fmt"

type Category int

const (
	Fiction Category = iota // Fiction是Category类型，值为0
	NonFiction // 
	Science
	Biography
	Computer
)

// CategoryNames 存储类别名称，便于显示
var CategoryNames = []string{
	"Fiction",
	"NonFiction",
	"Science",
	"Biography",
	"Computer",
}

// 为Category类型添加String方法（fmt.Stringer接口），便于显示
func (c Category) String() string {
	if int(c) < len(CategoryNames) {
		return CategoryNames[c]
	}
	return "Unknown"
}

type Printable interface {
	PrintDetails()
}

type Book struct {
	ID	   int
	Title  string
	Author string
	Category Category
	Price float64
}

func (b Book) PrintDetails() {
    fmt.Printf("ID: %d, 书名: %s, 作者: %s, 类别: %s, 价格: %.2f\n",
        b.ID, b.Title, b.Author, b.Category, b.Price)
}