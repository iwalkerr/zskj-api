package logic

import "xframe/frontend/core/shoppingcart/dao"

// 用户接口
type iShoppingCartService interface {
	GetAllShoppingCart() []*dao.Entity
	DeleteShoppingCartById(id int) error
	UpdateShoppingCart(id int) error
}

func New() iShoppingCartService {
	return &ShoppingCartService{dao.Entity{}}
}

type ShoppingCartService struct {
	dao dao.Entity
}

func (u *ShoppingCartService) GetAllShoppingCart() []*dao.Entity {
	return u.dao.SelectAll()
}

func (u *ShoppingCartService) DeleteShoppingCartById(id int) error {
	return u.dao.DeleteByKey(id)
}

func (u *ShoppingCartService) InsertShoppingCart() (int, error) {
	u.dao.ShoppingCartId = 1
	return u.dao.Insert()
}

func (u *ShoppingCartService) UpdateShoppingCart(id int) error {
	return u.dao.UpdateByKey(id)
}
