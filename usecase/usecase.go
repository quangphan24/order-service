package usecase

import (
	"order-service/repository"
	"order-service/service/grpc_client"
	"order-service/service/rabbitmq/publisher"
	"order-service/usecase/order"
)

type UseCase struct {
	OrderUseCase order.IOrderUseCase
}

func New(repo *repository.Repository, grpcClient *grpc_client.GrpcClient, pub *publisher.Publisher) *UseCase {
	return &UseCase{
		OrderUseCase: order.New(repo, grpcClient, pub),
	}
}
