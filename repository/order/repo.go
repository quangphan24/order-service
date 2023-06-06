package order

import (
	"gorm.io/gorm"
	"order-service/model"
)

type RepoOrder struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IRepoOrder {
	return &RepoOrder{db: db}
}

//go:generate mockery --name IRepoUser
type IRepoOrder interface {
	CreateOrder(order *model.Order) error
	UpdateOrder(order *model.Order) error
	GetOneOrder(id string) (*model.Order, error)
}

func (r *RepoOrder) CreateOrder(order *model.Order) error {
	return r.db.Model(&model.Order{}).Create(&order).Error
}
func (r *RepoOrder) UpdateOrder(order *model.Order) error {
	return r.db.Model(&model.Order{}).Where("id = ?", order.ID).Updates(&order).Error
}
func (r *RepoOrder) GetOneOrder(id string) (*model.Order, error) {
	var res *model.Order
	err := r.db.Model(&model.Order{}).Preload("OrderItems").Where("id = ?", id).Take(&res).Error
	return res, err
}
