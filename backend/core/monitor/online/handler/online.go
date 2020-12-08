package handler

import (
	"fmt"
	"strings"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	"xframe/backend/core/monitor/online/model"
	index "xframe/backend/core/system/index/handler"
	userService "xframe/backend/core/system/user/service"
	"xframe/pkg/utils/base"
	"xframe/pkg/ws"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	user := userService.GetProfile(c)

	resp.BuildTpl(c, "core/monitor/online/list").Write(gin.H{
		"ws":     index.GetWs(ws.OnSysUserGrpId, user.UserId),
		"userId": user.UserId,
	})
}

func ListAjax(c *gin.Context) {
	var req model.SelectPageReq
	if err := c.ShouldBind(&req); err != nil {
		total := 0
		resp.BuildTable(c, &constant.PageReq{Total: &total}, nil).Write()
		return
	}
	if req.Ids == "" {
		total := 0
		resp.BuildTable(c, &constant.PageReq{Total: &total}, nil).Write()
		return
	}

	idsArr := strings.Split(req.Ids, ",")
	req.IdsArr = base.RemoveRepByMap(idsArr)

	// 根据Ids查询用户列表
	list := userService.ListUsersByIds(&req)
	resp.BuildTable(c, req.PageReq, list).Write()
}

//用户强退
func ForceLogout(c *gin.Context) {
	userId := c.PostForm("id")
	if userId == "" {
		resp.Error(c).Msg("缺少参数").Log("用户强退", gin.H{"v": userId}).Write()
		return
	}
	// userService.ForceLogout(userId)
	resp.Success(c).Log("用户强退", gin.H{"userId": userId}).Write()
}

//批量强退
func BatchForceLogout(c *gin.Context) {
	ids := c.PostForm("ids")
	if ids == "" {
		resp.Error(c).Msg("缺少参数").Log("批量强退", gin.H{"ids": ids}).Write()
		return
	}
	if idarr := strings.Split(ids, ","); len(idarr) > 0 {
		for _, userId := range idarr {
			if userId != "" {
				fmt.Println(userId)
				// userService.ForceLogout(userId)
			}
		}
	}
	resp.Success(c).Log("批量强退", gin.H{"ids": ids}).Write()
}
