package orderhandler

import (
	pb "gateway/genprotos/order_pb"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new order
// @Description Create a new order with the provided details
// @Tags order
// @Accept json
// @Produce json
// @Param request body pb.CreateOrderRequest true "Order details"
// @Success 200 {object} pb.CreateOrderResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /order [post]
func (h *Orderhandler) CreateOrderHandler(c *gin.Context) {
	h.logger.Info("CreateOrderHandler")

	var req pb.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.CreateOrder(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Update order status
// @Description Update the status of an existing order
// @Tags order
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Param request body pb.UpdateOrderStatusRequest true "Updated order status"
// @Success 200 {object} pb.UpdateOrderStatusResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /order/{order_id}/status [put]
func (h *Orderhandler) UpdateOrderStatusHandler(c *gin.Context) {
	h.logger.Info("UpdateOrderStatusHandler")
	order_id := c.Param("order_id")

	var req = pb.UpdateOrderStatusRequest{
		OrderId: order_id,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.UpdateOrderStatus(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary List orders for a user
// @Description Get a list of orders for a specific user
// @Tags order
// @Accept json
// @Produce json
// @Param request body pb.ListOfOrdersRequest true "Additional request parameters"
// @Success 200 {object} pb.ListOfOrdersResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /order/user/{user_id} [get]
func (h *Orderhandler) ListOfOrdersHandler(c *gin.Context) {
	h.logger.Info("ListOfOrdersHandler")

	var req pb.ListOfOrdersRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.ListOfOrders(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Get orders for a kitchen
// @Description Get orders for a specific kitchen
// @Tags order
// @Accept json
// @Produce json
// @Param request body pb.GetOrderByKitchenIdRequest true "GetOrderByKitchenIdRequest"
// @Success 200 {object} pb.GetOrderByKitchenIdResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /order/kitchen/{kitchen_id} [get]
func (h *Orderhandler) GetOrderByKitchenIdHandler(c *gin.Context) {
	h.logger.Info("GetOrderByKitchenIdHandler")

	var req pb.GetOrderByKitchenIdRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.GetOrderByKitchenId(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}
