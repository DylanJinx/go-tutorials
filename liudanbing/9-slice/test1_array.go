package main

import "fmt"

func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
}

func main() {
	// 固定长度的数组
	var myArray1 [10]int

	myArray2 := [8]int{1, 2, 3, 4}
	
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 type is %T\n", myArray1)
	fmt.Printf("myArray2 type is %T\n", myArray2)

	myArray3 := [4]int{11, 22, 33, 44}
	printArray(myArray3) // 值传递

	
}