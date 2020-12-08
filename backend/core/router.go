package core

import (
	"xframe/backend/core/middleware/auth"
	_ "xframe/backend/core/monitor"
	_ "xframe/backend/core/system"
	index "xframe/backend/core/system/index/handler"
	user "xframe/backend/core/system/user/handler"
	"xframe/pkg/router"
	"xframe/pkg/ws"
)

func init() {
	g1 := router.New("admin", "/")
	g1.GET("/ws/:grpId/:uid", "", ws.RunWsServer) // websocket 实时通讯,grpId没有群或组为0
	g1.GET("/login", "", index.Login)
	g1.GET("/logout", "", index.Logout)
	g1.GET("/captchaImage", "", index.CaptchaImage)
	g1.POST("/checklogin", "", index.CheckLogin)
	g1.GET("/index", "", auth.Auth, index.Index)
	g1.GET("/500", "", index.Error)
	g1.GET("/404", "", index.NotFound)
	g1.GET("/403", "", index.Unauth)

	// 系统路由
	g2 := router.New("admin", "/system", auth.Auth)
	g2.GET("/download", "", index.Download)
	g2.GET("/default_main", "", auth.Auth, index.DefaultMain)

	// 个人中心路由
	g8 := router.New("admin", "/system/user/profile", auth.Auth)
	g8.GET("/", "", user.Profile)
	g8.POST("/checkEmailUnique", "", user.CheckEmailUnique)
	g8.POST("/checkPhoneUnique", "", user.CheckPhoneUnique)
	g8.POST("/checkPhoneUniqueAll", "", user.CheckPhoneUniqueAll)
	g8.POST("/checkLoginNameUnique", "", user.CheckLoginNameUnique)
	g8.POST("/checkEmailUniqueAll", "", user.CheckEmailUniqueAll)
	g8.GET("/avatar", "", user.Avatar)
	g8.POST("/updateAvatar", "", user.UpdateAvatar)
	g8.POST("/update", "", user.Update)
	g8.POST("/checkPassword", "", user.CheckPassword)
	g8.POST("/resetSavePwd", "", user.UpdatePassword)
}
