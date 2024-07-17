package orderhandler

import (
	pb "gateway/genprotos/order_pb"

	"github.com/gin-gonic/gin"
)

// @Summary Add a new review
// @Description Add a new review for a kitchen or order
// @Tags review
// @Accept json
// @Produce json
// @Param request body pb.AddReviewRequest true "Review details"
// @Success 200 {object} pb.AddReviewResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /review [post]
func (h *Orderhandler) AddReviewHandler(c *gin.Context) {
	h.logger.Info("AddReviewHandler")

	var req pb.AddReviewRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.AddReview(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary List reviews for a kitchen
// @Description Get a list of reviews for a specific kitchen
// @Tags review
// @Accept json
// @Produce json
// @Param request body pb.ListReviewsRequest true "Additional request parameters"
// @Success 200 {object} pb.ListReviewsResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /review/kitchen/{kitchen_id} [post]
func (h *Orderhandler) ListReviewsHandler(c *gin.Context) {
	h.logger.Info("ListReviewsHandler")

	var req pb.ListReviewsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.ListReviews(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Create kitchen working hours
// @Description Create working hours for a specific kitchen
// @Tags kitchen
// @Accept json
// @Produce json
// @Param kitchen_id path string true "Kitchen ID"
// @Param request body pb.CreateKitchenWorkingHoursRequest true "Kitchen working hours details"
// @Success 201 {object} pb.CreateKitchenWorkingHoursResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /kitchen/{kitchen_id}/working-hours [post]
func (h *Orderhandler) CreateKitchenWorkingHoursHandler(c *gin.Context) {
	h.logger.Info("CreateKitchenWorkingHoursHandler")
	kitchen_id := c.Param("kitchen_id")

	var req = pb.CreateKitchenWorkingHoursRequest{
		KitchenId: kitchen_id,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.order.CreateKitchenWorkingHours(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(201, resp)
}

// GetUserActivityHandler godoc
// @Summary Get user activity
// @Description Retrieve activity statistics for a specific user
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} pb.GetUserActivityResponse
// @Failure 500 {object} gin.H
// @Router /user/{user_id}/activity [get]
func (h *Orderhandler) GetUserActivityHandler(c *gin.Context) {
	h.logger.Info("GetUserActivityHandler")
	user_id := c.Param("user_id")

	var req = pb.GetUserActivityRequest{
		UserId: user_id,
	}

	resp, err := h.order.GetUserActivity(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// // GetKitchenStatisticsHandler godoc
// // @Summary Get kitchen statistics
// // @Description Retrieve statistics for a specific kitchen
// // @Tags kitchen
// // @Accept json
// // @Produce json
// // @Param kitchen_id path string true "Kitchen ID"
// // @Success 200 {object} pb.GetKitchenStatisticsResponse
// // @Failure 500 {object} gin.H
// // @Router /kitchen/{kitchen_id}/statistics [get]
// func (h *Orderhandler) GetKitchenStatisticsHandler(c *gin.Context) {
// 	h.logger.Info("GetKitchenStatisticsHandler")
// 	kitchen_id := c.Param("kitchen_id")

// 	var req = pb.GetKitchenStatisticsRequest{
// 		KitchenId: kitchen_id,
// 	}

// 	resp, err := h.order.GetKitchenStatistics(c, &req)
// 	if err != nil {
// 		c.IndentedJSON(500, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.IndentedJSON(200, resp)
// }
