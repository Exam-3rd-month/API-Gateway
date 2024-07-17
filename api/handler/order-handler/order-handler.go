package orderhandler

import (
	pb "gateway/genprotos/order_pb"
	"log/slog"
)

type Orderhandler struct {
	logger *slog.Logger
	order  pb.OrderServiceClient
}

func NewOrderHandler(logger *slog.Logger, order pb.OrderServiceClient) *Orderhandler {
	return &Orderhandler{
		logger: logger,
		order:  order,
	}
}
