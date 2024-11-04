package main

import (
	"fmt"
)

func main() {
	// 1. 声明slice1是一个切片，并且初始化，默认值是1,2,3，长度是3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1)

	// 2. 声明slice2是一个切片，但是没有初始化，nil，长度是0
	var slice2 []int
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)
	// slice2[0] = 1 // 这里会引发一个panic，因为没有初始化，需要make初始化
	slice2 = make([]int, 3) // 使用make函数开辟3个容量
	slice2[0] = 100
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	// 3. 声明slice3是一个切片，同时给slice3分配3个空间，初始化值为0
	var slice3 []int = make([]int, 3)
	fmt.Printf("len = %d, slice3 = %v\n", len(slice3), slice3)

	// 4. 声明slice4是一个切片，同时给slice4分配3个空间，初始化值为0，通过:=自动推导类型
	slice4 := make([]int, 3)
	fmt.Printf("len = %d, slice4 = %v\n", len(slice4), slice4)

	var slice []int 
	if slice == nil {
		fmt.Println("slice is nil")
	} else {
		fmt.Println("slice is not nil")
	}
}