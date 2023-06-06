package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"order-service/conf"
	grpc2 "order-service/delivery/grpc"
	serviceHttp "order-service/delivery/http"
	"order-service/migrations"
	proto "order-service/proto/order"
	"order-service/repository"
	"order-service/service/grpc_client"
	"order-service/service/rabbitmq/consumer"
	"order-service/service/rabbitmq/publisher"
	"order-service/usecase"
)

const VERSION = "1.0.0"

// @title Example API
// @version 1.0

// @BasePath /api
// @schemes http http

// @securityDefinitions.apikey AuthToken
// @in header
// @name Authorization

// @description Transaction API.
func main() {
	conf.SetEnv()

	confMysql := conf.GetConfig().MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", confMysql.DBUser, confMysql.DBPass, confMysql.DBHost, confMysql.DBPort, confMysql.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()
	//migrations
	migrations.Up(db)

	repo := repository.New(db)
	grpcClient, err := grpc_client.NewGrpcClient()
	if err != nil {
		log.Fatal(err)
	}

	// connect rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//defer ch.Close()
	pub := publisher.NewPublisher(conn)

	uc := usecase.New(repo, grpcClient, pub)
	RunConsumer(conn, uc)

	go RunGRPC(uc)

	h := serviceHttp.NewHTTPHandler(uc)
	//go func() {
	//	h.Listener = httpL
	//	errs <- h.Start("")
	//}()
	if err := h.Start("0.0.0.0:8070"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
func RunConsumer(conn *amqp.Connection, uc *usecase.UseCase) {
	cons := consumer.NewConsumer(conn, uc)
	cons.StartConsumer()
}

func RunGRPC(uc *usecase.UseCase) {
	port := ":8071"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, grpc2.ServerGRPC{UseCase: uc})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
