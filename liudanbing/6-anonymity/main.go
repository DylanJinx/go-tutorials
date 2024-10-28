package main

import (
	"fmt"
	// "mylocalproject/lib1" // 如果导入但是没有使用这个包里的函数，会报错
	// 如果只想调用包里面的init()函数，不想调用其他函数，可以使用下划线
	_ "mylocalproject/lib1"      // 使用下划线，表示匿名导入，不会报错，但是这种导入方式不能调用包里的函数，只能运行init()函数
	mylib2 "mylocalproject/lib2" // 可以给包起别名
	. "mylocalproject/lib3"      // 使用点号，表示将lib3中的函数全部写入到main包中，可以不用写包名，直接调用包里的函数
)

func main() {
	fmt.Println("main() ...")
	mylib2.Test()
	Test()
}