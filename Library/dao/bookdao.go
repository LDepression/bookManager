package dao

import (
	"Library/Library/model"
	"Library/Library/utils"
)

func GetBooks() ([]*model.Book, error) {
	defer utils.RwMutexBook.RUnlock()
	utils.RwMutexBook.RLock()

	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	var books []*model.Book
	rows, _ := utils.Db.Query(sqlStr)
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}
func GetBooksByPage(pageNo int) (*model.Page, error) {
	defer utils.RwMutexBook.RUnlock()
	utils.RwMutexBook.RLock()
	sqlStr1 := "select count(*) from books"
	row := utils.Db.QueryRow(sqlStr1)
	//查询图书的记录数
	var TotalRecord int
	row.Scan(&TotalRecord)
	//设置每页有多少图书
	var PageSize = 4
	//图书的页数
	var TotalPageSize int
	if TotalRecord%PageSize == 0 {
		TotalPageSize = TotalRecord / PageSize
	} else {
		TotalPageSize = TotalRecord/PageSize + 1
	}

	sqlStr := "select id,title,author,price,sales,stock,img_path,kind,introduction from books limit ?,?"
	var books []*model.Book
	rows2, _ := utils.Db.Query(sqlStr, (pageNo-1)*PageSize, PageSize)
	for rows2.Next() {
		book := &model.Book{}
		rows2.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &book.Kind, &book.Introduction)
		books = append(books, book)
	}
	page := &model.Page{
		Books:         books,
		PageNo:        pageNo,
		PageSize:      PageSize,
		TotalPageSize: TotalPageSize,
		TotalRecord:   TotalRecord,
	}
	return page, nil
}
func AddBook(book *model.Book) error {
	defer utils.RwMutexBook.Unlock()
	utils.RwMutexBook.Lock()
	sqlStr := "insert into books(title,author,price,sales,stock,img_path,kind,introduction) values(?,?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath, book.Kind, book.Introduction)
	if err != nil {
		return err
	}
	return nil
}
func DeleteBookById(id int) error {
	defer utils.RwMutexBook.Unlock()
	utils.RwMutexBook.Lock()
	sqlStr := "delete from books where id=?"
	_, err := utils.Db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}
func UpdateBook(book *model.Book) error {
	defer utils.RwMutexBook.Unlock()
	utils.RwMutexBook.Lock()
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=?,img_path=?,kind=?,introduction=? where id=?"
	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath, book.Kind, book.Introduction, book.ID)
	if err != nil {
		return err
	}
	return nil
}
func SearchBooksByName(bookName string) ([]*model.Book, error) {
	defer utils.RwMutexBook.RUnlock()
	utils.RwMutexBook.RLock()
	rows, err := utils.Db.Query("select id,title,author,price,sales,stock,img_path,kind from books where title like ?", "%"+bookName+"%")
	var books []*model.Book
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			book := &model.Book{}
			rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &book.Kind)
			books = append(books, book)
		}
		return books, nil
	}

}
func GetBookById(id int) *model.Book {
	defer utils.RwMutexBook.RUnlock()
	utils.RwMutexBook.RLock()
	sqlStr := "select id,title,author,price,sales,stock,img_path,kind,introduction from books where id=?"
	row := utils.Db.QueryRow(sqlStr, id)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &book.Kind, &book.Introduction)
	return book
}
func GetBooksByKind(kind string) ([]*model.Book, error) {
	defer utils.RwMutexBook.RUnlock()
	utils.RwMutexBook.RLock()
	rows, err := utils.Db.Query("select id,title,author,price,sales,stock,img_path,kind,introduction from books where kind=?", kind)
	var books []*model.Book
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			book := &model.Book{}
			rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &book.Kind, &book.Introduction)
			books = append(books, book)
		}
		return books, nil
	}
}
