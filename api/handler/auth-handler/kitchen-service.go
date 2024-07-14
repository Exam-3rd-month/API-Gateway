package authhandler

import (
    pb "gateway/genprotos/auth_pb"
    "github.com/gin-gonic/gin"
)

// CreateKitchenHandler creates a new kitchen in the system.
// @Summary Create a new kitchen
// @Description Create a new kitchen with provided details
// @Tags kitchen
// @Accept json
// @Produce json
// @Param request body auth_pb.CreateKitchenRequest true "Kitchen creation details"
// @Success 200 {object} auth_pb.CreateKitchenResponse
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /kitchen [post]
func (h *AuthHandler) CreateKitchenHandler(c *gin.Context) {
    h.logger.Info("CreateKitchenHandler")

    var req pb.CreateKitchenRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.IndentedJSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    resp, err := h.auth.CreateKitchen(c, &req)
    if err != nil {
        c.IndentedJSON(500, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.IndentedJSON(200, resp)
}

// UpdateKitchenHandler updates an existing kitchen.
// @Summary Update a kitchen
// @Description Update an existing kitchen with provided details
// @Tags kitchen
// @Accept json
// @Produce json
// @Param kitchen_id path string true "Kitchen ID"
// @Param request body auth_pb.UpdateKitchenRequest true "Kitchen update details"
// @Success 200 {object} auth_pb.UpdateKitchenResponse
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /kitchen/{kitchen_id} [put]
func (h *AuthHandler) UpdateKitchenHandler(c *gin.Context) {
    h.logger.Info("UpdateKitchenHandler")
    kitchen_id := c.Param("kitchen_id")

    var req = pb.UpdateKitchenRequest{
        KitchenId: kitchen_id,
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.IndentedJSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    resp, err := h.auth.UpdateKitchen(c, &req)
    if err != nil {
        c.IndentedJSON(500, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.IndentedJSON(200, resp)
}

// GetKitchenHandler retrieves details of a specific kitchen.
// @Summary Get kitchen details
// @Description Get details of a specific kitchen based on Kitchen ID
// @Tags kitchen
// @Accept json
// @Produce json
// @Param kitchen_id path string true "Kitchen ID"
// @Success 200 {object} auth_pb.GetKitchenResponse
// @Failure 500 {object} gin.H "Internal server error"
// @Router /kitchen/{kitchen_id} [get]
func (h *AuthHandler) GetKitchenHandler(c *gin.Context) {
    h.logger.Info("GetKitchenHandler")
    kitchen_id := c.Param("kitchen_id")

    var req = pb.GetKitchenRequest{
        KitchenId: kitchen_id,
    }

    resp, err := h.auth.GetKitchen(c, &req)
    if err != nil {
        c.IndentedJSON(500, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.IndentedJSON(200, resp)
}
