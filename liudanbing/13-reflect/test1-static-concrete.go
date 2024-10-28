package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var y interface{}
	y = x 
	fmt.Println(y) // 10
	value, ok := y.(int)
	if ok {
		fmt.Printf("value type is %T\n", value) // value type is int
	}

	var z string = "hi"
	y = z
	fmt.Println(y) // hi
	value2, ok2 := y.(string)
	if ok2 {
		fmt.Printf("value2 type is %T\n", value2) // value2 type is string
	}
}