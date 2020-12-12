package logic

import (
	"errors"
	"fmt"
	"strings"
	"xframe/frontend/config"
	"xframe/frontend/core/user/dao"
	"xframe/pkg/token"
	"xframe/pkg/utils/gmd5"
)

// 用户接口
type UserService interface {
	// 用户登陆验证
	LoginUser(param *dao.Login) (string, error)
	// 用户注册
	GetAllUser() []*dao.Entity
	DeleteUserById(id int) error
	UpdateUser(id int) error
}

func New() UserService {
	return &userService{&dao.Entity{}}
}

type userService struct {
	dao *dao.Entity
}

// 用户登陆验证
func (u *userService) LoginUser(param *dao.Login) (string, error) {
	// 1.根据用户名查询
	userId, pwd := u.dao.GetPwdByUsername(param.Username)

	// 2.加密规则
	enPwd := param.Username + param.Password + config.UserSalt
	enPwd = gmd5.MustEncryptString(enPwd)

	if strings.Compare(pwd, enPwd) != 0 {
		return "", errors.New("用户名密码错误")
	}

	// 3.加密数据
	jwtString, err := token.New(userId, config.OutTime, config.RefreshTime).CreateToken(config.JwtEncryptKey)
	if err != nil {
		return "", errors.New("token生成错误")
	}

	fmt.Println(jwtString)

	return jwtString, nil
}

func (u *userService) GetAllUser() []*dao.Entity {
	return u.dao.SelectAll()
}

func (u *userService) DeleteUserById(id int) error {
	return u.dao.DeleteByKey(id)
}

func (u *userService) InsertUser() (int, error) {
	u.dao.UserId = 1
	return u.dao.Insert()
}

func (u *userService) UpdateUser(id int) error {
	return u.dao.UpdateByKey(id)
}
