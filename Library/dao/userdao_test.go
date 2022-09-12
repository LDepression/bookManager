package dao

import (
	"Library/Library/model"
	"fmt"
	"testing"
)

func testUser(t *testing.T) {
	fmt.Println("测试userdao中的函数")
	t.Run("测试用户登录的账号和密码", testCheckUserNameAndPwd)
	//t.Run("测试保存用户的代码", testSaveUser)

}
func testCheckUserNameAndPwd(t *testing.T) {
	user, _ := CheckUserNameAndPwd("admin1", "123456")
	fmt.Println("用户信息是:", user)
}
func testSaveUser(t *testing.T) {
	user := &model.User{
		UserName: "admin1",
		PassWord: "123456",
		Email:    "1197285120@qq.com",
	}
	SaveUser(user)
}
func tstGetUserById(t *testing.T) {
	fmt.Println("测试通过userID来查询用户")
	user := GetUserById(1)
	fmt.Println(user)
}
func testGetPermissionByUserID(t *testing.T) {
	fmt.Println("测试得到权限")
	p := GetPermissionByUserID(12)
	fmt.Println("权限是:", p.Permission)
}
