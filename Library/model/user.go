package model

var (
	userPerm   = 0
	managePerm = 1
)

type User struct {
	ID         int
	UserName   string
	PassWord   string
	Email      string
	IsLogin    bool
	Permission int
}
