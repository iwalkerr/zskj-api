package logic

import "xframe/frontend/core/orders/dao"

// 用户接口
type OrdersService interface {
	GetAllOrders() []*dao.Entity
	DeleteOrdersById(id int) error
	UpdateOrders(id int) error
}

func New() OrdersService {
	return &ordersService{&dao.Entity{}}
}

type ordersService struct {
	dao *dao.Entity
}

func (u *ordersService) GetAllOrders() []*dao.Entity {
	return u.dao.SelectAll()
}

func (u *ordersService) DeleteOrdersById(id int) error {
	return u.dao.DeleteByKey(id)
}

func (u *ordersService) InsertOrders() (int, error) {
	u.dao.OrderId = 1
	return u.dao.Insert()
}

func (u *ordersService) UpdateOrders(id int) error {
	return u.dao.UpdateByKey(id)
}
