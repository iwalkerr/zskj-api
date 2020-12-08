package service

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
	"xframe/backend/common/constant"
	"xframe/backend/core/monitor/operlog/model"
	deptService "xframe/backend/core/system/dept/service"
	userService "xframe/backend/core/system/user/service"
	"xframe/pkg/utils/ip"

	"github.com/gin-gonic/gin"
)

func SelectPageList(req *model.SelectPageReq) []model.Entity {
	return model.SelectPageList(req)
}

//批量删除记录
func DeleteRecordByIds(ids string) int {
	idarr := strings.Split(ids, ",")
	return model.DeleteBatch(idarr)
}

//清空记录
func DeleteRecordAll() error {
	return model.DeleteRecordAll()
}

//根据主键查询用户信息
func SelectRecordById(id int) model.Entity {
	return model.SelectRecordById(id)
}

//新增记录
func Add(c *gin.Context, title, inContent string, outContent *constant.CommonRes) error {
	user := userService.GetProfile(c)
	if user == nil {
		return errors.New("用户未登陆")
	}

	var operl model.Entity
	outJson, _ := json.Marshal(outContent)
	operl.JsonResult = string(outJson)
	operl.Title = title
	operl.OperParam = inContent
	operl.BusinessType = int(outContent.Btype)
	//操作类别（0其它 1后台用户 2手机端用户）
	operl.OperatorType = 1
	//操作状态（0正常 1异常）
	if outContent.Code == 0 {
		operl.Status = 0
	} else {
		operl.Status = 1
	}

	operl.OperName = user.LoginName
	operl.RequestMethod = c.Request.Method

	dept := deptService.SelectDeptById(user.DeptId)
	if dept.DeptId > 0 {
		operl.DeptName = dept.DeptName
	} else {
		operl.DeptName = ""
	}

	operl.OperUrl = c.Request.URL.Path
	operl.Method = c.Request.Method
	operl.OperIp = c.ClientIP()

	operl.OperLocation = ip.GetCityByIP(operl.OperIp)
	operl.OperTime = time.Now()

	return operl.Insert()
}
