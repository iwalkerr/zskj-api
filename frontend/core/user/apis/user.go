package apis

import (
	"encoding/json"
	"fmt"
	"xframe/frontend/common/response"
	"xframe/frontend/core/user/dao"
	userLogic "xframe/frontend/core/user/logic"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// 用户登陆
func Login(c *gin.Context) {
	var req dao.Login
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c).Msg("参数错误").JSON()
		return
	}

	if encrypt, err := userLogic.New().LoginUser(&req); err != nil {
		response.Error(c).Msg(err.Error()).JSON()
	} else {
		response.Success(c).Data(gin.H{"encrypt": encrypt}).JSON()
	}
}

// 注册
func Register(c *gin.Context) {

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
