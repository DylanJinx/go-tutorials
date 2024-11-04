package main

import (
	"fmt"
)

type Book struct {
	Title string
	Author string
	ISBN string
}

// ISBN -> Book
var Book_map = make(map[string]Book)

var Book_list = make([]Book, 0)

func init() {
	fmt.Println("Welcome to the library management system!")
}

func AddBook(title, author, isbn string) (bool, string) {
	if _, exists := Book_map[isbn]; exists { // 有两个返回值，第一个是value，第二个是是否存在
		return false, "The book already exists!"
	}

	book := Book{Title: title, Author: author, ISBN: isbn}
	Book_map[isbn] = book
	Book_list = append(Book_list, book)
	return true, "The book has been added successfully!"
}

func ViewBook() {
	fmt.Println("The books in the library are as follows:")
	for _, book := range Book_list { // range返回两个值，第一个是索引，第二个是值
		fmt.Printf("Title: %s, Author: %s, ISBN: %s\n", book.Title, book.Author, book.ISBN)
	}
}

func DeleteBook(isbn string) (bool, string) {
	if _, exists := Book_map[isbn]; !exists {
		return false, "The book does not exist!"
	}

	for i, book := range Book_list {
		if book.ISBN == isbn {
			// 
			Book_list = append(Book_list[:i], Book_list[i+1:]...) // 删除切片中的元素
			delete(Book_map, isbn) // 删除map中的元素
			return true, "The book has been deleted successfully!"
		}
	}

	return false, "The book has been deleted failed!"
}

func main() {
	defer fmt.Println("The program is over, thank you for using it!")

	var choice int

	for {
		fmt.Println("---------------------------------------------")
		fmt.Println("1. Add a book")
		fmt.Println("2. View all books")
		fmt.Println("3. Delete a book")
		fmt.Println("4. Exit")
		fmt.Print("Please enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
			case 1:
				var title, author, isbn string
				fmt.Println("Please enter the title of the book:")
				fmt.Scanln(&title)
				fmt.Println("Please enter the author of the book:")
				fmt.Scanln(&author)
				fmt.Println("Please enter the ISBN of the book:")
				fmt.Scanln(&isbn)
				_, msg := AddBook(title, author, isbn)
				fmt.Println(msg)

			case 2:
				ViewBook()

			case 3:
				var isbn string
				fmt.Println("Please enter the ISBN of the book you want to delete:")
				fmt.Scanln(&isbn)
				_, msg := DeleteBook(isbn)
				fmt.Println(msg)

			case 4:
				return

			default:
				fmt.Println("Invalid choice, please re-enter!")
		}
	}
}