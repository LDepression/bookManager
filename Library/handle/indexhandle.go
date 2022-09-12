package handle

import (
	"Library/Library/dao"
	"Library/Library/model"
	"html/template"
	"net/http"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	//解析模板
	session, f := dao.IsLogin(r)
	if f == true {
		userId := session.UserID
		user := dao.GetUserById(userId)
		//说明此用户在线上
		user.IsLogin = true
		p := dao.GetPermissionByUserID(userId)
		user.Permission = p.Permission
		t := template.Must(template.ParseFiles("views/index.html"))
		t.Execute(w, user)
	} else {
		user := &model.User{}
		t := template.Must(template.ParseFiles("views/index.html"))
		t.Execute(w, user)
	}
}
