package main

import "fmt"

type Person struct {
	Name string
}

// 指针接收者方法
func (p *Person) ChangeName(newName string) {
	p.Name = newName // 直接修改 p 所指向对象的 Name 字段
}

func main() {
	p1 := Person{Name: "A1"}
	p2 := Person{Name: "A2"}

	p1.ChangeName("B1")
	fmt.Println(p1.Name)

	p2p := &p2
	p2p.ChangeName("B2")
	fmt.Println(p2.Name)

	var p3 Person = Person{Name: "A3"}
	p3.ChangeName("B3")
	
}