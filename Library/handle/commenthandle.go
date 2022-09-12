package handle

import (
	"Library/Library/dao"
	"Library/Library/model"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func GetAllComment(w http.ResponseWriter, r *http.Request) {
	session, _ := dao.IsLogin(r)
	userID := session.UserID
	userName := session.UserName
	user := dao.GetUserById(userID)
	bookID := r.FormValue("bookID")
	iBookID, _ := strconv.Atoi(bookID)
	book := dao.GetBookById(iBookID)
	comments, _ := dao.GetAllComment(iBookID)
	//fmt.Println("userName:", userName)
	Comments := model.Comments{
		Comment:  comments,
		UserName: userName,
		UserID:   userID,
		Book:     book,
		User:     user,
	}
	t := template.Must(template.ParseFiles("views/pages/user/comment.html"))
	t.Execute(w, Comments)
}
func AddComment(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("1111111111")
	session, _ := dao.IsLogin(r)
	userID := session.UserID
	text := r.PostFormValue("text")
	bookID := r.FormValue("BookID")
	iBookID, _ := strconv.Atoi(bookID)
	user := dao.GetUserById(userID)
	userName := user.UserName
	comment := &model.Comment{
		BookID:   iBookID,
		UserID:   userID,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		Text:     text,
		UserName: userName,
	}
	dao.AddComment(comment)
	//GetAllComment(w, r)
	w.Write([]byte("评论添加成功"))
}
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentID := r.FormValue("commentID")
	iCommentID, _ := strconv.Atoi(commentID)
	dao.DeleteComment(iCommentID)
	w.Write([]byte("删除成功"))
}
