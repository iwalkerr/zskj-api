package dao

import "time"

type Entity struct {
	UserId      int       `db:"user_id" json:"user_id"`
	Name        string    `db:"name" json:"name"`                 // 用户显示名
	HeadPicture string    `db:"head_picture" json:"head_picture"` // 用户头像
	Username    string    `db:"real_name" json:"real_name"`       // 用户唯一名称
	Password    string    `db:"password" json:"-"`                // 加密密码
	Signature   string    `db:"signature" json:"signature"`       // 个性签名
	Sex         string    `db:"sex" json:"sex"`                   // 性别 1男，0女
	Birthday    string    `db:"birthday" json:"birthday"`         // 生日
	Phone       string    `db:"phone" json:"phone"`               // 手机号
	CreateTime  time.Time `db:"created_time" json:"created_time"` // 创建时间
	UpdateTime  time.Time `db:"updated_time" json:"updated_time"` // 更新时间
}

type LoginResp struct {
	Name        string `db:"name" json:"name"`
	HeadPicture string `db:"head_picture" json:"head_picture"`
	Username    string `db:"real_name" json:"real_name"`
	Signature   string `db:"signature" json:"signature"`
	Sex         string `db:"sex" json:"sex"`
	Phone       string `db:"phone" json:"phone"`
	Birthday    string `db:"birthday" json:"birthday"`
	Password    string `db:"password" json:"-"`
	UserId      int    `db:"user_id" json:"-"`
}

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required,min=8"` // 最少8位
}

type Register struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required,min=8"`
}
