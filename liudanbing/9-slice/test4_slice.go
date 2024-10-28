// 容量的追加
// 切片的长度和容量不同，长度表示左指针到右指针之间的距离，容量表示左指针到底层数组之间的距离
package main

import (
	"fmt"
)

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) //len = 3, cap = 5, slice = [0 0 0]

	// 追加1个元素
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) //len = 4, cap = 5, slice = [0 0 0 1]

	// 追加多个元素
	numbers = append(numbers, 2, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) //len = 6, cap = 10, slice = [0 0 0 1 2 3]
	// cap = 10，因为容量不够，所以会重新分配内存，容量变为原来的2倍

	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2) //len = 3, cap = 3, slice = [0 0 0]
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2) //len = 4, cap = 6, slice = [0 0 0 1]
}