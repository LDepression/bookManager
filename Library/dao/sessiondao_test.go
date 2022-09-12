package dao

import (
	"Library/Library/model"
	"fmt"
	"testing"
)

func testAddSession(t *testing.T) {
	fmt.Println("测试添加session")
	session := &model.Session{
		SessionID: "130130130",
		UserName:  "admin1",
		UserID:    1,
	}
	AddSession(session)
}
func testDeleteSession(t *testing.T) {
	fmt.Println("测试删除session")
	session := GetSessionByID("130130130")
	DeleteSession(session)
}
