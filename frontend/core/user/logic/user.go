package logic

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"xframe/frontend/config"
	"xframe/frontend/core/user/dao"
	"xframe/pkg/token"
	"xframe/pkg/utils/gmd5"
	"xframe/pkg/utils/random"
)

// 用户接口
type UserService interface {
	LoginUser(param *dao.Login) (*dao.LoginResp, string, error) // 用户登陆验证
	SendPhoneNote(phone string)                                 // 发送验证码
	ExistPhone(phone string) bool                               // 查询手机号码是否存在
	SaveUser(username, password string) (int, error)            // 保存用户注册信息
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

func (u *userService) SaveUser(username, password string) (int, error) {
	// 加密规则
	enPwd := username + password + config.UserSalt
	enPwd = gmd5.MustEncryptString(enPwd)
	// 用户显示名
	name := "哈哈哈哈"
	userName := "DHHE3343322"

	u.dao.Password = enPwd
	u.dao.Phone = username
	u.dao.HeadPicture = "/static/upload/userhead/default.jpg" // 默认头像，可以动态随机选择
	u.dao.Name = name
	u.dao.Username = userName
	u.dao.CreateTime = time.Now()
	u.dao.UpdateTime = time.Now()

	return u.dao.Insert()
}

// TODO: 发送短信
func (u *userService) SendPhoneNote(phone string) {
	code := random.GenValidateCode(6)
	fmt.Println("发送成功", code)
}

// 查看手机号码是否存在
func (u *userService) ExistPhone(phone string) bool {
	return u.dao.ExistPhone(phone)
}

// 用户登陆验证
func (u *userService) LoginUser(param *dao.Login) (*dao.LoginResp, string, error) {
	// 1.根据用户名查询
	user := u.dao.GetPwdByUsername(param.Username)

	// 2.加密规则
	enPwd := param.Username + param.Password + config.UserSalt
	enPwd = gmd5.MustEncryptString(enPwd)

	if strings.Compare(user.Password, enPwd) != 0 {
		return nil, "", errors.New("用户名密码错误")
	}

	// 3.加密数据
	jwtString, err := token.New(strconv.Itoa(user.UserId), config.OutTime, config.RefreshTime).CreateToken(config.JwtEncryptKey)
	if err != nil {
		return nil, "", errors.New("token生成错误")
	}

	return &user, jwtString, nil
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
