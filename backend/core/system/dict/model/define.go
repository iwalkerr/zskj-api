package model

import "xframe/backend/common/constant"

type Entity struct {
	DictId    int    `db:"dict_id" json:"dict_id"`       // 字典ID
	DictType  string `db:"dict_type" json:"dict_type"`   // 字典类型
	DictSort  int    `db:"dict_sort" json:"dict_sort"`   // 字典顺序
	DictLabel string `db:"dict_label" json:"dict_label"` // 字典名称
	DictValue string `db:"dict_value" json:"dict_value"` // 字典值
	ParentId  int    `db:"parent_id" json:"parent_id"`   // 父字典ID
	ListClass string `db:"list_class" json:"list_class"` // class属性
	IsDefault string `db:"is_default" json:"is_default"` // 是否是默认
	Status    string `db:"status" json:"status"`         // 字典状态
	constant.ModelData
}

//查询用户列表请求参数
type SelectPageReq struct {
	DictId    string `form:"dictId"` // 字典ID
	Status    string `form:"status"`
	DictLabel string `form:"dictLabel"`
	*constant.PageReq
}

type DictData struct {
	DictId    int    `db:"dict_id" json:"dict_id"`       // 字典ID
	DictLabel string `db:"dict_label" json:"dict_label"` // 字典名称
	DictValue string `db:"dict_value" json:"dict_value"` // 字典值
	ListClass string `db:"list_class" json:"list_class"` // class属性
}

//新增页面请求参数
type AddReq struct {
	ParentId  int    `form:"parentId"` // 父字典ID
	DictLabel string `form:"dictLabel"  binding:"required"`
	DictValue string `form:"dictValue"  binding:"required"`
	DictType  string `form:"dictType"  binding:"required"`
	DictSort  int    `form:"dictSort"  binding:"required"`
	ListClass string `form:"listClass" binding:"required"`
	IsDefault string `form:"isDefault" binding:"required"`
	Status    string `form:"status"    binding:"required"`
	Remark    string `form:"remark"`
}

//修改页面请求参数
type EditReq struct {
	DictId int `form:"dictId" binding:"required"`
	AddReq
}
