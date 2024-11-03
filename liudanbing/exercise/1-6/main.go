package main

import (
	"fmt"
	"commandLineCalculator/calculator"
)

func main() {
	var a, b float64
	var operator string

	fmt.Printf("Please input the first number: ")
	fmt.Scanln(&a) // 这里传递的是 a 的地址，Scanln 可以修改 a 的值，如果没有 &，即 fmt.Scanln(a)，Scanln 会试图将输入值写回到一个副本中，而不是原始变量 a。
	fmt.Printf("Please input the operator(+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Printf("Please input the second number: ")
	fmt.Scanln(&b)

	result, ok := calculator.Calculate(a, b, operator)

	if ok {
		fmt.Printf("The result is: %f\n", result)
	} else {
		fmt.Println("Calculation failed. Please check the input.")
	}
}