package dao

import (
	"Library/Library/model"
	"fmt"
	"testing"
	"time"
)

func TestAddComment(t *testing.T) {
	fmt.Println("测试添加一条评论")
	comment := &model.Comment{
		ID:       1,
		BookID:   33,
		UserID:   1,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		Text:     "你好呀",
		UserName: "admin1",
	}
	AddComment(comment)
}
func testGetAllComment(t *testing.T) {
	comments, _ := GetAllComment(33)
	for i, v := range comments {
		fmt.Printf("第%d个评论是%v\n", i+1, v)
	}
}
func testDeleteComment(t *testing.T) {
	DeleteComment(1)
}
