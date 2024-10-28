package main

import (
	"fmt"
)

func printArray(myArray []int) {
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组，切片slice
	fmt.Printf("myArray type is %T\n", myArray) // result: myArray type is []int

	printArray(myArray) // 切片传递的时候是引用传递
	// 切片传递的是整体数组的指针

	fmt.Println("myArray[0] = ", myArray[0]) // result: myArray[0] = 100
}