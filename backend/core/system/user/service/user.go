package service

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
	"xframe/backend/common/constant"
	"xframe/backend/common/db"
	"xframe/backend/core/middleware/sessions"
	onlineModel "xframe/backend/core/monitor/online/model"
	role "xframe/backend/core/system/role/model"
	"xframe/backend/core/system/user/model"
	"xframe/pkg/cache"
	"xframe/pkg/utils/excel"
	"xframe/pkg/utils/gmd5"
	"xframe/pkg/utils/random"

	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	// 退出登陆,清理工作
	user := GetProfile(c)
	if user != nil {
		ClearMenuCache(user)
	}

	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
}

func ListUsersByIds(req *onlineModel.SelectPageReq) []model.Entity {
	return model.ListUsersByIds(req)
}

func UpdatePassword(c *gin.Context, profile *model.PasswordReq) error {
	to := GetProfile(c)

	var user model.Entity
	user.UserId = to.UserId
	user.DeptId = to.DeptId
	user.LoginName = to.LoginName
	user.UserName = to.UserName
	user.Email = to.Email
	user.PhoneNumber = to.PhoneNumber
	user.Sex = to.Sex
	user.Avatar = to.Avatar
	user.Status = to.Status

	if err := user.FindUser(); err != nil {
		return err
	}

	//校验密码
	oldPassword := gmd5.MustEncryptString(user.LoginName + profile.OldPassword)
	token := user.LoginName + oldPassword + user.Salt
	token = gmd5.MustEncryptString(token)

	if token != user.Password {
		return errors.New("原密码不正确")
	}

	newPassword := gmd5.MustEncryptString(user.LoginName + profile.NewPassword)

	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := user.LoginName + newPassword + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken

	// 更新数据
	if err := user.UpdatePwd(); err != nil {
		return errors.New("保存数据失败")
	}

	to.UserId = user.UserId
	to.DeptId = user.DeptId
	to.LoginName = user.LoginName
	to.UserName = user.UserName
	to.Email = user.Email
	to.PhoneNumber = user.PhoneNumber
	to.Sex = user.Sex
	to.Avatar = user.Avatar
	to.Status = user.Status

	SaveUserToSession(c, to)
	return nil
}

func CheckPassword(user *model.Entity, password string) bool {
	if err := user.FindUser(); err != nil {
		return false
	}
	password = gmd5.MustEncryptString(user.LoginName + password)

	//校验密码
	token := user.LoginName + password + user.Salt
	token = gmd5.MustEncryptString(token)

	if strings.Compare(token, user.Password) == 0 {
		return true
	} else {
		return false
	}
}

func UpdateProfile(c *gin.Context, profile *model.ProfileReq) error {
	user := GetProfile(c)
	user.UserName = profile.UserName
	user.Email = profile.Email
	user.PhoneNumber = profile.Phonenumber
	user.Sex = profile.Sex

	if err := user.UpdateData(); err != nil {
		return errors.New("保存数据失败")
	}

	SaveUserToSession(c, user)
	return nil
}

// 更新用户头像
func UpdateAvatar(c *gin.Context, avatar string) error {
	user := GetProfile(c)
	if avatar != "" {
		user.Avatar = avatar
	}

	if err := user.Update(); err != nil {
		return errors.New("保存数据失败")
	}

	SaveUserToSession(c, user)
	return nil
}

// 导出excel
func Export(req *model.SelectPageReq) (string, error) {
	userList := model.SelectExportList(req)
	if len(userList) == 0 {
		return "", errors.New("没有查到数据")
	}

	file, err := excel.NewFile()
	if err != nil {
		return "", errors.New("数据导出错误")
	}

	head := []interface{}{"用户名", "呢称", "Email", "电话号码", "性别", "部门", "领导", "状态", "删除标记", "创建人", "创建时间", "备注"}
	file.SetRow(1, head)

	for i, user := range userList {
		createTime := user.CreateTime.Format("2006-01-02 15:04:05")
		file.SetRow(i+2, []interface{}{user.LoginName, user.UserName, user.Email, user.PhoneNumber, user.Sex, user.DeptName, user.Leader, user.Status, user.DelFlag, user.CreateBy, createTime, user.Remark})
	}

	return file.SaveAs("系统用户表")
}

// 重置密码
func ResetPassword(req *model.ResetPwdReq) error {
	u := model.SelectRecordById(req.UserId)
	if u.UserId <= 0 {
		return errors.New("用户不存在")
	}

	password := gmd5.MustEncryptString(u.LoginName + req.Password)
	//新校验密码
	newSalt := random.GenerateSubId(6)
	newToken := u.LoginName + password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	u.Salt = newSalt
	u.Password = newToken

	if err := u.UpdatePwd(); err != nil {
		return errors.New("保存数据失败")
	}
	return nil
}

//批量删除用户记录
func DeleteRecordByIds(ids string) error {
	idarr := strings.Split(ids, ",")

	tx, err := db.Conn().Begin()
	if err != nil {
		return err
	}
	// 删除用户
	if err = model.DeleteRecordByIds(tx, idarr); err != nil {
		return errors.New("用户删除失败")
	}

	// 删除用户角色
	if err = model.DeleteUserRoleByIds(tx, idarr); err != nil {
		return errors.New("用户角色删除失败")
	}

	// 删除用户岗位
	if err = model.DeleteUserPostsByIds(tx, idarr); err != nil {
		return errors.New("用户岗位删除失败")
	}

	return tx.Commit()
}

// 新增用户
func AddSave(c *gin.Context, req *model.AddReq) (int, error) {
	var u model.Entity
	u.LoginName = req.LoginName
	u.UserName = req.UserName
	u.Email = req.Email
	u.PhoneNumber = req.Phonenumber
	u.Status = req.Status
	u.Sex = req.Sex
	u.DeptId = req.DeptId
	u.Remark = req.Remark

	//生成密码
	newSalt := random.GenerateSubId(6)
	newToken := req.LoginName + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	u.Salt = newSalt
	u.Password = newToken

	u.CreateTime = time.Now()
	createUser := GetProfile(c)

	if createUser != nil {
		u.CreateBy = createUser.LoginName
	}
	u.DelFlag = "0"

	tx, _ := db.Conn().Begin()

	// 插入用户信息
	err := u.AddUser(tx)
	if err != nil {
		return 0, err
	}

	// 添加岗位数据
	if req.PostIds != "" {
		postIds := strings.Split(req.PostIds, ",")
		userPosts := make([]model.UserPostEntity, 0)
		for i := range postIds {
			postId, _ := strconv.Atoi(postIds[i])
			if postId > 0 {
				var userPost model.UserPostEntity
				userPost.UserId = u.UserId
				userPost.PostId = postId
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			err := model.AddUserPost(tx, userPosts)
			if err != nil {
				return 0, err
			}
		}
	}

	// 添加角色数据
	if req.RoleIds != "" {
		roleIds := strings.Split(req.RoleIds, ",")
		userRoles := make([]model.UserRoleEntity, 0)
		for i := range roleIds {
			roleId, _ := strconv.Atoi(roleIds[i])
			if roleId > 0 {
				var userRole model.UserRoleEntity
				userRole.UserId = u.UserId
				userRole.RoleId = roleId
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			err := model.AddUserRole(tx, userRoles)
			if err != nil {
				return 0, err
			}
		}
	}

	return u.UserId, tx.Commit()
}

//检查邮箱是否存在,存在返回true,否则false
func CheckEmailUniqueAll(email string) bool {
	return model.CheckEmailUniqueAll(email)
}

// 检查登陆名是否存在,存在返回true,否则false
func CheckLoginName(loginName string) bool {
	return model.CheckLoginName(loginName)
}

//检查手机号是否已使用 ,存在返回true,否则false
func CheckPhoneUniqueAll(phoneNumber string) bool {
	return model.CheckPhoneUniqueAll(phoneNumber)
}

func EditSave(c *gin.Context, req *model.EditReq) (int, error) {
	u := model.SelectRecordById(req.UserId)
	if u.UserId <= 0 {
		return 0, errors.New("数据不存在")
	}

	u.UserName = req.UserName
	u.Email = req.Email
	u.PhoneNumber = req.Phonenumber
	u.Status = req.Status
	u.Sex = req.Sex
	u.DeptId = req.DeptId
	u.Remark = req.Remark

	u.UpdateTime = time.Now()

	updateUser := GetProfile(c)
	if updateUser != nil {
		u.UpdateBy = updateUser.LoginName
	}

	// 开启事务
	tx, _ := db.Conn().Begin()
	if err := u.UpdateSave(tx); err != nil {
		return 0, err
	}

	// 更新岗位
	if req.PostIds != "" {
		postIds := strings.Split(req.PostIds, ",")
		userPosts := make([]model.UserPostEntity, 0)
		for i := range postIds {
			postId, _ := strconv.Atoi(postIds[i])
			if postId > 0 {
				var userPost model.UserPostEntity
				userPost.UserId = u.UserId
				userPost.PostId = postId
				userPosts = append(userPosts, userPost)
			}
		}
		if len(userPosts) > 0 {
			err := model.DelAddUserPosts(tx, userPosts, u.UserId)
			if err != nil {
				return 0, err
			}
		}
	} else {
		err := model.DelAddUserPosts(tx, []model.UserPostEntity{}, u.UserId)
		if err != nil {
			return 0, err
		}
	}

	// 更新角色数据
	if req.RoleIds != "" {
		roleIds := strings.Split(req.RoleIds, ",")
		userRoles := make([]model.UserRoleEntity, 0)
		for i := range roleIds {
			roleId, _ := strconv.Atoi(roleIds[i])
			if roleId > 0 {
				var userRole model.UserRoleEntity
				userRole.UserId = u.UserId
				userRole.RoleId = roleId
				userRoles = append(userRoles, userRole)
			}
		}
		if len(userRoles) > 0 {
			err := model.DelAddUserRoles(tx, userRoles, u.UserId)
			if err != nil {
				return 0, err
			}
		}
	} else {
		err := model.DelAddUserRoles(tx, []model.UserRoleEntity{}, u.UserId)
		if err != nil {
			return 0, err
		}
	}

	return u.UserId, tx.Commit()
}

// 检查邮箱是否已使用
func CheckEmailUnique(userId int, email string) bool {
	user := model.Entity{Email: email, UserId: userId}
	return user.CheckEmailUnique()
}

//检查手机号是否已使用,存在返回true,否则false
func CheckPhoneUnique(userId int, phonenumber string) bool {
	user := model.Entity{PhoneNumber: phonenumber, UserId: userId}
	return user.CheckPhoneUnique()
}

//根据主键查询用户信息
func SelectRecordById(id int) model.Entity {
	return model.SelectRecordById(id)
}

// 查询未分配用户角色列表
func SelectUnAllocatedList(req *role.AllocatedReq) []model.Entity {
	return model.SelectUnAllocatedList(req)
}

// 查询已分配用户角色列表
func SelectAllocatedList(param *role.AllocatedReq) []model.Entity {
	return model.SelectAllocatedList(param)
}

func SelectRecordList(param *model.SelectPageReq) []model.Entity {
	return model.SelectPageList(param)
}

// 是否被锁
func IsLock(c *gin.Context) bool {
	user := GetProfile(c)
	if user == nil || user.UserId == 0 {
		return true
	}

	return user.Status != "0"
}

//判断是否是系统管理员
func IsAdmin(userId int) bool {
	if userId == 1 {
		return true
	} else {
		return false
	}
}

// 获得用户信息详情
func GetProfile(c *gin.Context) *model.ToSession {
	session := sessions.Default(c)
	tmp := session.Get(constant.USER_SESSION_MARK)
	s, ok := tmp.(string)
	if !ok {
		return nil
	}

	var u model.ToSession
	if err := json.Unmarshal([]byte(s), &u); err != nil {
		return nil
	}
	return &u
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func SignIn(c *gin.Context, loginName, password string) (string, int, error) {
	user := model.Entity{LoginName: loginName}
	if err := user.FindUser(); err != nil {
		return "", 0, err
	}

	token := user.LoginName + password + user.Salt
	token = gmd5.MustEncryptString(token)

	if strings.Compare(user.Password, token) != 0 {
		return "", user.UserId, errors.New("用户名密码错误")
	}

	var d model.ToSession
	d.UserId = user.UserId
	d.DeptId = user.DeptId
	d.LoginName = user.LoginName
	d.UserName = user.UserName
	d.Email = user.Email
	d.PhoneNumber = user.PhoneNumber
	d.Sex = user.Sex
	d.Avatar = user.Avatar
	d.Status = user.Status
	d.CreateTime = user.CreateTime

	// 更新登陆时间
	model.UpdateLoginTime(d.UserId, time.Now(), c.ClientIP())

	return SaveUserToSession(c, &d), user.UserId, nil
}

//保存用户信息到session
func SaveUserToSession(c *gin.Context, user *model.ToSession) string {
	session := sessions.Default(c)
	session.Set(constant.USER_ID, user.UserId)

	sessionId := session.ID()

	tmp, _ := json.Marshal(user)
	session.Set(constant.USER_SESSION_MARK, string(tmp))
	_ = session.Save()

	return sessionId
}

//清空用户菜单缓存
func ClearMenuCache(user *model.ToSession) {
	cache.Instance().Delete(constant.MENU_CACHE + strconv.Itoa(user.UserId))
}
