package logic

import "xframe/frontend/core/seller/dao"

// 用户接口
type SellerService interface {
	GetAllSeller() []*dao.Entity
	DeleteSellerById(id int) error
	UpdateSeller(id int) error
}

func New() SellerService {
	return &sellerService{&dao.Entity{}}
}

type sellerService struct {
	dao *dao.Entity
}

func (u *sellerService) GetAllSeller() []*dao.Entity {
	return u.dao.SelectAll()
}

func (u *sellerService) DeleteSellerById(id int) error {
	return u.dao.DeleteByKey(id)
}

func (u *sellerService) InsertSeller() (int, error) {
	u.dao.SellerId = 1
	return u.dao.Insert()
}

func (u *sellerService) UpdateSeller(id int) error {
	return u.dao.UpdateByKey(id)
}
