package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"xframe/backend/common/cfg"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	loginlogService "xframe/backend/core/monitor/loginlog/service"
	"xframe/backend/core/system/index/service"
	userService "xframe/backend/core/system/user/service"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type RegisterReq struct {
	UserName     string `form:"username"  binding:"required,min=1,max=30"`
	Password     string `form:"password" binding:"required,min=5,max=150"`
	ValidateCode string `form:"validateCode" binding:"required,min=1,max=10"`
	IdKey        string `form:"idkey" binding:"required,min=5,max=30"`
}

// 去登陆页面
func Login(c *gin.Context) {
	if strings.EqualFold(c.Request.Header.Get("X-Requested-With"), "XMLHttpRequest") {
		resp.Error(c).Msg("未登录或登录超时。请重新登录").Write()
		return
	}

	resp.BuildTpl(c, "core/system/layout/login").Write(gin.H{
		"systemName":  cfg.Instance().Admin.BusinessName,
		"companyName": cfg.Instance().Admin.CompanyName,
	})
}

// 退出登陆
func Logout(c *gin.Context) {
	userService.SignOut(c)
	c.Redirect(302, "/login")
}

// 定义存储库
var store = base64Captcha.DefaultMemStore

// 图形验证码
func CaptchaImage(c *gin.Context) {
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	cc := base64Captcha.NewCaptcha(driver, store)
	// 获取
	id, b64s, err := cc.Generate()
	if err != nil {
		log.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "操作成功",
		"data":  b64s,
		"idkey": id,
	})
}

// 检查登陆
func CheckLogin(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg("请求参数错误").Write()
		return
	}

	// 比对验证码
	if ok := store.Verify(req.IdKey, req.ValidateCode, true); !ok {
		resp.Error(c).Code(constant.FAIL).Msg("验证码不正确").Write()
		return
	}

	// 检查用户是否被锁
	if isLock := service.CheckLock(req.UserName); isLock {
		resp.Error(c).Msg("账号已锁定，请30分钟后再试").Write()
		return
	}

	// 验证用户是否被锁
	if err := service.IsUserLock(req.UserName); err != nil {
		resp.Error(c).Msg(err.Error()).Write()
		log.Printf("登陆错误: %s,ip:%s ===> %+v", err.Error(), c.ClientIP(), req)
		return
	}

	sessionId, userId, err := userService.SignIn(c, req.UserName, req.Password)
	if err != nil {
		errTimes := loginlogService.SetPasswordCounts(req.UserName)
		having := 5 - errTimes

		// 写入日志
		loginlogService.AddUserLog(c, req.UserName, "账号或密码不正确", "1", "", 0, false)
		resp.Error(c).Msg("账号或密码不正确,还有" + strconv.Itoa(having) + "次之后账号将锁定").Write()
		return
	}

	// 登陆成功
	loginlogService.AddUserLog(c, req.UserName, "登陆成功", "0", sessionId, userId, true)

	resp.Success(c).Msg("登陆成功").Write()
}
