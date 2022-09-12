package dao

import (
	"Library/Library/model"
	"fmt"
	"testing"
)

func testBook(t *testing.T) {
	//t.Run("测试获取全部图书", testGetBooks)
	//t.Run("测试添加图书", testAddBook)
	//t.Run("测试通过图书ID删除图书", testDeleteBookById)
	//t.Run("测试更改图书的相关信息", testUpdateBook)
	//t.Run("测试搜索图书", testSearchBooksByName)
	t.Run("测试通过id获取图书", testGetBookById)
}
func testGetBooks(t *testing.T) {
	fmt.Println("测试获取全部图书")
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%d个图书是%v\n", k+1, v)
	}
}
func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:        "笑傲江湖",
		Author:       "金庸",
		Price:        30.0,
		Sales:        200,
		Stock:        100,
		ImgPath:      "static/images/default.jpg",
		Kind:         "武侠",
		Introduction: "11111",
	}
	AddBook(book)
}
func testDeleteBookById(t *testing.T) {
	bookId := 31
	DeleteBookById(bookId)
}
func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:           55,
		Title:        "笑傲江湖",
		Author:       "金庸",
		Price:        28.0,
		Sales:        200,
		Stock:        100,
		ImgPath:      "static/images/default.jpg",
		Kind:         "武侠",
		Introduction: "11111",
	}
	UpdateBook(book)
}
func testSearchBooksByName(t *testing.T) {
	books, _ := SearchBooksByName("到")
	for k, book := range books {
		fmt.Printf("第%d本图书的相关信息是%v\n", k+1, book)
	}
}
func testGetBookById(t *testing.T) {
	book := GetBookById(1)
	fmt.Println(book)
}
func testGetBooksByPage(t *testing.T) {
	fmt.Println("测试带分页的图书")
	page, _ := GetBooksByPage(1)
	for i, v := range page.Books {
		fmt.Printf("第%d本书是%v\n", i+1, v)
	}
	fmt.Println(page)
}
func testGetBooksByKind(t *testing.T) {
	fmt.Println("测试通过类型查图书")
	books, _ := GetBooksByKind("人文")
	for i, v := range books {
		fmt.Printf("第%d个图书是%v\n", i+1, v)
	}
}
