package handle

import (
	"Library/Library/dao"
	"Library/Library/model"
	"Library/Library/utils"
	"fmt"
	"net/http"
	"text/template"
)

func LogOut(w http.ResponseWriter, r *http.Request) {
	session, _ := dao.IsLogin(r)
	//删除
	dao.DeleteSession(session)
	cookie, _ := r.Cookie("user")
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	IndexHandle(w, r)
}
func Login(w http.ResponseWriter, r *http.Request) {
	_, f := dao.IsLogin(r)
	if f == true {
		IndexHandle(w, r)
	} else {
		username := r.PostFormValue("username")
		pwd := r.PostFormValue("password")
		user, _ := dao.CheckUserNameAndPwd(username, pwd)
		if user.ID > 0 {
			//账号密码正确
			//生成UUID作为sessionID
			uuid := utils.CreateUUID()
			session := &model.Session{
				SessionID: uuid,
				UserName:  username,
				UserID:    user.ID,
			}
			//将session添加到数据库中去
			dao.AddSession(session)
			//fmt.Println(1111)
			//在这里添加设置权限
			userr := dao.GetPermissionByUserID(user.ID)
			//fmt.Println("lyclyc的权限是", userr.Permission)
			user1 := model.User{
				ID:         user.ID,
				UserName:   username,
				PassWord:   pwd,
				IsLogin:    true,
				Permission: userr.Permission,
			}
			//创建cookie,将cookie的值设置为uuid
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			fmt.Println("登录成功")
			t.Execute(w, user1)
		} else {
			//账号或密码有误
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			fmt.Println("登录失败")
			t.Execute(w, "账号或密码有误")
		}
	}
}
func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user := &model.User{
		UserName: username,
		PassWord: password,
		Email:    email,
	}
	User, _ := dao.CheckUserName(username)
	if User.ID > 0 {
		//说明此时有这个用户
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "<font style='color:red'>该用户已经注册过了</font>")
	} else {
		//说明用户还没有注册
		dao.SaveUser(user)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, user)
	}
}
func CheckUserNameOK(w http.ResponseWriter, r *http.Request) {
	//获取用户名
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户已经存在了
		w.Write([]byte("<font style='color:red'>用户名已存在</font>"))
	} else {
		//用户名不存在
		w.Write([]byte("<font style='color:green'>用户名可用</font>"))
	}
}
func AddManager(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	manager := &model.User{
		UserName:   username,
		PassWord:   password,
		Email:      email,
		IsLogin:    false,
		Permission: 1,
	}
	Manager, _ := dao.CheckUserName(username)
	if Manager.ID > 0 {
		//说明此时有这个用户
		t := template.Must(template.ParseFiles("views/pages/manager/regist.html"))
		t.Execute(w, "<font style='color:red'>该用户已经注册过了</font>")
	} else {
		//说明用户还没有注册
		dao.SaveManager(manager)
		//fmt.Println("添加成功")
		t := template.Must(template.ParseFiles("views/pages/manager/regist_success.html"))
		t.Execute(w, manager)
	}
}
