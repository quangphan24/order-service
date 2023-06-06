package grpc_client

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"order-service/conf"
	proto "order-service/proto/authen"
	proto2 "order-service/proto/stock"
	"order-service/proto/user"
)

type GrpcClient struct {
	AutheClient proto.AuthenServiceClient
	StockClient proto2.StockServiceClient
	UserClient  user.UserServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	authConn, err := grpc.Dial(conf.GetConfig().GRPCServer.AuthenServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Error("did not connect authen service: %v", err)
		return nil, err
	}

	//user service
	userConn, err := grpc.Dial(conf.GetConfig().GRPCServer.UserServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Error("did not connect order service: %v", err)
		return nil, err
	}
	//stock service
	stockConn, err := grpc.Dial(conf.GetConfig().GRPCServer.StockServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Error("did not connect stock service: %v", err)
		return nil, err
	}

	return &GrpcClient{
		AutheClient: proto.NewAuthenServiceClient(authConn),
		StockClient: proto2.NewStockServiceClient(stockConn),
		UserClient:  user.NewUserServiceClient(userConn),
	}, nil
}
