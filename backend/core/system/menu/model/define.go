package model

import "xframe/backend/common/constant"

type Entity struct {
	MenuId     int    `db:"menu_id" json:"menu_id"`     // 菜单ID
	Name       string `db:"name" json:"menu_name"`      // 菜单名称
	Url        string `db:"url" json:"url"`             // 菜单路径
	ParentId   int    `db:"parent_id" json:"parent_id"` // 父菜单ID
	SortId     int    `db:"sort_id" json:"sort_id"`     // 排序
	Icon       string `db:"icon" json:"icon"`           // 图标
	SysType    string `db:"sys_type" json:"sys_type"`   // 系统菜单类型，业务类型，系统类型2种
	MenuType   string `db:"menu_type" json:"menu_type"` // 菜单类型，F:按钮类型，C:菜单类型，M:目录类型
	Visible    string `db:"visible" json:"visible"`     // 状态 菜单是否隐藏
	Perms      string `db:"perms" json:"perms"`         // 菜单权限
	ParentName string `db:"parentName"`
	constant.ModelData
	SubMenu []Entity `json:"subMenu,omitempty"`
	Color   string   // 颜色
}

type SelectPageReq struct {
	MenuId   string `form:"menuId"`   // 菜单ID
	MenuName string `form:"menuName"` // 菜单名称
	Visible  string `form:"status"`   // 是否可见
	SysType  string `form:"sysType"`  // 菜单类型

	*constant.PageReq
}

//新增页面请求参数
type AddReq struct {
	ParentId int    `form:"parentId"`
	SysType  string `form:"sysType"  binding:"required"`
	MenuType string `form:"menuType"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
	OrderNum int    `form:"orderNum"`
	Url      string `form:"url"`
	Icon     string `form:"icon"`
	Perms    string `form:"perms"`
	Visible  string `form:"visible"`
	Color    string `form:"color"`
}

type EditReq struct {
	MenuId int `form:"menuId" binding:"required"`
	AddReq
}

//检查菜单名称请求参数
type CheckMenuNameReq struct {
	MenuId   int    `form:"menuId"  binding:"required"`
	ParentId int    `form:"parentId"`
	MenuName string `form:"menuName"  binding:"required"`
}

//检查菜单名称请求参数
type CheckMenuNameAllReq struct {
	ParentId int    `form:"parentId"`
	MenuName string `form:"menuName"  binding:"required"`
}
