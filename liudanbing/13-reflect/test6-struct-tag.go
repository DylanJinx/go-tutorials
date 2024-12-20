package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"my name"`
	Sex string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem() // 注：t := reflect.TypeOf(str), Elem()方法是获取指针指向的元素类型
	
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagInfo, ", doc: ", tagDoc)
	}
}

func main() {
	var re resume
	findTag(&re) // 注：findTag(re)
}