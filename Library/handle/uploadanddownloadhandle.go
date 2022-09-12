package handle

import (
	"Library/Library/dao"
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

// 统一错误输出接口
func errorHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

//UploadBook 上传图书
func UploadBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	r.ParseForm()
	//Title:=r.FormValue("title")
	uploadFile, handle, _ := r.FormFile("book")
	ext := strings.ToLower(path.Ext(handle.Filename))
	if ext != ".txt" {
		errorHandle(errors.New("只支持txt文件上传"), w)
		return
		//defer os.Exit(2)
	}
	os.Mkdir("./uploaded/", 0777)
	saveFile, err := os.OpenFile("./uploaded/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	errorHandle(err, w)
	io.Copy(saveFile, uploadFile)

	defer uploadFile.Close()
	defer saveFile.Close()
	// 上传图片成功
	w.Write([]byte("查看上传图书: <a target='_blank' href='/upload/" + handle.Filename + "'>" + handle.Filename + "</a>"))

}

func ShowBookHandle(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("." + req.URL.Path)
	errorHandle(err, w)

	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	errorHandle(err, w)
	w.Write(buff)
}
func DownLoad(w http.ResponseWriter, req *http.Request) {
	fileName := req.FormValue("title")
	fileName = fileName + ".txt"
	//设置响应头
	header := w.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+fileName)
	b, _ := ioutil.ReadFile("./uploaded/" + fileName)
	w.Write(b)
}

//Upload 去上传界面
func Upload(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	iBookID, _ := strconv.Atoi(bookID)
	book := dao.GetBookById(iBookID)
	t := template.Must(template.ParseFiles("views/pages/manager/upload.html"))
	t.Execute(w, book)
}
