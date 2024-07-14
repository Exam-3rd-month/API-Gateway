package authhandler

import (
    pb "gateway/genprotos/auth_pb"
    "github.com/gin-gonic/gin"
)

// @Summary Register a new user
// @Description Register a new user in the system
// @Tags auth
// @Accept json
// @Produce json
// @Param request body pb.LoginRequest true "Registration details"
// @Success 200 {object} pb.LoginResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/register [post]
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
    h.logger.Info("RegisterHandler")

    var req pb.LoginRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.IndentedJSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    resp, err := h.auth.Login(c, &req)
    if err != nil {
        c.IndentedJSON(500, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.IndentedJSON(200, resp)
}

// @Summary Login user
// @Description Login an existing user
// @Tags auth
// @Accept json
// @Produce json
// @Success 400 {object} gin.H
// @Router /auth/login [post]
func (h *AuthHandler) LoginHandler(c *gin.Context) {
    h.logger.Info("LoginHandler")
    c.IndentedJSON(400, gin.H{"Message": "Not implemented"})
}

// @Summary Get user profile
// @Description Get profile information for a specific user
// @Tags profile
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} pb.GetProfileResponse
// @Failure 500 {object} gin.H
// @Router /user/profile/{id} [get]
func (h *AuthHandler) GetProfileHandler(c *gin.Context) {
    h.logger.Info("GetProfileHandler")
    id := c.Param("id")

    resp, err := h.auth.GetProfile(c, &pb.GetProfileRequest{
        UserId: id,
    })

    if err != nil {
        c.IndentedJSON(500, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.IndentedJSON(200, resp)
}

// @Summary Update user profile
// @Description Update profile information for a specific user
// @Tags profile
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body pb.UpdateProfileRequest true "Profile update details"
// @Success 200 {object} pb.UpdateProfileResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /user/profile/{id} [put]
func (h *AuthHandler) UpdateProfileHandler(c *gin.Context) {
    h.logger.Info("UpdateProfileHandler")
    id := c.Param("id")

    var req = pb.UpdateProfileRequest{
        UserId: id,
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.IndentedJSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    resp, err := h.auth.UpdateProfile(c, &req)
    if err != nil {
        c.IndentedJSON(500, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.IndentedJSON(200, resp)
}

// @Summary Reset password
// @Description Reset user password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body pb.ResetPasswordRequest true "Password reset details"
// @Success 200 {object} pb.ResetPasswordResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /user/reset-password [post]
func (h *AuthHandler) ResetPasswordHandler(c *gin.Context) {
    h.logger.Info("ResetPasswordHandler")

    var req = pb.ResetPasswordRequest{}

    if err := c.ShouldBindJSON(&req); err != nil {
        c.IndentedJSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    resp, err := h.auth.ResetPassword(c, &req)
    if err != nil {
        c.IndentedJSON(500, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.IndentedJSON(200, resp)
}
