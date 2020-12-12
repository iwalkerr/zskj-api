package dao

import "time"

type Entity struct {
	UserId      int       `db:"user_id" json:"user_id"`
	Name        string    `db:"name" json:"name"`                 // 用户显示名
	HeadPicture string    `db:"head_picture" json:"head_picture"` // 用户头像
	Username    string    `db:"real_name" json:"real_name"`       // 用户唯一名称
	Password    string    `db:"password"`                         // 加密密码
	Signature   string    `db:"signature" json:"signature"`       // 个性签名
	Sex         string    `db:"sex" json:"sex"`                   // 性别 1男，0女
	Birthday    string    `db:"birthday" json:"birthday"`         // 生日
	CreateTime  time.Time `db:"create_time" json:"create_time"`   // 创建时间
	UpdateTime  time.Time `db:"update_time" json:"update_time"`   // 更新时间
}

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required,min=8"` // 最少8位
}
