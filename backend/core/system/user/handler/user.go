package handler

import (
	"os"
	"strconv"
	"time"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	deptService "xframe/backend/core/system/dept/service"
	post "xframe/backend/core/system/post/model"
	postService "xframe/backend/core/system/post/service"
	role "xframe/backend/core/system/role/model"
	roleService "xframe/backend/core/system/role/service"
	"xframe/backend/core/system/user/model"
	"xframe/backend/core/system/user/service"

	"github.com/gin-gonic/gin"
)

// 检查密码
func CheckPassword(c *gin.Context) {
	var req model.CheckPasswordReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("用户管理", req).Write()
		return
	}
	user := service.GetProfile(c)
	var d model.Entity
	d.LoginName = user.LoginName
	if flag := service.CheckPassword(&d, req.Password); flag {
		_, _ = c.Writer.WriteString("true")
	} else {
		_, _ = c.Writer.WriteString("false")
	}
}

//修改用户密码
func UpdatePassword(c *gin.Context) {
	var req model.PasswordReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("用户管理", req).Write()
	}
	if err := service.UpdatePassword(c, &req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("修改用户密码失败").Log("用户管理", req).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Edit).Log("修改用户密码", req).Write()
	}
}

//修改密码页面
func EditPwd(c *gin.Context) {
	user := service.GetProfile(c)
	resp.BuildTpl(c, "core/system/user/profile/resetPwd").Write(gin.H{"user": user})
}

//修改用户信息
func Update(c *gin.Context) {
	var req model.ProfileReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("缺少参数").Log("用户管理", req).Write()
		return
	}
	if err := service.UpdateProfile(c, &req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("缺少参数").Log("用户管理", req).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Edit).Log("用户管理", req).Write()
	}
}

func Avatar(c *gin.Context) {
	user := service.GetProfile(c)
	resp.BuildTpl(c, "core/system/user/profile/avatar").Write(gin.H{"user": user})
}

func UpdateAvatar(c *gin.Context) {
	user := service.GetProfile(c)
	curDir, err := os.Getwd()
	if err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("缺少参数").Log("保存头像", gin.H{"userid": user.UserId}).Write()
		return
	}
	saveDir := curDir + "/public/upload/"
	fileHead, err := c.FormFile("avatarfile")
	if err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("没有获取到上传文件").Log("保存头像", gin.H{"userid": user.UserId}).Write()
		return
	}
	curdate := time.Now().UnixNano()
	filename := user.LoginName + strconv.FormatInt(curdate, 10) + ".png"
	dts := saveDir + filename
	if err := c.SaveUploadedFile(fileHead, dts); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("保存头像", gin.H{"userid": user.UserId}).Write()
		return
	}
	avatar := "/upload/" + filename
	// 更新头像
	if err = service.UpdateAvatar(c, avatar); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("保存头像", gin.H{"userid": user.UserId}).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Edit).Log("保存头像", gin.H{"userid": user.UserId}).Write()
	}
}

func CheckEmailUnique(c *gin.Context) {
	var req model.CheckEmailReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	if res := service.CheckEmailUnique(req.UserId, req.Email); res {
		_, _ = c.Writer.WriteString("1")
	} else {
		_, _ = c.Writer.WriteString("0")
	}
}

//检查手机号是否存在
func CheckPhoneUnique(c *gin.Context) {
	var req *model.CheckPhoneReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	if res := service.CheckPhoneUnique(req.UserId, req.Phonenumber); res {
		_, _ = c.Writer.WriteString("1")
	} else {
		_, _ = c.Writer.WriteString("0")
	}
}

//检查手机号是否存在
func CheckPhoneUniqueAll(c *gin.Context) {
	var req model.CheckPhoneAllReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	if isHasPhone := service.CheckPhoneUniqueAll(req.Phonenumber); isHasPhone {
		_, _ = c.Writer.WriteString("1")
	} else {
		_, _ = c.Writer.WriteString("0")
	}
}

//检查邮箱是否存在
func CheckEmailUniqueAll(c *gin.Context) {
	var req model.CheckEmailAllReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	if result := service.CheckEmailUniqueAll(req.Email); result {
		_, _ = c.Writer.WriteString("1")
	} else {
		_, _ = c.Writer.WriteString("0")
	}
}

//检查登陆名是否存在
func CheckLoginNameUnique(c *gin.Context) {
	var req model.CheckLoginNameReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	if result := service.CheckLoginName(req.LoginName); result {
		_, _ = c.Writer.WriteString("1")
	} else {
		_, _ = c.Writer.WriteString("0")
	}
}

func Profile(c *gin.Context) {
	user := service.GetProfile(c)
	// 获取部门名称
	deptName := deptService.GetDeptName(user.DeptId)

	var d model.Entity
	d.UserId = user.UserId
	d.DeptId = user.DeptId
	d.LoginName = user.LoginName
	d.UserName = user.UserName
	d.Email = user.Email
	d.PhoneNumber = user.PhoneNumber
	d.Sex = user.Sex
	d.Avatar = user.Avatar
	d.Status = user.Status
	d.DeptName = deptName
	d.CreateTime = user.CreateTime

	resp.BuildTpl(c, "core/system/user/profile/profile").Write(gin.H{"user": d})
}

func List(c *gin.Context) {
	resp.BuildTpl(c, "core/system/user/list").Write()
}

func ListAjax(c *gin.Context) {
	var req model.SelectPageReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("用户管理", req).Write()
		return
	}
	rows := service.SelectRecordList(&req)
	resp.BuildTable(c, req.PageReq, rows).Write()
}

func Add(c *gin.Context) {
	var paramsRole role.SelectPageReq
	var paramsPost post.SelectPageReq
	// 获取所有角色
	roleList := roleService.SelectRecordAll(&paramsRole)
	// 所有岗位
	postList := postService.SelectListAll(&paramsPost)

	user := model.Entity{Status: "0"}
	resp.BuildTpl(c, "core/system/user/edit").Write(gin.H{
		"roles": roleList,
		"posts": postList,
		"user":  user,
		"msg":   "add",
	})
}

func AddSave(c *gin.Context) {
	var req model.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Add).Msg(err.Error()).Log("新增用户", req).Write()
		return
	}
	//判断登陆名是否已注册
	if isHasName := service.CheckLoginName(req.LoginName); isHasName {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("登陆名已经存在").Log("新增用户", req).Write()
		return
	}
	//判断手机号码是否已注册
	if isHadPhone := service.CheckPhoneUniqueAll(req.Phonenumber); isHadPhone {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("手机号码已经存在").Log("新增用户", req).Write()
		return
	}
	//判断邮箱是否已注册
	if isHadEmail := service.CheckEmailUniqueAll(req.Email); isHadEmail {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("邮箱已经存在").Log("新增用户", req).Write()
		return
	}
	if uid, err := service.AddSave(c, &req); err == nil && uid > 0 {
		resp.Success(c).Data(uid).Btype(constant.Buniss_Add).Log("新增用户", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Add).Log("新增用户", req).Write()
	}
}

func Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}

	user := service.SelectRecordById(id)
	if user.UserId <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "用户不存在"})
		return
	}
	var deptName string
	// 获取部门信息
	if user.DeptId > 0 {
		deptName = deptService.GetDeptName(user.DeptId)
	}
	roles := roleService.SelectRoleContactVo(id)
	posts := postService.SelectPostsByUserId(id)
	resp.BuildTpl(c, "core/system/user/edit").Write(gin.H{
		"user":     user,
		"deptName": deptName,
		"roles":    roles,
		"posts":    posts,
		"msg":      "edit",
	})
}

func EditSave(c *gin.Context) {
	var req model.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("修改用户", req).Write()
		return
	}
	//判断手机号码是否已注册
	if isHadPhone := service.CheckPhoneUnique(req.UserId, req.Phonenumber); isHadPhone {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("手机号码已经存在").Log("修改用户", req).Write()
		return
	}
	//判断邮箱是否已注册
	if isHadEmail := service.CheckEmailUnique(req.UserId, req.Email); isHadEmail {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("邮箱已经存在").Log("修改用户", req).Write()
		return
	}
	if uid, err := service.EditSave(c, &req); err == nil && uid > 0 {
		resp.Success(c).Data(uid).Btype(constant.Buniss_Edit).Log("修改用户", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Edit).Log("修改用户", req).Write()
	}
}

func Remove(c *gin.Context) {
	var req constant.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Del).Msg(err.Error()).Log("删除用户", req).Write()
	}
	if err := service.DeleteRecordByIds(req.Ids); err == nil {
		resp.Success(c).Data(1).Btype(constant.Buniss_Del).Log("删除用户", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Del).Log("删除用户", req).Write()
	}
}

func Export(c *gin.Context) {
	var req model.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("导出Excel", req).Write()
	}
	if url, err := service.Export(&req); err == nil {
		resp.Success(c).Msg(url).Log("导出Excel", req).Write()
	} else {
		resp.Error(c).Msg("导出Excel失败").Log("导出Excel", req).Write()
	}
}

func ResetPwd(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("userId"))
	if id <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}
	if user := service.SelectRecordById(id); user.UserId > 0 {
		resp.BuildTpl(c, "core/system/user/resetPwd").Write(gin.H{"user": user})
	} else {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "用户不存在"})
	}
}

func ResetPwdSave(c *gin.Context) {
	var req model.ResetPwdReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("重置密码", req).Write()
		return
	}
	if err := service.ResetPassword(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Log("重置密码", req).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Edit).Log("重置密码", req).Write()
	}
}
