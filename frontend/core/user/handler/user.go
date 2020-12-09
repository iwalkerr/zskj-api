package handler

import (
	"fmt"
	"xframe/frontend/core/user/logic"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	list := logic.New().GetAllUser()
	fmt.Println(list)
}
