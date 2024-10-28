// OOP: object-oriented programming
// encapsulation: 封装

package main

import "fmt"

// struct与下面3个方法构成了一个类
// 当func的首字母大写时，表示该方法可以被其他包调用
// 那么struct的首字母大写时，表示该类可以被其他包调用
type Hero struct {
	// 类的属性首字母大写，表示该属性可以被其他包调用，否则只能在本包中使用
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	fmt.Println("hero = ", this)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	// this是调用该方法的对象的一个副本
	this.Name = newName
}

func (this *Hero) SetName2(newName string) {
	// this是调用该方法的对象的一个指针
	this.Name = newName
}

func main() {
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}
	hero.Show() // hero =  {zhang3 100 1}

	hero.SetName("liu4")
	hero.Show() // hero =  {zhang3 100 1}

	hero.SetName2("liu4")
	hero.Show() // hero =  {liu4 100 1}
}