/*
	知识点二：defer和return谁先谁后
	- returnFunc()函数先执行，returnFunc()函数执行完毕后，deferFunc()函数才会执行
*/

package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called ...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called ...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}