package auth

import (
	"encoding/json"
	"xframe/frontend/common/response"
	"xframe/frontend/common/rsa"
	"xframe/frontend/config"
	"xframe/pkg/token"

	"github.com/gin-gonic/gin"
)

type JwtData struct {
	Token string `form:"token"`
}

// 验证拦截
func Auth(c *gin.Context) {
	// rsa解密
	bytes, err := rsa.Decrypy(c)
	if err != nil {
		response.Error(c).Msg("解密错误").JSON()
		return
	}

	var jwt JwtData
	_ = json.Unmarshal(bytes, &jwt)

	if err := jWTAuth(c, jwt.Token); err != nil {
		response.Forbidden(c).Msg("无效的Token").JSON()
		return
	}

	c.Set("b", bytes)

	c.Next()
}

// 验证jwt
func jWTAuth(c *gin.Context, jwt string) error {
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	t, err := token.VerifyAuthToken(jwt, config.JwtEncryptKey, config.OutTime, config.RefreshTime)
	if err != nil {
		return err
	}
	// 将当前请求的uid信息保存到请求的上下文上
	c.Set("uid", t.Claim.StandardClaims.Id)
	//判断是否有新的token生成
	if t.NewToken != "" {
		c.Set("nt", t.NewToken)
	}
	return nil
}
