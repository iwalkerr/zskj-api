package model

import (
	"time"
	"xframe/backend/common/constant"
)

type Entity struct {
	UserId      int       `db:"user_id" json:"user_id"`         // 用户ID
	DeptId      int       `db:"dept_id" json:"dept_id"`         // 部门ID
	LoginName   string    `db:"login_name" json:"login_name"`   // 登录账号
	UserName    string    `db:"user_name" json:"user_name"`     // 用户昵称
	Email       string    `db:"email" json:"email"`             // 用户邮箱
	PhoneNumber string    `db:"phonenumber" json:"phonenumber"` // 手机号码
	Sex         string    `db:"sex" json:"sex"`                 // 用户性别（1男 0女 2未知）
	Avatar      string    `db:"avatar" json:"avatar"`           // 头像路径
	Status      string    `db:"status" json:"status"`           // 帐号状态（0正常 1停用）
	UserType    string    `db:"user_type" json:"-"`             // 用户类型（00系统用户）
	DelFlag     string    `db:"del_flag" json:"-"`              // 删除标志（0代表存在 2代表删除）
	LoginIp     string    `db:"login_ip" json:"login_ip"`       // 最后登陆IP
	LoginDate   time.Time `db:"login_date" json:"login_date"`   // 最后登陆时间
	DeptName    string    `db:"dept_name" json:"dept_name"`     // 部门名称
	Leader      string    `db:"leader" json:"leader"`           // 领导
	Password    string    `json:"-"`                            // 密码
	Salt        string    `json:"-"`                            // 盐加密
	// SessionId   string    `json:"-"`                            // sessionId
	constant.ModelData
}

//查询用户列表请求参数
type SelectPageReq struct {
	LoginName   string `form:"loginName"`   //登陆名
	Status      string `form:"status"`      //状态
	Phonenumber string `form:"phonenumber"` //手机号码
	DeptId      int    `form:"deptId"`      // 部门ID
	*constant.PageReq
}

// 用户与角色
type UserRoleEntity struct {
	UserId int `json:"user_id"` // 用户ID
	RoleId int `json:"role_id"` // 角色ID
}

//新增用户资料请求参数
type EditReq struct {
	UserId      int    `form:"userId" binding:"required"`
	UserName    string `form:"userName"  binding:"required"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
	Email       string `form:"email"  binding:"required,email"`
	DeptId      int    `form:"deptId" binding:"required"`
	Sex         string `form:"sex"  binding:"required"`
	Status      string `form:"status"`
	RoleIds     string `form:"roleIds"`
	PostIds     string `form:"postIds"`
	Remark      string `form:"remark"`
}

//新增用户资料请求参数
type AddReq struct {
	LoginName   string `form:"loginName" binding:"required,min=5,max=30"`
	UserName    string `form:"userName" binding:"required"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
	Email       string `form:"email"  binding:"required,email"`
	Password    string `form:"password"  binding:"required,min=5,max=30"`
	DeptId      int    `form:"deptId" binding:"required"`
	Sex         string `form:"sex"  binding:"required"`
	Status      string `form:"status"`
	RoleIds     string `form:"roleIds"`
	PostIds     string `form:"postIds"`
	Remark      string `form:"remark"`
}

//重置密码请求参数
type ResetPwdReq struct {
	UserId   int    `form:"userId"  binding:"required"`
	Password string `form:"password" binding:"required,min=5,max=30"`
}

// 用户与岗位
type UserPostEntity struct {
	UserId int `json:"user_id" binding:"required"` // 用户ID
	PostId int `json:"post_id" binding:"required"` // 岗位ID
}

//检查email请求参数
type CheckEmailAllReq struct {
	Email string `form:"email"  binding:"required,email"`
}

//检查phone请求参数
type CheckPhoneAllReq struct {
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
}

//检查phone请求参数
type CheckLoginNameReq struct {
	LoginName string `form:"loginName"  binding:"required"`
}

//用户列表数据结构
type UserListEntity struct {
	Entity
	DeptName string `json:"dept_name"` // 部门名称
	Leader   string `json:"leader"`    // 负责人
}

//检查email请求参数
type CheckEmailReq struct {
	UserId int    `form:"userId"  binding:"required"`
	Email  string `form:"email"  binding:"required,email"`
}

//检查phone请求参数
type CheckPhoneReq struct {
	UserId      int    `form:"userId"  binding:"required"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
}

//修改用户资料请求参数
type ProfileReq struct {
	UserName    string `form:"userName"  binding:"required,min=5,max=30"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
	Email       string `form:"email"  binding:"required,email"`
	Sex         string `form:"sex"  binding:"required"`
}

//检查密码请求参数
type CheckPasswordReq struct {
	Password string `form:"password"  binding:"required"`
}

//修改密码请求参数
type PasswordReq struct {
	OldPassword string `form:"oldPassword" binding:"required"`
	NewPassword string `form:"newPassword" binding:"required,min=5,max=30"`
	Confirm     string `form:"confirm" binding:"required,min=5,max=30"`
}

type ToSession struct {
	UserId      int       `db:"user_id" json:"user_id"`         // 用户ID
	DeptId      int       `db:"dept_id" json:"dept_id"`         // 部门ID
	LoginName   string    `db:"login_name" json:"login_name"`   // 登录账号
	UserName    string    `db:"user_name" json:"user_name"`     // 用户昵称
	Email       string    `db:"email" json:"email"`             // 用户邮箱
	PhoneNumber string    `db:"phonenumber" json:"phonenumber"` // 手机号码
	Sex         string    `db:"sex" json:"sex"`                 // 用户性别（1男 0女 2未知）
	Avatar      string    `db:"avatar" json:"avatar"`           // 头像路径
	Status      string    `db:"status" json:"status"`           // 帐号状态（0正常 1停用）
	CreateTime  time.Time `db:"create_time" json:"create_time"` // 创建时间
}
