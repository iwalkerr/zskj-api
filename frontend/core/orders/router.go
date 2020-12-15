package orders

import (
	orders "xframe/frontend/core/orders/apis"
	"xframe/pkg/router"
)

func init() {
	// 用户订单
	g1 := router.New("api", "/api/v1/orders")
	g1.GET("/", "", orders.List)
}
