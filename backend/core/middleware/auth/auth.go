package auth

import (
	"net/http"
	"strings"
	"xframe/backend/common/constant"
	indexService "xframe/backend/core/system/index/service"
	menuService "xframe/backend/core/system/menu/service"
	userService "xframe/backend/core/system/user/service"
	"xframe/pkg/router"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(c *gin.Context) {
	// 用户是否被锁
	if isLock := userService.IsLock(c); isLock {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	// 获取用户信息
	user := userService.GetProfile(c)
	menus := menuService.SelectMenuNormalByUser(user.UserId)
	if menus == nil {
		c.Redirect(http.StatusFound, "/500")
		c.Abort()
		return
	}
	// 根据url判断是否有权限
	url := c.Request.URL.Path
	strEnd := url[len(url)-1:]
	if strings.EqualFold(strEnd, "/") {
		url = strings.TrimRight(url, "/")
	}
	// 获取权限标识
	permi := router.FindPermission(url)
	if len(permi) <= 0 {
		c.Next()
		return
	}
	// 检查是否具有权限
	hasPermission := false
	if indexService.CheckPermi(*menus, permi, &hasPermission); hasPermission {
		c.Next()
		return
	}
	// 是否有操作权限
	ajaxString := c.Request.Header.Get("X-Requested-With")
	if strings.EqualFold(ajaxString, "XMLHttpRequest") {
		c.JSON(http.StatusOK, constant.CommonRes{Code: 403, Msg: "您没有操作权限"})
	} else {
		c.Redirect(http.StatusFound, "/403")
	}
	c.Abort()
}
