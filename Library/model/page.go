package model

type Page struct {
	Books         []*Book //每一页显示出来的切片
	PageNo        int     //当前页数
	PageSize      int     //每页显示的条数
	TotalPageSize int     //总页数
	TotalRecord   int     //总记录数，通过查询数据库得到
	IsLogin       bool
}

//IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

//IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageSize
}

//ToPrevPageNo 获取上一页
func (p *Page) ToPrevPageNo() int {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}
func (p *Page) ToNextPageNo() int {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageSize
	}
}
