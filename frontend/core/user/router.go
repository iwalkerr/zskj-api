package user

import (
	"xframe/frontend/core/middleware/auth"
	user "xframe/frontend/core/user/apis"
	"xframe/pkg/router"
)

func init() {
	// 用户路由
	g1 := router.New("api", "/api/v1/user")
	g1.POST("/login", "", user.Login)       // 用户登陆
	g1.POST("/register", "", user.Register) // 用户注册

	g1.POST("/head_pic", "", auth.Auth, user.HeadPic) // 测试
}
