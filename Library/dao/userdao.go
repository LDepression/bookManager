package dao

import (
	"Library/Library/model"
	"Library/Library/utils"
	"crypto/md5"
)

//CheckUserNameAndPwd 检查用户名和密码
func CheckUserNameAndPwd(userName, passWord string) (*model.User, error) {
	defer utils.RwMutexUser.RUnlock()
	//读加锁
	utils.RwMutexUser.RLock()
	newPwd := utils.CreatMdStr(md5.Sum([]byte(passWord)))
	//写sql语句
	sqlStr := "select id,username,password,email from user where username=? and password=?"
	row := utils.Db.QueryRow(sqlStr, userName, newPwd)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	return user, err
}

//CheckUserName 检查用户名避免用户名重复
func CheckUserName(userName string) (*model.User, error) {
	defer utils.RwMutexUser.RUnlock()
	utils.RwMutexUser.RLock()
	sqlStr := "select id,username,password,email from user where username=?"
	row := utils.Db.QueryRow(sqlStr, userName)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	return user, err
}

//SaveUser 保存用户
func SaveUser(user *model.User) error {
	defer utils.RwMutexUser.Unlock()
	utils.RwMutexUser.Lock()
	newPassWord := utils.CreatMdStr(md5.Sum([]byte(user.PassWord)))
	sqlStr := "insert into user(username,password,email) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, user.UserName, newPassWord, user.Email)
	return err
}
func GetUserById(userId int) *model.User {
	defer utils.RwMutexUser.RUnlock()
	utils.RwMutexUser.RLock()
	sqlStr := "select id,username,password,email from user where id=?"
	row := utils.Db.QueryRow(sqlStr, userId)
	user := &model.User{}
	row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	return user
}

//GetPermissionByUserID 通过用户的ID获取用户的权限
func GetPermissionByUserID(userID int) *model.User {
	defer utils.RwMutexUser.RUnlock()
	utils.RwMutexUser.RLock()
	sqlStr := "select id,username,password,email,permission from user where id=?"
	row := utils.Db.QueryRow(sqlStr, userID)
	user := &model.User{}
	row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email, &user.Permission)
	return user
}
func SaveManager(user *model.User) error {
	defer utils.RwMutexUser.Unlock()
	utils.RwMutexUser.Lock()
	newPassWord := utils.CreatMdStr(md5.Sum([]byte(user.PassWord)))
	sqlStr := "insert into user(username,password,email,permission) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, user.UserName, newPassWord, user.Email, 1)
	return err
}
