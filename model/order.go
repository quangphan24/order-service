package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	BaseModel
	ID         string      `json:"id"`
	Total      int64       `json:"total"`
	WalletId   string      `json:"wallet_id"`
	Status     string      `json:"status"` // new | completed
	OrderItems []OrderItem `json:"order_items"`
}

func (Order) TableName() string {
	return "orders"
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return nil
}
