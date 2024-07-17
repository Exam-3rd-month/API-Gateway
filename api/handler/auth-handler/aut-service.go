package authhandler

import (
	sendcode "gateway/api/sendCode"
	pb "gateway/genprotos/auth_pb"
	"time"

	r "gateway/api/redis"

	"github.com/gin-gonic/gin"
)

// RegisterHandler godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags auth
// @Accept json
// @Produce json
// @Param request body pb.RegisterRequest true "Registration details"
// @Success 200 {object} pb.RegisterResponse "Successful registration response"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	h.logger.Info("RegisterHandler")

	var req pb.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.auth.Register(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Login user
// @Description Authenticates a user and returns a token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body pb.LoginRequest true "Login credentials"
// @Success 200 {object} pb.LoginResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/login [post]
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	h.logger.Info("LoginHandler")
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

// @Summary Get user profile
// @Description Get the profile of a user by their ID
// @Tags user
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
// @Description Update the profile of a user by their ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body pb.UpdateProfileRequest true "Updated profile information"
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

// @Summary Reset user password
// @Description Reset the password for a user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body pb.ResetPasswordRequest true "Password reset information"
// @Success 200 {object} pb.ResetPasswordResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/reset-password [post]
func (h *AuthHandler) ResetPasswordHandler(c *gin.Context) {
	h.logger.Info("ResetPasswordHandler")

	var req = pb.ResetPasswordRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	code, err := r.GetVerificationCode(req.Email)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.VerifyCode != code {
		c.IndentedJSON(400, gin.H{
			"error": "Verify code",
		})
		return
	}

	if req.NewPassword != req.ConfirmPassword {
		c.IndentedJSON(400, gin.H{
			"error": "Passwords do not match",
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

// @Summary Logout user
// @Description Logs out a user by their user ID
// @Tags auth
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} pb.LogoutResponse
// @Failure 500 {object} gin.H
// @Router /auth/logout/{user_id} [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	h.logger.Info("LogoutHandler")
	user_id := c.Param("user_id")

	resp, err := h.auth.Logout(c, &pb.LogoutRequest{
		UserId: user_id,
	})

	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Check token validity
// @Description Checks if a user's token is valid
// @Tags auth
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} pb.IsValidTokenResponse
// @Failure 500 {object} gin.H
// @Router /auth/is-valid-token/{user_id} [get]
func (h *AuthHandler) IsValidToken(c *gin.Context) {
	h.logger.Info("IsValidTokenHandler")
	user_id := c.Param("user_id")

	resp, err := h.auth.IsValidToken(c, &pb.IsValidTokenRequest{
		UserId: user_id,
	})

	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, resp)
}

// GetAccessToResetPassword godoc
// @Summary Get access to reset password
// @Description Send a verification code to the user's email for password reset
// @Tags auth
// @Accept json
// @Produce json
// @Param request body pb.IsValidUserRequest true "User email"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/reset-password-access [post]
func (h *AuthHandler) GetAccessToResetPassword(c *gin.Context) {
	h.logger.Info("GetAccessToResetPasswordHandler")

	var req pb.IsValidUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.auth.IsValidUser(c, &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	code, err := sendcode.SendVerificationCode(req.Email)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = r.SaveVerificationCode(req.Email, code, 2*time.Minute)
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "We have sent you a code to your email"})
}
