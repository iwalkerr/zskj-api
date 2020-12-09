package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	// 验证用户有效性
	fmt.Println("进入验证拦截")

	// dd, err := aes.AESEncrypt([]byte("1kkkkkk"), []byte("123456781234567812345678"))
	// if err != nil {
	// 	c.Abort()
	// 	return
	// }

	// _, err = aes.AESDecrypt(dd, []byte("123456781234567812345678"))
	// if err != nil {
	// 	c.Abort()
	// 	return
	// }

	c.Next()
}
