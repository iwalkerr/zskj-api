package handler

import (
	"fmt"
	userLogic "xframe/frontend/core/user/logic"
	"xframe/pkg/response"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	list := userLogic.New().GetAllUser()
	fmt.Println(list)

	response.Success(c).Msg("success").Data(list).JSON()
}
