package main

import "fmt"

type Person struct {
    Name string
}

// 指针接收者方法
func (p *Person) ChangeName(newName string) {
    (*p).Name = newName  // 直接修改 p 所指向对象的 Name 字段
}

func main() {
    p := Person{Name: "Alice"}

    // // 传递指针 &p 给 ChangeName 方法
    // p.ChangeName("Bob")
    // fmt.Println(p.Name)  // Bob，因为 ChangeName 修改了原始对象
	fmt.Println(p)
	//fmt.Printf("%p", &p)
	fmt.Printf("%T",p)
	p.ChangeName("bob")
	fmt.Println(p)
}
