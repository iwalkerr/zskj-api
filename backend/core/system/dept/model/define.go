package model

import "xframe/backend/common/constant"

type Entity struct {
	DeptId    int    `db:"dept_id" json:"dept_id"`     // 部门id
	ParentId  int    `db:"parent_id" json:"parent_id"` // 父部门id
	Ancestors string `db:"ancestors" json:"ancestors"` // 祖级列表
	DeptName  string `db:"dept_name" json:"dept_name"` // 部门名称
	OrderNum  int    `db:"order_num" json:"order_num"` // 显示顺序
	Leader    string `db:"leader" json:"leader"`       // 负责人
	Phone     string `db:"phone" json:"phone"`         // 联系电话
	Email     string `db:"email" json:"email"`         // 邮箱
	Status    string `db:"status" json:"status" `      // 部门状态（0正常 1停用）
	DelFlag   string `db:"del_flag" json:"del_flag"`   // 删除标志（0代表存在 2代表删除）
	constant.ModelData
	ParentName string
}

type SelectPageReq struct {
	DeptId   string `form:"deptId"`   // 部门ID
	DeptName string `form:"deptName"` // 部门名称
	Visible  string `form:"status"`   // 是否可见

	*constant.PageReq
}

//新增页面请求参数
type AddReq struct {
	ParentId int    `form:"parentId"`
	DeptName string `form:"deptName"  binding:"required"`
	OrderNum int    `form:"orderNum" binding:"required"`
	Leader   string `form:"leader"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Status   string `form:"status"`
}

//修改页面请求参数
type EditReq struct {
	DeptId int `form:"deptId" binding:"required"`
	AddReq
}

//检查菜单名称请求参数
type CheckDeptNameReq struct {
	DeptId   int    `form:"deptId"  binding:"required"`
	ParentId int    `form:"parentId"`
	DeptName string `form:"deptName"  binding:"required"`
}

//检查菜单名称请求参数
type CheckDeptNameAllReq struct {
	ParentId int    `form:"parentId"`
	DeptName string `form:"deptName"  binding:"required"`
}
