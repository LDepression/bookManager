package model

type Comment struct {
	ID       int
	BookID   int
	UserID   int
	Time     string
	Text     string
	AllText  []string
	Book     *Book
	UserName string
}
