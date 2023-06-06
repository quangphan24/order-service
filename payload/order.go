package payload

import "order-service/model"

type CreateOrderRequest struct {
	WalletID string            `json:"wallet_id"`
	Items    []model.OrderItem `json:"items"`
}

type UpdateOrderRequest struct {
	Status string `json:"status"`
	Id     string `json:"id" param:"id"`
}
