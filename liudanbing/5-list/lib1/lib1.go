package lib1

import "fmt"

// 首字母大写，表示该函数是公开的
func Test() {
	fmt.Println("lib1. Test() ...")
}

func init() {
	fmt.Println("lib1. init() ...")
}