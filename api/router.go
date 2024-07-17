// Package api API.
//
// @title LocalEats API
// @version 1.0
// @description API Endpoints for LocalEats
// @termsOfService http://swagger.io/terms/
//
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
//
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host localhost:8080
// @BasePath /
// @schemes http https
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package api

import (
	"fmt"
	"log/slog"
	"strings"

	"gateway/api/handler"
	"gateway/internal/config"

	_ "gateway/api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Claims struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("your_secret_key_here")

type API struct{}

func New() *API {
	return &API{}
}

// @security BearerAuth
func (a *API) NewRouter(config *config.Config, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Swagger
	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url, ginSwagger.PersistAuthorization(true)))

	router.Use(VerifyJWTMiddleware)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := handler.NewHandlerConfig(logger, config)

	// Public routes
	auth := router.Group("/auth")
	{
		auth.POST("/reset-password-access", handler.AuthHandler.GetAccessToResetPassword)
		auth.POST("/register", handler.AuthHandler.RegisterHandler)
		auth.POST("/login", handler.AuthHandler.LoginHandler)
		auth.GET("/is/valid/token/:user_id", handler.AuthHandler.IsValidToken)
		auth.POST("/logout/:user_id", handler.AuthHandler.Logout)
		auth.POST("/reset-password", handler.AuthHandler.ResetPasswordHandler)
	}

	// Auth-Service
	user := router.Group("/user")
	{
		user.GET("/profile/:id", handler.AuthHandler.GetProfileHandler)
		user.PUT("/profile/:id", handler.AuthHandler.UpdateProfileHandler)
		user.GET("/:user_id/activity", handler.OrderHandler.GetUserActivityHandler)
	}

	kitchen := router.Group("/kitchen")
	{
		kitchen.POST("", handler.AuthHandler.CreateKitchenHandler)
		kitchen.PUT("/:kitchen_id", handler.AuthHandler.UpdateKitchenHandler)
		kitchen.GET("/:kitchen_id", handler.AuthHandler.GetKitchenHandler)
		kitchen.POST("/search", handler.AuthHandler.SearchKitchensHandler)
		kitchen.POST("/:kitchen_id/working-hours", handler.OrderHandler.CreateKitchenWorkingHoursHandler)
	}

	// Order-Service
	dish := router.Group("/dish")
	{
		dish.POST("", handler.OrderHandler.AddDishHandler)
		dish.PUT("/:dish_id", handler.OrderHandler.UpdateDishHandler)
		dish.DELETE("/:dish_id", handler.OrderHandler.DeleteDishHandler)
		dish.GET("/kitchen/:kitchen_id", handler.OrderHandler.ListDishesHandler)
		dish.PUT("/:dish_id/nutrition", handler.OrderHandler.UpdateDishNutritionInfoHandler)
	}

	order := router.Group("/order")
	{
		order.POST("", handler.OrderHandler.CreateOrderHandler)
		order.PUT("/:order_id/status", handler.OrderHandler.UpdateOrderStatusHandler)
		order.GET("/user/:user_id", handler.OrderHandler.ListOfOrdersHandler)
		order.GET("/kitchen/:kitchen_id", handler.OrderHandler.GetOrderByKitchenIdHandler)
	}

	review := router.Group("/review")
	{
		review.POST("", handler.OrderHandler.AddReviewHandler)
		review.POST("/kitchen/:kitchen_id", handler.OrderHandler.ListReviewsHandler)
	}

	payment := router.Group("/payment")
	{
		payment.POST("", handler.OrderHandler.CreatePaymentHandler)
	}

	return router
}

func VerifyJWTMiddleware(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/auth/login" || ctx.Request.URL.Path == "/auth/register" ||ctx.Request.URL.Path == "/auth/reset-password-access"||ctx.Request.URL.Path == "/auth/reset-password" || strings.HasPrefix(ctx.Request.URL.Path, "/swagger") {
		ctx.Next()
		return
	}

	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		ctx.JSON(401, gin.H{"error": "Authorization header is missing"})
		ctx.Abort()
		return
	}

	if !strings.HasPrefix(tokenStr, "Bearer ") {
		ctx.JSON(401, gin.H{"error": "Invalid token format"})
		ctx.Abort()
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		ctx.JSON(401, gin.H{"error": fmt.Sprintf("Invalid token: %v", err)})
		ctx.Abort()
		return
	}

	if !token.Valid {
		ctx.JSON(401, gin.H{"error": "Token is not valid"})
		ctx.Abort()
		return
	}

	ctx.Set("user_id", claims.UserID)
	ctx.Set("username", claims.Username)

	ctx.Next()
}
