package logic

import "xframe/frontend/core/orders/dao"

// 用户接口
type iOrdersService interface {
	GetAllOrders() []*dao.Entity
	DeleteOrdersById(id int) error
	UpdateOrders(id int) error
}

func New() iOrdersService {
	return &OrdersService{dao.Entity{}}
}

type OrdersService struct {
	dao dao.Entity
}

func (u *OrdersService) GetAllOrders() []*dao.Entity {
	return u.dao.SelectAll()
}

func (u *OrdersService) DeleteOrdersById(id int) error {
	return u.dao.DeleteByKey(id)
}

func (u *OrdersService) InsertOrders() (int, error) {
	u.dao.OrderId = 1
	return u.dao.Insert()
}

func (u *OrdersService) UpdateOrders(id int) error {
	return u.dao.UpdateByKey(id)
}
