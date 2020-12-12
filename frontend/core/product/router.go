package product

import (
	product "xframe/frontend/core/product/apis"
	"xframe/pkg/router"
)

func init() {
	// 商品模块
	g1 := router.New("api", "/api/v1/product")
	g1.GET("/", "", product.List)

}
