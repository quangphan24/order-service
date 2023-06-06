package order

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"order-service/payload"
	"order-service/util"
)

func (u *OrderUseCase) UpdateOrder(ctx context.Context, req payload.UpdateOrderRequest) error {
	//get one
	order, err := u.repo.GetOneOrder(req.Id)
	if err != nil {
		return err
	}
	if order.Status == req.Status {
		return errors.New("Status not change")
	}
	order.Status = req.Status
	err = u.repo.UpdateOrder(order)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if req.Status == "completed" {
		// go updateWallet
		go func() {
			data := struct {
				Amount   int64  `json:"amount"`
				WalletId string `json:"wallet_id"`
			}{
				Amount:   order.Total,
				WalletId: order.WalletId,
			}
			err := u.publisher.Publish(util.UPDATE_WALLET_AMOUNT, data)
			if err != nil {
				logrus.Error(err)
				return
			}
		}()
		// update quantity of product
		go func() {
			type UpdateQuantity struct {
				Id       int64 `json:"id"`
				Quantity int64 `json:"quantity"`
			}
			for _, item := range order.OrderItems {
				err := u.publisher.Publish(util.UPDATE_QUANTITY_PRODUCT, UpdateQuantity{Id: item.ProductId, Quantity: item.Quantity})
				if err != nil {
					logrus.Error(err)
					return
				}
			}
		}()
	}
	return nil
}
