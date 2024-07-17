package orderhandler

import (
	pb "gateway/genprotos/order_pb"

	"github.com/gin-gonic/gin"
)

// @Summary Add a new dish
// @Description Add a new dish to the menu
// @Tags dish
// @Accept json
// @Produce json
// @Param request body pb.AddDishRequest true "Dish details"
// @Success 200 {object} pb.AddDishResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dish [post]
func (h *Orderhandler) AddDishHandler(c *gin.Context) {
	h.logger.Info("AddDishHandler")

	var req pb.AddDishRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.AddDish(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Update a dish
// @Description Update an existing dish's details
// @Tags dish
// @Accept json
// @Produce json
// @Param dish_id path string true "Dish ID"
// @Param request body pb.UpdateDishRequest true "Updated dish details"
// @Success 200 {object} pb.UpdateDishResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dish/{dish_id} [put]
func (h *Orderhandler) UpdateDishHandler(c *gin.Context) {
	h.logger.Info("UpdateDishHandler")
	dish_id := c.Param("dish_id")

	var req = pb.UpdateDishRequest{
		DishId: dish_id,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.UpdateDish(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Delete a dish
// @Description Delete an existing dish
// @Tags dish
// @Accept json
// @Produce json
// @Param dish_id path string true "Dish ID"
// @Success 200 {object} pb.DeleteDishResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dish/{dish_id} [delete]
func (h *Orderhandler) DeleteDishHandler(c *gin.Context) {
	h.logger.Info("DeleteDishHandler")
	dish_id := c.Param("dish_id")

	var req = pb.DeleteDishRequest{
		DishId: dish_id,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.DeleteDish(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary List dishes
// @Description Get a list of dishes for a specific kitchen
// @Tags dish
// @Accept json
// @Produce json
// @Param kitchen_id path string true "Kitchen ID"
// @Success 200 {object} pb.ListDishesResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dish/kitchen/{kitchen_id} [get]
func (h *Orderhandler) ListDishesHandler(c *gin.Context) {
	h.logger.Info("ListDishesHandler")
	kitchen_id := c.Param("kitchen_id")

	var req = pb.ListDishesRequest{
		KitchenId: kitchen_id,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.ListDishes(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// UpdateDishNutritionInfoHandler godoc
// @Summary Update dish nutrition info
// @Description Update nutrition information for a specific dish
// @Tags dish
// @Accept json
// @Produce json
// @Param dish_id path string true "Dish ID"
// @Param request body pb.UpdateDishNutritionInfoRequest true "Nutrition info update request"
// @Success 200 {object} pb.UpdateDishNutritionInfoResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dish/{dish_id}/nutrition [put]
func (h *Orderhandler) UpdateDishNutritionInfoHandler(c *gin.Context) {
	h.logger.Info("UpdateDishNutritionInfoHandler")
	dish_id := c.Param("dish_id")

	var req = pb.UpdateDishNutritionInfoRequest{
		DishId: dish_id,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.UpdateDishNutritionInfo(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}