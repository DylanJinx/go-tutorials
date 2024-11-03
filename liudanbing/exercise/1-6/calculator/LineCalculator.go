package calculator

import "fmt"

const (
	ADD = "+"
	SUB = "-"
	MUL = "*"
	DIV = "/"
)

func init() {
	fmt.Println("Welcome to use the command - line calculator!")
}

// 首字母大写，表示这个函数是公开的，可以被其他包调用
func Calculate(a, b float64, operator string) (float64, bool) {
	switch operator {
	case ADD:
		return a + b, true
	case SUB:
		return a - b, true
	case MUL:
		return a * b, true
	case DIV:
		if b == 0 {
			return 0, false
		}
		return a / b, true
	// 必须有default，否则编译不通过
	default:
		return 0, false
	}
}