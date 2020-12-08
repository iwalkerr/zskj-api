package service

import (
	"log"
	"strings"
	"time"
	"xframe/backend/core/monitor/loginlog/model"
	"xframe/pkg/cache"
	"xframe/pkg/utils/ip"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

const USER_NOPASS_TIME string = "user_nopass_"
const USER_LOCK string = "user_lock_"

//移除密码错误次数
func RemovePasswordCounts(loginName string) {
	cache.Instance().Delete(USER_NOPASS_TIME + loginName)
}

//批量删除记录
func DeleteRecordByIds(ids string) int {
	return model.DeleteRecordByIds(strings.Split(ids, ","))
}

func SelectPageList(req *model.SelectPageReq) []model.Entity {
	return model.SelectPageList(req)
}

// 记录日志
func AddUserLog(c *gin.Context, username, msg, status, sessionId string, userId int, isLogin bool) {
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	loginIp := c.ClientIP()
	browser, _ := ua.Browser()
	loginLocation := ip.GetCityByIP(loginIp)
	os := ua.OS()

	// 记录日志
	var userlog model.Entity
	userlog.LoginName = username
	userlog.Ipaddr = loginIp
	userlog.Os = os
	userlog.Browser = browser
	userlog.LoginTime = time.Now()
	userlog.LoginLocation = loginLocation
	userlog.Msg = msg
	userlog.Status = status
	if err := userlog.Insert(); err != nil {
		log.Println("日志记录失败,", err)
	}
}

//记录密码尝试次数
func SetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, _ := cache.Instance().Get(USER_NOPASS_TIME + loginName)
	if curTimeObj != nil {
		if count, ok := curTimeObj.(int); ok {
			curTimes = count
		}
	}
	curTimes = curTimes + 1
	cache.Instance().Set(USER_NOPASS_TIME+loginName, curTimes, 1*time.Minute)

	if curTimes >= 4 {
		Lock(loginName)
	}
	return curTimes
}

//锁定账号
func Lock(loginName string) {
	cache.Instance().Set(USER_LOCK+loginName, true, 30*time.Minute)
}

//解除锁定
func Unlock(loginName string) {
	cache.Instance().Delete(USER_LOCK + loginName)
}
