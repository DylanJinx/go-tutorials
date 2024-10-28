// 一个变量都说由type和value的pair构成的
package main

import (
	"fmt"
)

func main() {
	// a pair<static type: string, value: >
	var a string
	// a pair<static type: string, value: "hello">
	a = "hello"

	// allType pair<static type: interface{}, value: >
	var allType interface{}
	// allType pair<concrete type: string, value: "hello">
	allType = a

	str, _ := allType.(string) //尝试将 allType 中的值断言为 string 类型。如果断言成功，则 str 将持有 allType 的值 "hello"
	// str pari<static type: string, value: "hello">
	fmt.Println(str)
}