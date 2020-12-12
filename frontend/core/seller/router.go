package seller

import (
	seller "xframe/frontend/core/seller/apis"
	"xframe/pkg/router"
)

func init() {
	// 卖家模块
	g1 := router.New("api", "/api/v1/seller")
	g1.GET("/", "", seller.List)
}
