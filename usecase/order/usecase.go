package order

import (
	"context"
	"order-service/model"
	"order-service/payload"
	"order-service/repository"
	"order-service/repository/order"
	"order-service/service/grpc_client"
	"order-service/service/rabbitmq/publisher"
)

type OrderUseCase struct {
	repo       order.IRepoOrder
	grpcClient *grpc_client.GrpcClient
	publisher  *publisher.Publisher
}
type IOrderUseCase interface {
	CreateOrder(ctx context.Context, req payload.CreateOrderRequest) (*model.Order, error)
	UpdateOrder(ctx context.Context, req payload.UpdateOrderRequest) error
}

func New(repo *repository.Repository, grpcClient *grpc_client.GrpcClient, pub *publisher.Publisher) IOrderUseCase {
	return &OrderUseCase{
		repo:       repo.RepoOrder,
		grpcClient: grpcClient,
		publisher:  pub,
	}
}
