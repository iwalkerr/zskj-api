package monitor

import (
	"xframe/backend/core/middleware/auth"
	job "xframe/backend/core/monitor/job/handler"
	loginlog "xframe/backend/core/monitor/loginlog/handler"
	online "xframe/backend/core/monitor/online/handler"
	operlog "xframe/backend/core/monitor/operlog/handler"
	"xframe/backend/core/monitor/server"
	"xframe/pkg/router"
)

func init() {
	// 服务监控
	g1 := router.New("admin", "/monitor/server", auth.Auth)
	g1.GET("/", "monitor:server:view", server.Server)

	// 在线用户
	g2 := router.New("admin", "/monitor/online", auth.Auth)
	g2.GET("/", "monitor:online:view", online.List)
	g2.POST("/list", "monitor:online:list", online.ListAjax)
	g2.POST("/batchForceLogout", "monitor:online:batchForceLogout", online.BatchForceLogout)
	g2.POST("/forceLogout", "monitor:online:forceLogout", online.ForceLogout)

	// 定时任务
	g3 := router.New("admin", "/monitor/job", auth.Auth)
	g3.GET("/", "monitor:job:view", job.List)
	g3.POST("/list", "monitor:job:list", job.ListAjax)
	g3.GET("/add", "monitor:job:add", job.Add)
	g3.POST("/add", "monitor:job:add", job.AddSave)
	g3.GET("/edit", "monitor:job:edit", job.Edit)
	g3.POST("/edit", "monitor:job:edit", job.EditSave)
	g3.POST("/remove", "monitor:job:remove", job.Remove)
	g3.GET("/detail", "monitor:job:detail", job.Detail)
	g3.POST("/start", "monitor:job:changeStatus", job.Start)
	g3.POST("/stop", "monitor:job:changeStatus", job.Stop)

	// 操作日志
	g4 := router.New("admin", "/monitor/operlog", auth.Auth)
	g4.GET("/", "monitor:operlog:view", operlog.List)
	g4.POST("/list", "monitor:operlog:list", operlog.ListAjax)
	g4.POST("/remove", "monitor:operlog:remove", operlog.Remove)
	g4.POST("/clean", "monitor:operlog:remove", operlog.Clean)
	g4.GET("/detail", "monitor:operlog:detail", operlog.Detail)

	//登陆日志
	g5 := router.New("admin", "/monitor/loginlog", auth.Auth)
	g5.GET("/", "monitor:loginlog:view", loginlog.List)
	g5.POST("/list", "monitor:loginlog:list", loginlog.ListAjax)
	g5.POST("/remove", "monitor:loginlog:remove", loginlog.Remove)
	g5.POST("/unlock", "monitor:loginlog:unlock", loginlog.Unlock)
}
