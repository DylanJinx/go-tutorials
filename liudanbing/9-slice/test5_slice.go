package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3} // len = 3, cap = 3, slice = [1 2 3]
	// 传递的是引用
	s1 := s[0:2] // [0, 2), len = 2, cap = 3, slice = [1 2]
	fmt.Println(s1)

	s1[0] = 100
	fmt.Println(s) // [100 2 3]
	fmt.Println(s1) // [100 2]

	//copy
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2) // [100 2 3]
	s2[0] = 200
	fmt.Println(s) // [100 2 3]
	fmt.Println(s2) // [200 2 3]

	// 创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numbers)

	// 打印原始切片
	fmt.Println("numbers ==", numbers) // numbers == [0 1 2 3 4 5 6 7 8]

	// 打印子切片从索引1(包含) 到索引4(不包含)
	fmt.Println("numbers[1:4] ==", numbers[1:4]) // numbers[1:4] == [1 2 3]

	// 默认下限为0
	fmt.Println("numbers[:3] ==", numbers[:3]) // numbers[:3] == [0 1 2]

	// 默认上限为len(s)
	fmt.Println("numbers[4:] ==", numbers[4:]) // numbers[4:] == [4 5 6 7 8]

	numbers1 := make([]int, 0, 5)
	fmt.Println(numbers1)

	// 打印子切片从索引0(包含) 到索引2(不包含)
	numbers2 := numbers[:2]
	fmt.Println(numbers2)

	// 打印子切片从索引2(包含) 到索引5(不包含)
	numbers3 := numbers[2:5]
	fmt.Println(numbers3)
}