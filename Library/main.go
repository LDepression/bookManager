package main

import (
	"Library/Library/handle"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/login", handle.Login)
	http.HandleFunc("/logOut", handle.LogOut)
	http.HandleFunc("/register", handle.Register)
	http.HandleFunc("/checkUserNameOK", handle.CheckUserNameOK)
	//http.HandleFunc("/getAllBooks", handle.GetAllBooks)
	http.HandleFunc("/managerBooks", handle.ManagerBooks)
	http.HandleFunc("/bookKind", handle.BookKind)
	http.HandleFunc("/addManager", handle.AddManager)
	http.HandleFunc("/deleteBook", handle.DeleteBook)
	http.HandleFunc("/getBooksByPage", handle.GetBooksByPage)
	http.HandleFunc("/getBooksByPageUser", handle.GetBooksByPageUser)
	http.HandleFunc("/searchBooks", handle.SearchBooks)
	http.HandleFunc("/searchBooksUser", handle.SearchBooksUser)
	http.HandleFunc("/toUpdatePage", handle.ToUpdatePage)
	http.HandleFunc("/addAndUpdateBook", handle.AddAndUpdateBook)
	http.HandleFunc("/findDetail", handle.FindDetail)
	http.HandleFunc("/uploadBook", handle.UploadBook)
	http.HandleFunc("/upload", handle.Upload)
	http.HandleFunc("/download", handle.DownLoad)
	http.HandleFunc("/addComment", handle.AddComment)
	http.HandleFunc("/deleteComment", handle.DeleteComment)
	http.HandleFunc("/getAllComment", handle.GetAllComment)
	http.HandleFunc("/main", handle.IndexHandle)
	http.ListenAndServe(":9090", nil)
}
