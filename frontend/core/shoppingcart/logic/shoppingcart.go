package logic

import "xframe/frontend/core/shoppingcart/dao"

// 用户接口
type ShoppingCartService interface {
	GetAllShoppingCart() []*dao.Entity
	DeleteShoppingCartById(id int) error
	UpdateShoppingCart(id int) error
}

func New() ShoppingCartService {
	return &shoppingCartService{&dao.Entity{}}
}

type shoppingCartService struct {
	dao *dao.Entity
}

func (u *shoppingCartService) GetAllShoppingCart() []*dao.Entity {
	return u.dao.SelectAll()
}

func (u *shoppingCartService) DeleteShoppingCartById(id int) error {
	return u.dao.DeleteByKey(id)
}

func (u *shoppingCartService) InsertShoppingCart() (int, error) {
	u.dao.ShoppingCartId = 1
	return u.dao.Insert()
}

func (u *shoppingCartService) UpdateShoppingCart(id int) error {
	return u.dao.UpdateByKey(id)
}
