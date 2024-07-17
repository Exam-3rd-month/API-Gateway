package orderhandler

import (
	pb "gateway/genprotos/order_pb"

	"github.com/gin-gonic/gin"
)

// CreatePaymentHandler godoc
// @Summary Create a new payment
// @Description Creates a new payment for an order using credit card information
// @Tags payment
// @Accept json
// @Produce json
// @Param request body pb.CreatePaymentRequest true "Payment information"
// @Success 200 {object} pb.CreatePaymentResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payment [post]
func (h *Orderhandler) CreatePaymentHandler(c *gin.Context) {
    h.logger.Info("CreatePaymentHandler")
    var req pb.CreatePaymentRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.IndentedJSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    resp, err := h.order.CreatePayment(c, &req)
    if err != nil {
        c.IndentedJSON(500, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.IndentedJSON(200, resp)
}

