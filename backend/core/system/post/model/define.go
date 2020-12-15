package model

import (
	"time"
	"xframe/backend/common/constant"
)

type Entity struct {
	PostId     int       `json:"post_id"`     // 岗位ID
	PostCode   string    `json:"post_code"`   // 岗位编码
	PostName   string    `json:"post_name"`   // 岗位名称
	PostSort   int       `json:"post_sort"`   // 显示顺序
	Status     string    `json:"status"`      // 状态（0正常 1停用)
	CreateBy   string    `json:"create_by"`   // 创建者
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateBy   string    `json:"update_by"`   // 更新者
	UpdateTime time.Time `json:"update_time"` // 更新时间
	Remark     string    `json:"remark"`      // 备注
	Flag       bool      `json:"flag"`        // 标记
}

//分页请求参数
type SelectPageReq struct {
	PostCode string `form:"postCode"` //岗位编码
	Status   string `form:"status"`   //状态
	PostName string `form:"postName"` //岗位名称
	*constant.PageReq
}

//检查名称请求参数
type CheckPostNameAllReq struct {
	PostName string `form:"postName"  binding:"required"`
}

//检查编码请求参数
type CheckPostCodeAllReq struct {
	PostCode string `form:"postCode"  binding:"required"`
}

//新增页面请求参数
type AddReq struct {
	PostName string `form:"postName"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
	PostSort int    `form:"postSort"  binding:"required"`
	Status   string `form:"status"    binding:"required"`
	Remark   string `form:"remark"`
}

//修改页面请求参数
type EditReq struct {
	PostId int `form:"postId" binding:"required"`
	AddReq
}

//检查名称请求参数
type CheckPostNameReq struct {
	PostId   int    `form:"postId"  binding:"required"`
	PostName string `form:"postName"  binding:"required"`
}

//检查编码请求参数
type CheckPostCodeReq struct {
	PostId   int    `form:"postId"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
}
