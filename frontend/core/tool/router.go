package tool

import (
	swag "xframe/frontend/core/tool/handler"
	"xframe/pkg/router"
)

func init() {
	// 系统工具
	g1 := router.New("api", "/tool")
	g1.GET("/swagger", "tool:swagger:view", swag.Swagger)
}
