package test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"xframe/frontend/core/user/dao"
	"xframe/pkg/utils/gorsa"
)

func TestDeRsa(t *testing.T) {
	param := "a5T04Eo/wCCjDB8WsVhO7l8MNc41AX66ERydHLr+nwto3z1t08wtyzF3bFUVCifUPAeiduC/IPU4yQbedXjwOStDciO8mTQDcqS9i59zB1z0obBTZgDp/gwHNdSJjGa/0+Tkkw/EXuvuiHHPT4VbAPkkA+iXca+26XQIil7UvnI9U7Rz3jiZNG2QcD0KVTtpmXeXbZKrBiKw7Uzv2/CDYCEmwFyf9k0DWsr4goQYirvetz625l22J27bu1a7hHmEt3ugOi1BRADh81yJaZDBU3xw/5SFSt5r039H7vP+qOIZrc8w8o9yY7As5prmF8KCooS9mdJmd9+VU2hGVlyS1JK9tYY91PU8Eb+pkD9TTFIfKo4+Y0UsQeTFWdwBL9TQUBVDTdMkBy2g7kWN387DKE/PLlcuIggE4Ot1SU+NzBT8jS4W80oFFerBMKBGkHXe3F0dZxUeusm7TccGeAPyt9QB2pw+bzyxvMY9t+aq2C7TZCN9RpIx8Lcn3Fz9FnOmTd0O8f5XmznA7Mdt/JrzB5opfv6GsBDFI2VuOSv5uJVi5TdmsqhHUefLqQHidnw8V8Ju44pB/9BjNCOV39v7QZYrvuvMPmOaSArF6zaZVpF9oXQ2kY8jgIrMVtE0oQWriaeDPdZbeKTZcqH0yVDlJ9SEuW0OR4G3dI2UI8coHU0="
	decodeBytes, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		return
	}
	plainText, err := gorsa.RsaDecryptBlock(decodeBytes, "../frontend/res/keys/privateKey.pem")
	fmt.Println(err)

	var user UserData
	_ = json.Unmarshal(plainText, &user)

	fmt.Printf("%+v\n", user)

	// fmt.Println(string(plainText), err)
}

type QueryParam struct {
	Jwt      string `json:"token"`
	Username string `json:"username"`
	Status   int    `json:"status"`
	Keywords string `json:"keywords"`
}

type UserData struct {
	User  dao.Entity
	Nt    string `json:"nt"`
	Token string `json:"token"`
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
