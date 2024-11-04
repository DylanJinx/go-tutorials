package main

import (
	"fmt"
)

func main() {
	// 第一种声明方式
	// 声明myMap1是一个map类型，key是string，value是string
	var myMap1 map[string]string

	if myMap1 == nil {
		fmt.Println("myMap1 is nil")
	} 

	myMap1 = make(map[string]string, 2)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python" // 超过容量，会重新分配内存

	fmt.Println(myMap1) //map[one:java three:python two:c++] 无序的，使用的是hashmap

	// 第二种声明方式
	myMap2 := make(map[int]string)
	// 可以不用make，直接赋值
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"
	fmt.Println(myMap2) //map[1:java 2:c++ 3:python]

	// 第三种声明方式
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "python",
	}
	fmt.Println(myMap3) //map[one:java three:python two:c++]

}