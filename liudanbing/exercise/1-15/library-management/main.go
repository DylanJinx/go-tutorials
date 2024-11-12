package main

import (
	"bufio"
	"fmt"
	"library-management/fileio"
	"library-management/library"
	"library-management/models"
	"os"
	"strconv" // 字符串和数字的转换
)

func main() {
	lib := library.NewLibrary() // 创建一个新的图书馆
	scanner := bufio.NewScanner(os.Stdin) // 创建一个新的扫描器
	for {
		fmt.Println("\n======== library management system ========")
		fmt.Println("Please select an option:")
		fmt.Println("1. Add a book")
		fmt.Println("2. Delete a book")
		fmt.Println("3. Update a book")
		fmt.Println("4. Search for books")
		fmt.Println("5. List all books")
		fmt.Println("6. Save to file")
		fmt.Println("7. Load from file")
		fmt.Println("8. Exit")
		fmt.Print("Your choice(1-8): ")

		scanner.Scan() // 扫描输入
		choice := scanner.Text() // 获取输入

		switch choice {
		case "1":
			addBook(lib, scanner)
		case "2":
			deleteBook(lib, scanner)
		case "3":
			updateBook(lib, scanner)
		case "4":
			queryBooks(lib, scanner)
		case "5":
			listBooks(lib)
		case "6":
			saveData(lib, scanner)
		case "7":
			loadData(lib, scanner)
		case "8":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func addBook(lib *library.Library, scanner *bufio.Scanner) {
	fmt.Print("pleace enter the title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("pleace enter the author: ")
	scanner.Scan()
	author := scanner.Text()

	fmt.Println("choice the category number:")
	for i, category := range models.CategoryNames {
		fmt.Printf("%d. %s\n", i, category)
	}

	fmt.Print("please enter the category number: ")
	scanner.Scan()
	categoryInput, err := strconv.Atoi(scanner.Text()) // 将输入转换为整数
	if err != nil || categoryInput < 0 || categoryInput >= len(models.CategoryNames) {
		fmt.Println("Invalid category number!")
		return
	}

	category := models.Category(categoryInput)

	fmt.Print("please enter the price: ")
	scanner.Scan()
	price, err := strconv.ParseFloat(scanner.Text(), 64) // 将输入转换为浮点数64位
	if err != nil {
		fmt.Println("Invalid price!")
		return
	}

	id := lib.AddBook(title, author, category, price)
	fmt.Printf("Book added with ID: %d\n", id)
}

func deleteBook(lib *library.Library, scanner *bufio.Scanner) {
	fmt.Print("Enter the ID of the book you want to delete: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID!")
		return
	}

	success := lib.DeleteBook(id)
	if success {
		fmt.Println("Book deleted successfully!")
	} else {
		fmt.Println("Book not found!")
	}
}

func updateBook(lib *library.Library, scanner *bufio.Scanner) {
	fmt.Print("Enter the ID of the book you want to update: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID!")
		return
	}

	book, exists := lib.GetBook(id)
	if !exists {
		fmt.Println("Book not found!")
		return
	}

	fmt.Printf("Current book title: %s\n", book.Title)
	fmt.Print("Enter a new title (leave blank to indicate no changes): ")
	scanner.Scan()
	title := scanner.Text()
	if title == "" {
		title = book.Title
	}

	fmt.Printf("Current book author: %s\n", book.Author)
	fmt.Print("Enter a new author (leave blank to indicate no changes): ")
	scanner.Scan()
	author := scanner.Text()
	if author == "" {
		author = book.Author
	}

	fmt.Printf("Current book category: %s\n", book.Category)
	fmt.Println("Choose a new category(leave blank to indicate no changes):")
	for i, category := range models.CategoryNames {
		fmt.Printf("%d. %s\n", i, category)
	}
	fmt.Print("Enter the category number: ")
	scanner.Scan()
	categoryInputStr := scanner.Text()
	var category models.Category
	if categoryInputStr == "" {
		category = book.Category
	} else {
		categoryInput, err := strconv.Atoi(categoryInputStr)
		if err != nil || categoryInput < 0 || categoryInput >= len(models.CategoryNames) {
			fmt.Println("Invalid category number!")
			return
		}
		category = models.Category(categoryInput)
	}

	fmt.Printf("Current book price: %.2f\n", book.Price)
	fmt.Print("Enter a new price (leave blank to indicate no changes): ")
	scanner.Scan()
	priceInput := scanner.Text()
	var price float64
	if priceInput == "" {
		price = book.Price
	} else {
		price, err = strconv.ParseFloat(priceInput, 64)
		if err != nil {
			fmt.Println("Invalid price!")
			return
		}
	}

	success := lib.UpdateBook(id, title, author, category, price)
	if success {
		fmt.Println("Book updated successfully!")
	} else {
		fmt.Println("Book not found!")
	}

}

func queryBooks(lib *library.Library, scanner *bufio.Scanner) {
	fmt.Print("Enter the ID of the book you want to search: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID!")
		return
	}

	book, exists := lib.GetBook(id)
	if exists {
		book.PrintDetails()
	} else {
		fmt.Println("Book not found!")
	}
}

func listBooks(lib *library.Library) {
	books := lib.ListBooks()
	if len(books) == 0 {
		fmt.Println("No books in the library!")
		return
	}

	fmt.Println("Books in the library:")
	for i, book := range books {
		fmt.Printf("index: %d. ", i+1)
		book.PrintDetails()
	}
}

func saveData(lib *library.Library, scanner *bufio.Scanner) {
	fmt.Print("Enter the name of the file you want to save (e.g. books.json): ")
	scanner.Scan()
	filename := scanner.Text()
	err := fileio.SaveToFile(lib, filename)
	if err != nil {
		fmt.Println("Error saving data to file:", err)
	} else {
		fmt.Println("Data saved to file successfully!")
	}
}

func loadData(lib *library.Library, scanner *bufio.Scanner) {
	fmt.Print("Enter the name of the file you want to load (e.g. books.json): ")
	scanner.Scan()
	filename := scanner.Text()
	err := fileio.LoadFromFile(lib, filename)
	if err != nil {
		fmt.Println("Error loading data from file:", err)
	} else {
		fmt.Println("Data loaded from file successfully!")
	}
}