package core

import (
	orders "xframe/frontend/core/orders/handler"
	product "xframe/frontend/core/product/handler"
	seller "xframe/frontend/core/seller/handler"
	shoppingcart "xframe/frontend/core/shoppingcart/handler"
	teams "xframe/frontend/core/teams/handler"
	_ "xframe/frontend/core/tool"
	user "xframe/frontend/core/user/handler"
	"xframe/pkg/router"
)

func init() {
	// 用户路由
	g1 := router.New("api", "/api/v1/user")
	g1.GET("/", "", user.List)

	// 用户订单
	g2 := router.New("api", "/api/v1/orders")
	g2.GET("/", "", orders.List)

	// 商品模块
	g3 := router.New("api", "/api/v1/product")
	g3.GET("/", "", product.List)

	// 卖家模块
	g4 := router.New("api", "/api/v1/seller")
	g4.GET("/", "", seller.List)

	// 购物车模块
	g5 := router.New("api", "/api/v1/shoppingcart")
	g5.GET("/", "", shoppingcart.List)

	// 群聊模块
	g6 := router.New("api", "/api/v1/teams")
	g6.GET("/", "", teams.List)
}
