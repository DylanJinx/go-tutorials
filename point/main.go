package main

import (
	"fmt"
)

// func modify(c int32) *int32 {
// 	c++
// 	return &c
// }

type Book struct {
	Class byte
	Price int32
}

func modify1(b Book) Book {
	fmt.Printf("1 address %p\n", &b)
	b.Price++
	return b
}

func modify2(b *Book) Book {
	fmt.Printf("2 address %p\n", b)
	b.Price++
	return *b
}

func modify3(b Book) *Book {
	fmt.Printf("3 address %p\n", &b)
	b.Price++
	return &b
}

func modify4(b *Book) *Book {
	fmt.Printf("4 address %p\n", b)
	b.Price++
	return b
}

func main() {
	// var a int32 = 1
	// b := modify(a)
	// fmt.Println("b = ", *b)

	book_1 := Book{Class: 1, Price: 100}
	fmt.Printf("book_1 address %p\n", &book_1)
	book1 := modify1(book_1)
	fmt.Printf("book1 address %p\n", &book1)

	fmt.Println("--------------------")

	book_2 := Book{Class: 2, Price: 200}
	fmt.Printf("book_2 address %p\n", &book_2)
	book2 := modify2(&book_2)
	fmt.Printf("book2 address %p\n", &book2)

	fmt.Println("--------------------")

	book_3 := Book{Class: 3, Price: 300}
	fmt.Printf("book_3 address %p\n", &book_3)
	book3 := modify3(book_3)
	fmt.Printf("book3 address %p\n", book3)

	fmt.Println("--------------------")

	book_4 := Book{Class: 4, Price: 400}
	fmt.Printf("book_4 address %p\n", &book_4)
	book4 := modify4(&book_4)
	fmt.Printf("book4 address %p\n", book4)
}
