package repository

import (
	"gorm.io/gorm"
	"order-service/repository/order"
)

type Repository struct {
	RepoOrder order.IRepoOrder
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		RepoOrder: order.NewRepo(db),
	}
}
