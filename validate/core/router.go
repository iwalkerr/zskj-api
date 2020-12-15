package core

import (
	"xframe/pkg/router"
	"xframe/validate/core/middleware"
	product "xframe/validate/core/product/handler"
)

func init() {
	g1 := router.New("validate", "/validate/v1/check", middleware.Auth)
	g1.GET("/", "", product.List)
	g1.GET("/right", "", product.CheckRight)
}
