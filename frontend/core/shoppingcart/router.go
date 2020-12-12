package shoppingcart

import (
	shoppingcart "xframe/frontend/core/shoppingcart/apis"
	"xframe/pkg/router"
)

func init() {
	// 购物车模块
	g1 := router.New("api", "/api/v1/shoppingcart")
	g1.GET("/", "", shoppingcart.List)

}
