package handle

import (
	"Library/Library/dao"
	"Library/Library/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//GetAllBooks 获取所有的图书
//func GetAllBooks(w http.ResponseWriter, r *http.Request) {
//	books, _ := dao.GetBooks()
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	t.Execute(w, books)
//}

//DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	iBookId, _ := strconv.Atoi(bookId)
	dao.DeleteBookById(iBookId)
	GetBooksByPage(w, r)
}
func SearchBooks(w http.ResponseWriter, r *http.Request) {
	searchName := r.PostFormValue("keyword")
	books, _ := dao.SearchBooksByName(searchName)
	if books != nil {
		//此时能够搜索到
		t := template.Must(template.ParseFiles("views/pages/manager/search.html"))
		t.Execute(w, books)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/search.html"))
		t.Execute(w, "")
	}
}
func ToUpdatePage(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	iBookId, _ := strconv.Atoi(bookId)
	book := dao.GetBookById(iBookId)
	if book.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
}
func AddAndUpdateBook(w http.ResponseWriter, r *http.Request) {
	BookId := r.PostFormValue("bookId")
	Title := r.PostFormValue("title")
	Price := r.PostFormValue("price")
	Author := r.PostFormValue("author")
	Sales := r.PostFormValue("sales")
	Stock := r.PostFormValue("stock")
	ImgPath := r.PostFormValue("img_path")
	Kind := r.PostFormValue("kind")
	Introduction := r.PostFormValue("introduction")
	iBookId, _ := strconv.Atoi(BookId)
	iPrice, _ := strconv.ParseFloat(Price, 64)
	iSales, _ := strconv.Atoi(Sales)
	iStock, _ := strconv.Atoi(Stock)
	book := &model.Book{
		ID:           iBookId,
		Title:        Title,
		Author:       Author,
		Price:        iPrice,
		Sales:        iSales,
		Stock:        iStock,
		ImgPath:      ImgPath,
		Kind:         Kind,
		Introduction: Introduction,
	}
	if book.ID > 0 {
		//说明是修改图书
		dao.UpdateBook(book)
	} else {
		//说明是添加图书
		dao.AddBook(book)
	}
	GetBooksByPage(w, r)
}
func GetBooksByPage(w http.ResponseWriter, r *http.Request) {
	PageNo := r.FormValue("PageNo")
	if PageNo == "" {
		PageNo = "1"
	}
	iPageNo, _ := strconv.Atoi(PageNo)
	page, _ := dao.GetBooksByPage(iPageNo)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

//GetBooksByPageUser 顾客打开全部图书界面
func GetBooksByPageUser(w http.ResponseWriter, r *http.Request) {
	PageNo := r.FormValue("PageNo")
	if PageNo == "" {
		PageNo = "1"
	}
	iPageNo, _ := strconv.Atoi(PageNo)
	page, _ := dao.GetBooksByPage(iPageNo)
	t := template.Must(template.ParseFiles("views/pages/user/allbooks.html"))
	t.Execute(w, page)
}

//ManagerBooks 管理图书
func ManagerBooks(w http.ResponseWriter, r *http.Request) {
	permission := r.FormValue("permission")
	if permission == "" {
		permission = "0"
	}
	fmt.Println(permission)
	var f bool
	//f如果是false的话，说明是顾客；true，说明是管理人员
	iPermission, _ := strconv.Atoi(permission)
	if iPermission == 0 {
		f = false
		t := template.Must(template.ParseFiles("views/pages/manager/manager.html"))
		t.Execute(w, f)
	} else {
		f = true
		fmt.Println("lyclyc")
		t := template.Must(template.ParseFiles("views/pages/manager/manager.html"))
		t.Execute(w, f)
	}
}
func SearchBooksUser(w http.ResponseWriter, r *http.Request) {
	searchName := r.PostFormValue("keyword")
	books, _ := dao.SearchBooksByName(searchName)
	if books != nil {
		//此时能够搜索到
		t := template.Must(template.ParseFiles("views/pages/user/search.html"))
		t.Execute(w, books)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/search.html"))
		t.Execute(w, "")
	}
}
func BookKind(w http.ResponseWriter, r *http.Request) {
	kind := r.FormValue("kind")
	switch kind {
	case "subject":
		books, _ := dao.GetBooksByKind("学科")
		t := template.Must(template.ParseFiles("views/pages/user/search.html"))
		t.Execute(w, books)
	case "person":
		books, _ := dao.GetBooksByKind("人文")
		t := template.Must(template.ParseFiles("views/pages/user/search.html"))
		t.Execute(w, books)
	case "unique":
		books, _ := dao.GetBooksByKind("奇异")
		t := template.Must(template.ParseFiles("views/pages/user/search.html"))
		t.Execute(w, books)
	case "poem":
		books, _ := dao.GetBooksByKind("诗集")
		t := template.Must(template.ParseFiles("views/pages/user/search.html"))
		t.Execute(w, books)
	case "biography":
		books, _ := dao.GetBooksByKind("传记")
		t := template.Must(template.ParseFiles("views/pages/user/search.html"))
		t.Execute(w, books)
	case "history":
		books, _ := dao.GetBooksByKind("历史")
		t := template.Must(template.ParseFiles("views/pages/user/search.html"))
		t.Execute(w, books)
	case "manager":
		books, _ := dao.GetBooksByKind("管理")
		t := template.Must(template.ParseFiles("views/pages/user/search.html"))
		t.Execute(w, books)
	case "success":
		books, _ := dao.GetBooksByKind("励志")
		t := template.Must(template.ParseFiles("views/pages/user/search.html"))
		t.Execute(w, books)

	}

}
func FindDetail(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	iBookID, _ := strconv.Atoi(bookID)
	book := dao.GetBookById(iBookID)
	t := template.Must(template.ParseFiles("views/pages/book/introduce.html"))
	t.Execute(w, book)
}
