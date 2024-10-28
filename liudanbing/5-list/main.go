package main

import (
	"fmt"
	"mylocalproject/lib1" // 导入包必须要使用里面的函数，否则会报错
	"mylocalproject/lib2"
)

func main() {
	fmt.Println("main() ...")
	lib1.Test()
	lib2.Test()
}