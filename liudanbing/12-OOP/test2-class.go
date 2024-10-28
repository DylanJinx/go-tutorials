// Inheritance: 继承
package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human Eat()")
}

func (this *Human) Walk() {
	fmt.Println("Human Walk()")
}

type SuperMan struct {
	Human // SuperMan继承了Human的所有方法
	level int
}

// 重定义Human的Eat方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eat()")
}

// 子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan Fly()")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name, "sex = ", this.sex, "level = ", this.level)
}

func main() {
	h := Human{"zhang3", "female"}
	
	h.Eat()
	h.Walk()

	// 定义一个子类对象
	s := SuperMan{Human{"li4", "female"}, 88}
	s.Walk() // 父类的方法
	s.Eat() // 子类重写的方法
	s.Fly() // 子类的新方法
	s.Print()

	// 另一种定义子类对象的方式
	var s2 SuperMan
	s2.name = "wang5"
	s2.sex = "male"
	s2.level = 99
	s2.Print()
}