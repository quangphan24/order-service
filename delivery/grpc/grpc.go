package grpc

import (
	proto "order-service/proto/order"
	"order-service/usecase"
)

type ServerGRPC struct {
	proto.UnimplementedOrderServiceServer
	UseCase *usecase.UseCase
}
