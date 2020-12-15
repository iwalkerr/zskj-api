package service

import (
	userModel "xframe/backend/core/system/user/model"
	"xframe/pkg/cache"
)

const USER_NOPASS_TIME string = "user_nopass_"
const USER_LOCK string = "user_lock_"

// 检查账号是否锁定
func CheckLock(loginName string) (isLock bool) {
	rs, _ := cache.Instance().Get(USER_LOCK + loginName)
	if rs != nil {
		isLock = true
	}
	return isLock
}

// 用户是否已经被锁
func IsUserLock(loginName string) error {
	return userModel.IsUserLock(loginName)
}
