package logic

import (
	"xframe/frontend/core/user/dao"
)

// 用户接口
type iUserService interface {
	GetAllUser() []*dao.Entity
	DeleteUserById(id int) error
	UpdateUser(id int) error
}

func New() iUserService {
	return &userService{dao.Entity{}}
}

type userService struct {
	dao dao.Entity
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
