package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	BaseModel
	Id        string `json:"id"`
	OrderId   string `json:"order_id"`
	ProductId int64  `json:"product_id"`
	Quantity  int64  `json:"quantity"`
	Price     int64  `json:"price"`
	Total     int64  `json:"total"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
func (ot *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	ot.Id = uuid.New().String()
	return nil
}
