package model

type Book struct {
	ID           int
	Title        string
	Author       string
	Price        float64
	Sales        int
	Stock        int
	ImgPath      string
	Kind         string //图书的类型
	Introduction string
	IsLogin      bool
}
