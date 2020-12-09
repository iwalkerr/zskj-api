package logic

import "xframe/frontend/core/product/dao"

// 用户接口
type iProductService interface {
	GetAllProduct() []*dao.Entity
	DeleteProductById(id int) error
	UpdateProduct(id int) error
}

func New() iProductService {
	return &ProductService{dao.Entity{}}
}

type ProductService struct {
	dao dao.Entity
}

func (u *ProductService) GetAllProduct() []*dao.Entity {
	return u.dao.SelectAll()
}

func (u *ProductService) DeleteProductById(id int) error {
	return u.dao.DeleteByKey(id)
}

func (u *ProductService) InsertProduct() (int, error) {
	u.dao.ProductId = 1
	return u.dao.Insert()
}

func (u *ProductService) UpdateProduct(id int) error {
	return u.dao.UpdateByKey(id)
}
