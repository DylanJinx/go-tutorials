package main

import (
	"fmt"
)

// 接口本质是一个指针，interface里面有一个指针指向当前interface的具体类型，以及当前类型所包含的函数列表
type AnimalIF interface {
	// 接口里面的方法不能有方法体，不需要func关键字
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string // 获取动物的种类
}

// 实现接口的类
type Cat struct {	
	// AnimalIF // 不需要显示声明实现了哪个接口，只需要实现接口的方法即可
	color string
}

// 如果一个类没有实现接口的所有方法，那么接口的指针就不能指向该类
func (this *Cat) Sleep() {
	fmt.Println("Cat Sleep()")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的实现类
type Dog struct {
	color string
}

// 如果一个类没有实现接口的所有方法，那么接口的指针就不能指向该类
func (this *Dog) Sleep() {
	fmt.Println("Dog Sleep()")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() // 多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	var animal AnimalIF // 接口的数据类型，父类指针
	animal = &Cat{"Orange"}
	animal.Sleep() // Cat Sleep()

	animal = &Dog{"Black"}
	animal.Sleep() // Dog Sleep()

	fmt.Println("=====================================")
	cat := Cat{"Orange"}
	dog := Dog{"Black"}
	showAnimal(&cat)
	showAnimal(&dog)
}