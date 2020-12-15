package rsa

import (
	"encoding/base64"
	"xframe/frontend/config"
	"xframe/pkg/utils/gorsa"

	"github.com/gin-gonic/gin"
)

// 默认post，参数data
func Decrypy(c *gin.Context) ([]byte, error) {
	bytes, err := DecrypyParam(c.PostForm("data"))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// 加密数据
func RsaEncrypt(data []byte) (string, error) {
	cipherText, err := gorsa.RsaEncryptBlock(data, config.PublicKeyPath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// 解密请求参数
func DecrypyParam(param string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		return nil, err
	}
	plainText, err := gorsa.RsaDecryptBlock(decodeBytes, config.PrivateKeyPath)
	return plainText, err
}
