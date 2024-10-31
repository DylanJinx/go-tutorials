package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println("a = ", a)

	xx,yy := 100,200
	fmt.Println("xx = ", xx, "yy = ", yy)

	var (
		v1 int = 100
		v2 string = "v2"
		v3 bool = true
		v4 float32 = 3.14
	)

	fmt.Println("v1 = ", v1, "v2 = ", v2, "v3 = ", v3, "v4 = ", v4)
}