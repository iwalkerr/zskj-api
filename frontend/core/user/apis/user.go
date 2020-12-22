package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"xframe/frontend/common/response"
	"xframe/frontend/core/user/dao"
	userLogic "xframe/frontend/core/user/logic"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// @Tags 用户模块
// @Title 用户登陆
// @Summary APP登陆接口
// @Description 用户通过手机app登陆
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username formData string true "登陆用户名或手机号" default(zhangsan)
// @Param password formData string true "密码" default(12345678)
// @Success 200 {object} response.CommonRes
// @Router /api/v1/user/login [post]
func Login(c *gin.Context) {
	var req dao.Login
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c).Msg("参数错误").JSON()
		return
	}

	if user, encrypt, err := userLogic.New().LoginUser(&req); err != nil {
		response.Error(c).Msg(err.Error()).JSON()
	} else {
		response.Success(c).Data(gin.H{"user": user, "token": encrypt}).JSON()
	}
}

// @Tags 用户模块
// @Title 用户注册
// @Summary APP注册接口
// @Description 用户通过手机app注册
// @Accept x-www-form-urlencoded
// @Produce json
// @Param phone formData string true "登陆手机号码" default(13881887710)
// @Param password formData string true "密码" default(12345678)
// @Param authCode formData string true "手机收取的验证码" default(123456)
// @Success 200 {object} response.CommonRes
// @Router /api/v1/user/register [post]
func Register(c *gin.Context) {
	var req dao.Register
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c).Msg("参数错误").JSON()
		return
	}

	// TODD: 验证验证码
	if req.AuthCode != "123456" {
		response.Error(c).Msg("验证码错误").JSON()
		return
	}
	// 验证手机号码是否存在
	if ok := userLogic.New().ExistPhone(req.Phone); ok {
		response.Error(c).Msg("手机号码已存在").JSON()
		return
	}
	// 保存手机号码
	id, err := userLogic.New().SaveUser(req.Phone, req.Password)
	if err != nil {
		response.Error(c).Msg("用户注册失败").JSON()
		return
	}

	log.Println("注册成功 ===> ", id)
	response.Success(c).JSON()
}

// @Tags 用户模块
// @Title 获取验证码
// @Summary APP获取验证码
// @Description 用户通过手机app获取验证码
// @Accept x-www-form-urlencoded
// @Produce json
// @Param phone formData string true "登陆手机号码" default(13881887710)
// @Success 200 {object} response.CommonRes
// @Router /api/v1/user/getcode [post]
func GetCode(c *gin.Context) {
	var req dao.PhoneReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c).Msg("手机格式错误").JSON()
		return
	}

	// 验证手机的合法性

	userLogic.New().SendPhoneNote(req.Phone)
	response.Success(c).JSON()
}

type QueryParam struct {
	Username string `json:"username"`
	Status   int    `json:"status"  validate:"max=20,min=6"`
	Keywords string `json:"keywords"`
}

// 获取用户头像
func HeadPic(c *gin.Context) {
	var d QueryParam

	_ = json.Unmarshal(c.MustGet("b").([]byte), &d)
	if err := validator.New().Struct(d); err != nil {
		response.Error(c).Msg("参数错误").JSON()
		return
	}

	fmt.Printf("%+v\n", d)

	// uid := c.GetString("uid")
	// fmt.Println(uid, nt)

	response.Success(c).Data().JSON()
}
