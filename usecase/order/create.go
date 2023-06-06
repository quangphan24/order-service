package order

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"order-service/model"
	"order-service/payload"
	proto "order-service/proto/stock"
)

func (uc *OrderUseCase) CreateOrder(ctx context.Context, req payload.CreateOrderRequest) (*model.Order, error) {
	var (
		total int64
	)
	total = 0
	//check quantity of product in stock
	for _, item := range req.Items {
		product, err := uc.grpcClient.StockClient.GetProductById(ctx, &proto.Int{Value: int64(item.ProductId)})
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		if product.Quantity-item.Quantity < 0 {
			return nil, errors.New(fmt.Sprintf("The number of products %v is not enough", product.Name))
		}

		total = total + item.Price*item.Quantity
	}

	//create
	order := &model.Order{
		WalletId:   req.WalletID,
		Total:      total,
		Status:     "new",
		OrderItems: req.Items,
	}
	err := uc.repo.CreateOrder(order)

	return order, err
}
