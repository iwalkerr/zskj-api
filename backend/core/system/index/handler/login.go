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
	ValidateCode string `form:"validateCode" binding:"required,min=4,max=10"`
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

// 图形验证码
func CaptchaImage(c *gin.Context) {
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "操作成功",
		"data":  base64stringC,
		"idkey": idKeyC,
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
	verifyResult := base64Captcha.VerifyCaptcha(req.IdKey, req.ValidateCode)
	if !verifyResult {
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
