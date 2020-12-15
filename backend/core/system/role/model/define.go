package model

import (
	"time"
	"xframe/backend/common/constant"
)

type Entity struct {
	RoleId     int       `json:"role_id"`     // 角色ID
	RoleName   string    `json:"role_name"`   // 角色名称
	RoleKey    string    `json:"role_key"`    // 角色权限字符串
	RoleSort   int       `json:"role_sort"`   // 显示顺序
	DataScope  string    `json:"data_scope"`  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     string    `json:"status"`      // 角色状态（0正常 1停用）
	DelFlag    string    `json:"del_flag"`    // 删除标志（0代表存在 2代表删除）
	CreateBy   string    `json:"create_by"`   // 创建者
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateBy   string    `json:"update_by"`   // 更新者
	UpdateTime time.Time `json:"update_time"` // 更新时间
	Remark     string    `json:"remark"`      // 备注
	Flag       bool      `json:"flag" `       // 标记
}

//分页请求参数
type SelectPageReq struct {
	RoleName  string `form:"roleName"`  //角色名称
	Status    string `form:"status"`    //状态
	RoleKey   string `form:"roleKey"`   //角色键
	DataScope string `form:"dataScope"` //数据范围
	*constant.PageReq
}

//检查角色名称请求参数
type CheckRoleNameAllReq struct {
	RoleName string `form:"roleName"  binding:"required"`
}

//检查权限字符请求参数
type CheckRoleKeyAllReq struct {
	RoleKey string `form:"roleKey"  binding:"required"`
}

//新增页面请求参数
type AddReq struct {
	RoleName string `form:"roleName"  binding:"required"`
	RoleKey  string `form:"roleKey"  binding:"required"`
	RoleSort int    `form:"roleSort"  binding:"required"`
	Status   string `form:"status"`
	Remark   string `form:"remark"`
	MenuIds  string `form:"menuIds"`
}

//修改页面请求参数
type EditReq struct {
	RoleId   int    `form:"roleId" binding:"required"`
	RoleName string `form:"roleName"  binding:"required"`
	RoleKey  string `form:"roleKey"  binding:"required"`
	RoleSort string `form:"roleSort"  binding:"required"`
	Status   string `form:"status"`
	Remark   string `form:"remark"`
	MenuIds  string `form:"menuIds"`
}

type EntityRoleMenu struct {
	RoleId int `json:"role_id"` // 角色ID
	MenuId int `json:"menu_id"` // 菜单ID
}

// 用户与角色
type UserRoleEntity struct {
	RoleId  int    `form:"roleId" binding:"required"`  // 角色ID
	UserIds string `form:"userIds" binding:"required"` // 用户ID
}

type EntityRoleDept struct {
	RoleId int `json:"role_id"` // 角色ID
	DeptId int `json:"dept_id"` // 部门ID
}

//检查角色名称请求参数
type CheckRoleNameReq struct {
	RoleId   int    `form:"roleId"  binding:"required"`
	RoleName string `form:"roleName"  binding:"required"`
}

//检查权限字符请求参数
type CheckRoleKeyReq struct {
	RoleId  int    `form:"roleId"  binding:"required"`
	RoleKey string `form:"roleKey"  binding:"required"`
}

//数据权限保存请求参数
type DataScopeReq struct {
	RoleId    int    `form:"roleId"  binding:"required"`
	RoleName  string `form:"roleName"  binding:"required"`
	RoleKey   string `form:"roleKey"  binding:"required"`
	DataScope string `form:"dataScope"  binding:"required"`
	DeptIds   string `form:"deptIds"`
}

type AllocatedReq struct {
	RoleId        int    `form:"roleId"  binding:"required"`
	LoginName     string `form:"loginName"`
	PhoneNumber   string `form:"phonenumber"`
	OrderByColumn string `form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc"`         //排序方式
}

// 改变用户可用状态
type ChangeStatusReq struct {
	RoleId int    `form:"roleId"  binding:"required"`
	Status string `form:"status"  binding:"required"`
}
