package teams

import (
	teams "xframe/frontend/core/teams/apis"
	"xframe/pkg/router"
)

func init() {
	// 群聊模块
	g1 := router.New("api", "/api/v1/teams")
	g1.GET("/", "", teams.List)
}
