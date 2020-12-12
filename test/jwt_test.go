package test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"xframe/pkg/utils/gorsa"
)

func TestDeRsa(t *testing.T) {
	param := "EeDRhQQFHW1xkgQR3+T5x5s2SQMCoD2tlPDKFKuuXSKh8UPE29kwv+gRG4Tt14FNpv7j6wBPK1CHxFNjXPF4lvyJTHVz6cUUqfWNpEEEGvCnU12Dm31KumGLwVAWK9vQpX16pPlHXF5Tr/W4rQM3OUVr74gx+iDgn6GSTMeCzihs3jh9L5rSc/xECwiA50MET0yV2BIn+zFZO38jkHhSE0RXAmmWPr0ibchb/lA4KJvEZzrpbrtaxYpP/IVGKT0Pj5eCxvT0v/vrM0FUf3BxnM5yuFo1nfvh9t61078ciKVfFLzGO9y6MOOilVwU2TDH6RxDGAgqxbyUj8GQi22Otw=="
	decodeBytes, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		return
	}
	plainText, err := gorsa.RsaDecryptBlock(decodeBytes, "../frontend/res/keys/privateKey.pem")
	fmt.Println(string(plainText), err)
}

type QueryParam struct {
	Jwt      string `json:"token"`
	Username string `json:"username"`
	Status   int    `json:"status"`
	Keywords string `json:"keywords"`
}

func TestEnRsa(t *testing.T) {
	data := QueryParam{
		Jwt:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGFuZGFyZENsYWltcyI6eyJleHAiOjE2MTAzNjQ4NTYsImp0aSI6IjEiLCJpYXQiOjE2MDc3NzI4NTZ9LCJSZWZyZXNoVGltZSI6MTYwNzc3Mjg1N30.z9iAs5ED4e75lrMMOoiuOm0_Q5iAS3s-lX0ShDkFh5s",
		Username: "cco_username123",
		Status:   0,
		Keywords: "华为手机",
	}

	bytes, _ := json.Marshal(data)
	cipherText, err := gorsa.RsaEncryptBlock(bytes, "../frontend/res/keys/publicKey.pem")
	if err != nil {
		log.Println("获取公钥失败,err=", err.Error())
		return
	}

	text := base64.StdEncoding.EncodeToString(cipherText)

	fmt.Println(text)
}
