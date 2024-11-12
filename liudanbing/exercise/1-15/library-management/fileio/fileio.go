package fileio

import (
	"encoding/json"
	"library-management/library"
	"library-management/models"
	"os"
)

// SaveToFile 将图书馆的图书数据保存到文件中
func SaveToFile(lib *library.Library, filename string) error {
	file, err := os.Create(filename) // 创建文件
	if err != nil {
		return err
	}
	defer file.Close() // 函数返回前关闭文件

	books, nextID := lib.GetBooks() // 获取图书数据
	data := struct {
		Books map[int]models.Book `json:"books"`
		NextID int `json:"next_id"`
	}{
		Books: books,
		NextID: nextID,
	}

	encoder := json.NewEncoder(file) // 创建JSON编码器
	encoder.SetIndent("", "    ") // 设置缩进
	return encoder.Encode(data) // 编码并写入文件
}

// LoadFromFile 从文件中加载图书数据到图书馆
func LoadFromFile(lib *library.Library, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var data struct {
		Books map[int]models.Book `json:"books"`
		NextID int `json:"next_id"`
	}

	decoder := json.NewDecoder(file) // 创建JSON解码器
	err = decoder.Decode(&data) // 解码文件内容, 并存入data
	if err != nil {
		return err
	}

	lib.SetBooks(data.Books, data.NextID) // 设置图书数据
	return nil
}