package library

import (
	"library-management/models"
	"sync"
)

// Library 表示图书馆
type Library struct {
	Books map[int]models.Book
	NextID int	
	mu sync.Mutex // 互斥锁
}

// NewLibrary 创建一个新的图书馆
func NewLibrary() *Library {
	return &Library{
		Books: make(map[int]models.Book),
		NextID: 1,
	}
}

// AddBook 向图书馆中添加一本书，返回书的ID
func (lib *Library) AddBook(title, author string, category models.Category, price float64) int {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	id := lib.NextID
	lib.Books[id] = models.Book{
		ID: id,
		Title: title,
		Author: author,
		Category: category,
		Price: price,
	}
	lib.NextID++
	return id
}

// DeleteBook 删除指定ID的书，返回是否删除成功
func (lib *Library) DeleteBook(id int) bool {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	if _, exists := lib.Books[id]; exists {
		delete(lib.Books, id)
		return true
	}

	return false
}

// updateBook 更新指定ID的书，返回是否成功
func (lib *Library) UpdateBook(id int, title, author string, category models.Category, price float64) bool {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	if _, exists := lib.Books[id]; exists {
		lib.Books[id] = models.Book{
			ID: id,
			Title: title,
			Author: author,
			Category: category,
			Price: price,
		}
		return true
	}

	return false
}

// 查询指定ID的书籍，返回书籍和是否存在
func (lib *Library) GetBook(id int) (models.Book, bool) {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	book, exists := lib.Books[id]
	return book, exists
}

// 列出所有书籍，返回书籍切片
func (lib *Library) ListBooks() []models.Book {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	books := make([]models.Book, 0, len(lib.Books))
	for _, book := range lib.Books {
		books = append(books, book)
	}

	return books
}

// 设置Library的Books和NextID，用于加载数据
func (lib *Library) SetBooks(books map[int]models.Book, nextID int) {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	lib.Books = books
	lib.NextID = nextID
}

// 获取Library的Books和NextID，用于保存数据
func (lib *Library) GetBooks() (map[int]models.Book, int) {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	// 创建一个新的map，避免直接返回lib.Books，导致外部修改lib.Books
	booksCopy := make(map[int]models.Book)
	for id, book := range lib.Books {
		booksCopy[id] = book
	}
	
	return booksCopy, lib.NextID
}