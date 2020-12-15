package service

import (
	"errors"
	"strings"
	"time"
	"xframe/backend/core/system/config/model"
	userService "xframe/backend/core/system/user/service"

	"github.com/gin-gonic/gin"
)

// 获取资源路径
func GetOssUrl() string {
	return model.GetValueByKey("sys.resource.url")
}

func SelectListByPage(req *model.SelectPageReq) []model.Entity {
	return model.SelectListByPage(req)
}

//根据主键查询数据
func SelectRecordById(id int) model.Entity {
	return model.SelectRecordById(id)
}

//校验参数键名是否唯一
func CheckConfigKeyUnique(configKey string, configId int) string {
	entity := model.CheckConfigKeyUniqueAll(configKey)
	if entity > 0 && entity != configId {
		return "1"
	}
	return "0"
}

//修改数据
func EditSave(c *gin.Context, req *model.EditReq) (int, error) {
	entity := model.SelectRecordById(req.ConfigId)
	if entity.ConfigId <= 0 {
		return 0, errors.New("数据不存在")
	}

	entity.ConfigName = req.ConfigName
	entity.ConfigKey = req.ConfigKey
	entity.ConfigValue = req.ConfigValue
	entity.Remark = req.Remark
	entity.ConfigType = req.ConfigType
	entity.UpdateTime = time.Now()

	user := userService.GetProfile(c)
	if user != nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Update()
}

//检查角色名是否唯一
func CheckConfigKeyUniqueAll(configKey string) string {
	id := model.CheckConfigKeyUniqueAll(configKey)
	if id > 0 {
		return "1"
	}
	return "0"
}

// 根据ID批量删除数据
func DeleteRecordByIds(ids string) error {
	return model.DeleteBatch(strings.Split(ids, ","))
}

//添加数据
func AddSave(c *gin.Context, req *model.AddReq) (int, error) {
	var entity model.Entity
	entity.ConfigName = req.ConfigName
	entity.ConfigKey = req.ConfigKey
	entity.ConfigType = req.ConfigType
	entity.ConfigValue = req.ConfigValue
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()

	user := userService.GetProfile(c)
	if user != nil {
		entity.CreateBy = user.LoginName
	}

	return entity.Insert()
}
