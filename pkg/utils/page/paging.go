package page

import (
	"xframe/backend/common/constant"
	"xframe/backend/common/db"

	"github.com/jmoiron/sqlx"
)

type Paging struct {
	dbType        []int // 数据库类型,默认是主数据库，1代表从数据库
	pageNum       int   // 当前页
	pagesize      int   // 每页条数
	total         *int  // 总条数
	pageCount     int   // 总页数
	currentResult int   // 当前记录起始索引
}

//创建分页
func New(common *constant.PageReq, dbType ...int) *Paging {
	if common.PageNum < 1 {
		common.PageNum = 1
	}
	if common.PageSize < 1 {
		common.PageSize = 10
	}

	total := new(int)
	common.Total = total

	return &Paging{
		pageNum:  common.PageNum,
		pagesize: common.PageSize,
		dbType:   dbType,
		total:    total,
	}
}

// 数据库查出数据的条数
func (p *Paging) GetRows(query string, param []interface{}) (*sqlx.Rows, error) {
	// 数据库查出数据的总数
	_ = db.Conn(p.dbType...).QueryRow("select count(0) from ("+query+") tmp_count", param[:]...).Scan(p.total)

	query += " limit ?,?"
	param = append(param, p.getCurrentResult(), p.pagesize)
	rows, err := db.Conn(p.dbType...).Queryx(query, param[:]...)

	return rows, err
}

// 当前记录起始索引
func (p *Paging) getCurrentResult() int {
	p.currentResult = (p.getCurrentPage() - 1) * p.pagesize
	if p.currentResult < 0 {
		p.currentResult = 0
	}
	return p.currentResult
}

// 获取总共页数
func (p *Paging) getTotalPage() int {
	if *(p.total)%p.pagesize == 0 {
		p.pageCount = *(p.total) / p.pagesize
	} else {
		p.pageCount = *(p.total)/p.pagesize + 1
	}
	return p.pageCount
}

// 获取当前页
func (p *Paging) getCurrentPage() int {
	if p.pageNum <= 0 {
		p.pageNum = 1
	}
	if p.pageNum > p.getTotalPage() {
		p.pageNum = p.getTotalPage()
	}
	return p.pageNum
}
