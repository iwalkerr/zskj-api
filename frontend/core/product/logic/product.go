package logic

import "xframe/frontend/core/product/dao"

// 用户接口
type ProductService interface {
	GetAllProduct() []*dao.Entity
	DeleteProductById(id int) error
	UpdateProduct(id int) error
}

func New() ProductService {
	return &productService{&dao.Entity{}}
}

type productService struct {
	dao *dao.Entity
}

func (u *productService) GetAllProduct() []*dao.Entity {
	return u.dao.SelectAll()
}

func (u *productService) DeleteProductById(id int) error {
	return u.dao.DeleteByKey(id)
}

func (u *productService) InsertProduct() (int, error) {
	u.dao.ProductId = 1
	return u.dao.Insert()
}

func (u *productService) UpdateProduct(id int) error {
	return u.dao.UpdateByKey(id)
}
