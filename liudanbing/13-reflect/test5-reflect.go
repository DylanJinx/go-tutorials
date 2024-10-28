package main

import (
    "fmt"
    "reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

type User struct {
	Id int
	UserName string
	Age int
}

func (this User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType: ", inputType) // main.User
	fmt.Println("inputType: ", inputType.Name()) // User

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue: ", inputValue) // {1 Tom 20}

	// 通过type 获取里面的字段
	// 1. 获取interface的reflect.Type，通过Type得到NumField，进行遍历
	// 2. 得到的field就是一个数据类型
	// 3. 通过field有一个Interface()方法得到对应的value
	for i:= 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i) // 获取字段
		value := inputValue.Field(i).Interface() // 获取字段的值
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type 获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)

	user := User{1, "Tom", 20}

	fmt.Println("user name: ", user.UserName)
	DoFiledAndMethod(user)

}