package token

import (
	"encoding/base64"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type YClaims struct {
	//【JWT ID】     该jwt的唯一ID编号
	//【issuer】     发布者的url地址
	//【issued at】  该jwt的发布时间；unix 时间戳
	//【subject】    该JWT所面向的用户，用于处理特定应用，不是常用的字段
	//【audience】   接受者的url地址
	//【expiration】 该jwt销毁的时间；unix时间戳
	//【not before】 该jwt的使用时间不能早于该时间；unix时间戳
	StandardClaims *jwt.StandardClaims
	RefreshTime    int64 //【The refresh time】 该jwt刷新的时间；unix时间戳
}

type Token struct {
	Claim    *YClaims
	Token    string
	NewToken string
}

//创建Claims
func New(id string, timeOut, refreshTime int) *YClaims {
	if timeOut <= 0 {
		timeOut = 3600
	}
	if refreshTime <= 0 {
		refreshTime = timeOut / 2
	}

	var claims YClaims
	standardClaims := new(jwt.StandardClaims)
	standardClaims.Id = id
	standardClaims.ExpiresAt = time.Now().Add(time.Second * time.Duration(timeOut)).Unix()
	standardClaims.IssuedAt = time.Now().Unix()

	claims.RefreshTime = time.Now().Add(time.Second * time.Duration(refreshTime)).Unix()
	claims.StandardClaims = standardClaims
	return &claims
}

func (c *YClaims) SetIss(issuer string) *YClaims {
	c.StandardClaims.Issuer = issuer
	return c
}

func (c *YClaims) SetSub(subject string) *YClaims {
	c.StandardClaims.Subject = subject
	return c
}

func (c *YClaims) SetAud(audience string) *YClaims {
	c.StandardClaims.Audience = audience
	return c
}

func (c *YClaims) SetNbf(notBefore int64) *YClaims {
	c.StandardClaims.NotBefore = notBefore
	return c
}

func (c *YClaims) Valid() error {
	//标准验证
	return c.StandardClaims.Valid()
}

//创建token
func (c *YClaims) CreateToken(encryptKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	mySignKeyBytes, err := base64.URLEncoding.DecodeString(encryptKey)
	if err != nil {
		return "", err
	}
	return token.SignedString(mySignKeyBytes)
}

//验证token
func VerifyAuthToken(token, encryptKey string, timeOut, refreshTime int) (*Token, error) {
	mySignKeyBytes, err := base64.URLEncoding.DecodeString(encryptKey) //需要用和加密时同样的方式转化成对应的字节数组
	if err != nil {
		return nil, err
	}
	var yjgClaims YClaims

	parseAuth, err := jwt.ParseWithClaims(token, &yjgClaims, func(*jwt.Token) (interface{}, error) {
		return mySignKeyBytes, nil
	})
	if err != nil {
		return nil, err
	}

	//验证claims
	if err := parseAuth.Claims.Valid(); err != nil {
		return nil, err
	}

	rs := new(Token)
	rs.Claim = &yjgClaims
	rs.Token = token
	//判断是否需要刷新
	if time.Now().Unix() > yjgClaims.RefreshTime {
		//生成新token
		newToken, err := New(yjgClaims.StandardClaims.Id, timeOut, refreshTime).CreateToken(encryptKey)
		if err == nil {
			rs.NewToken = newToken
		}
	}
	return rs, nil
}
