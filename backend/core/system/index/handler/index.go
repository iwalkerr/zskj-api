package handler

import (
	"fmt"
	"io/ioutil"
	"os"
	"xframe/backend/common/cfg"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	menuService "xframe/backend/core/system/menu/service"
	userService "xframe/backend/core/system/user/service"
	"xframe/pkg/utils/ip"
	"xframe/pkg/ws"

	"github.com/gin-gonic/gin"
)

// 获取ws连接地址
func GetWs(grpId string, userId int) string {
	ipaddr := cfg.Instance().Admin.WebsocketIp
	if ipaddr == "" {
		ipaddr = ip.GetLocalIP()
	}
	port := cfg.Instance().Admin.Address
	ws := fmt.Sprintf("ws://%s%s/ws/%s/%d", ipaddr, port, grpId, userId)
	return ws
}

// 首页
func Index(c *gin.Context) {
	// param := c.Param("types")

	user := userService.GetProfile(c)
	menuList := menuService.SelectMenuNormalByUser(user.UserId)
	// 去除F菜单
	menuListF := menuService.RemoveAllMenuF(menuList)
	// 切换菜单
	// menuListF = menuService.ChangeMenu(c, menuListF, strconv.Itoa(user.UserId), param)

	// 设置用户头像
	if user.Avatar == "" {
		user.Avatar = "/resource/ajax/libs/ace/images/user.jpg"
	}

	resp.BuildTpl(c, "core/system/layout/index").Write(gin.H{
		"systemName": cfg.Instance().Admin.BusinessName,
		"menuList":   menuListF,
		"user":       user,
		"ws":         GetWs(ws.OnSysUserGrpId, user.UserId),
	})
}

// 默认主页面
func DefaultMain(c *gin.Context) {
	resp.BuildTpl(c, "core/system/layout/default").Write()
}

//下载文件
func Download(c *gin.Context) {
	fileName := c.Query("fileName")
	delete := c.Query("delete")
	if fileName == "" {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}
	// 创建路径
	curDir, err := os.Getwd()
	if err != nil {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "获取目录失败"})
		return
	}
	filepath := curDir + "/public/upload/" + fileName
	file, err := os.Open(filepath)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}
	b, _ := ioutil.ReadAll(file)
	c.Writer.Header().Add("Content-Disposition", "attachment; filename="+fileName)
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	_, _ = c.Writer.Write(b)

	if delete == "true" {
		os.Remove(filepath)
	}
}

func Unauth(c *gin.Context) {
	resp.BuildTpl(c, "core/errpage/unauth").Write()
}

func Error(c *gin.Context) {
	resp.BuildTpl(c, "core/errpage/500").Write()
}

func NotFound(c *gin.Context) {
	resp.BuildTpl(c, "core/errpage/404").Write()
}
