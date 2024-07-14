package orderhandler

import (
	"gateway/genprotos/order_pb"
	"log/slog"
)

type Orderhandler struct {
	logger *slog.Logger
	order  order_pb.OrderServiceClient
}

func NewOrderHandler(logger *slog.Logger, order order_pb.OrderServiceClient) *Orderhandler {
	return &Orderhandler{
		logger: logger,
		order:  order,
	}
}
