package logic

import "xframe/frontend/core/seller/dao"

// 用户接口
type iSellerService interface {
	GetAllSeller() []*dao.Entity
	DeleteSellerById(id int) error
	UpdateSeller(id int) error
}

func New() iSellerService {
	return &SellerService{dao.Entity{}}
}

type SellerService struct {
	dao dao.Entity
}

func (u *SellerService) GetAllSeller() []*dao.Entity {
	return u.dao.SelectAll()
}

func (u *SellerService) DeleteSellerById(id int) error {
	return u.dao.DeleteByKey(id)
}

func (u *SellerService) InsertSeller() (int, error) {
	u.dao.SellerId = 1
	return u.dao.Insert()
}

func (u *SellerService) UpdateSeller(id int) error {
	return u.dao.UpdateByKey(id)
}
