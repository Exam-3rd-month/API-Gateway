package handler

import (
	authhandler "gateway/api/handler/auth-handler"
	orderhandler "gateway/api/handler/order-handler"
	auth_pb "gateway/genprotos/auth_pb"
	order_pb "gateway/genprotos/order_pb"
	"gateway/internal/config"
	"log"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HandlerConfig struct {
	OrderHandler *orderhandler.Orderhandler
	AuthHandler  *authhandler.AuthHandler
}

func connect(port string) *grpc.ClientConn {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func NewHandlerConfig(logger *slog.Logger, config *config.Config) *HandlerConfig {
	return &HandlerConfig{
		OrderHandler: orderhandler.NewOrderHandler(logger, order_pb.NewOrderServiceClient(connect(config.Server.Order_Port))),
		AuthHandler:  authhandler.NewAuthHandler(logger, auth_pb.NewAuthServiceClient(connect(config.Server.Auth_Port))),
	}
}
