package model

type Comments struct {
	Comment  []*Comment
	UserName string
	UserID   int
	Book     *Book
	User     *User
}
