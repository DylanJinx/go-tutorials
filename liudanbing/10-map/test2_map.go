package main

import (
	"fmt"
)

func printMap(cityMap map[string]string) {
	// cityMap是引用传递，所以修改cityMap会影响原来的cityMap
	for key, value := range cityMap {
		fmt.Printf("key=%v value=%v\n", key, value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["USA"] = "Washington"
	cityMap["Japan"] = "Tokyo"

	// 遍历
	printMap(cityMap)

	fmt.Println("=====================================")
	// 删除
	delete(cityMap, "USA")
	printMap(cityMap)

	fmt.Println("=====================================")
	// 修改
	cityMap["China"] = "Shanghai"
	printMap(cityMap)

	fmt.Println("=====================================")
	ChangeValue(cityMap)
	printMap(cityMap)

	fmt.Println(cityMap["China"])

	// copy map
	newCityMap := make(map[string]string)
	for key, value := range cityMap {
		newCityMap[key] = value
	}
	fmt.Println(newCityMap)
}