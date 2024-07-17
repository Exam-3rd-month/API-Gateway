package authhandler

import (
	pb "gateway/genprotos/auth_pb"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new kitchen
// @Description Create a new kitchen with the provided details
// @Tags kitchen
// @Accept json
// @Produce json
// @Param request body pb.CreateKitchenRequest true "Kitchen creation details"
// @Success 200 {object} pb.CreateKitchenResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
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

// @Summary Update a kitchen
// @Description Update an existing kitchen's details
// @Tags kitchen
// @Accept json
// @Produce json
// @Param kitchen_id path string true "Kitchen ID"
// @Param request body pb.UpdateKitchenRequest true "Updated kitchen details"
// @Success 200 {object} pb.UpdateKitchenResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
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

// @Summary Get kitchen details
// @Description Get details of a specific kitchen by ID
// @Tags kitchen
// @Accept json
// @Produce json
// @Param kitchen_id path string true "Kitchen ID"
// @Success 200 {object} pb.GetKitchenResponse
// @Failure 500 {object} gin.H
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

// @Summary List kitchens
// @Description Get a list of kitchens
// @Tags kitchen
// @Accept json
// @Produce json
// @Param request body pb.ListKitchensRequest true "List kitchens request parameters"
// @Success 200 {object} pb.ListKitchensResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /kitchen/list [post]
func (h *AuthHandler) ListKitchensHandler(c *gin.Context) {
	h.logger.Info("ListKitchensHandler")
	var req pb.ListKitchensRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.auth.ListKitchens(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Search kitchens
// @Description Search for kitchens based on provided criteria
// @Tags kitchen
// @Accept json
// @Produce json
// @Param request body pb.SearchKitchensRequest true "Search criteria"
// @Success 200 {object} pb.SearchKitchensResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /kitchen/search [post]
func (h *AuthHandler) SearchKitchensHandler(c *gin.Context) {
	h.logger.Info("SearchKitchensHandler")
	var req pb.SearchKitchensRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.auth.SearchKitchens(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}
