package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:zxz123456@tcp(localhost:3306)/Library")
	if err != nil {
		panic(err.Error())
	}
	err := Db.Ping()
	if err != nil {
		fmt.Println("数据库密码错误")
		panic(err.Error())
	}
}
