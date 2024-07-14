package api

import (
	"fmt"
	"log/slog"
	"strings"

	"gateway/api/handler"
	"gateway/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	_ "gateway/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var jwtKey = []byte("secret-key")

type API struct{}

func New() *API {
	return &API{}
}

// @title LocalEats
// @version 1.0
// @description API Gateway of LocalEats
// @host localhost:8080
// BasePath: /
func (a *API) NewRouter(config *config.Config, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	url := ginSwagger.URL("swagger/docs.json") 
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := handler.NewHandlerConfig(logger, config)
	router.POST("/register", handler.AuthHandler.RegisterHandler)
	router.POST("/login", handler.AuthHandler.LoginHandler)

	router.Use(VerifyJWTMiddleware)

	user := router.Group("/user")
	{
		user.GET("/profile/:id", handler.AuthHandler.GetProfileHandler)
		user.PUT("/profile/:id", handler.AuthHandler.UpdateProfileHandler)
		user.POST("/reset-password", handler.AuthHandler.ResetPasswordHandler)
	}

	kitchen := router.Group("/kitchen")
	{
		kitchen.POST("", handler.AuthHandler.CreateKitchenHandler)
		kitchen.PUT("/:kitchen_id", handler.AuthHandler.UpdateKitchenHandler)
		kitchen.GET("/:kitchen_id", handler.AuthHandler.GetKitchenHandler)
	}

	return router
}

func VerifyJWTMiddleware(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")

	if !strings.HasPrefix(tokenStr, "Bearer ") {
		ctx.IndentedJSON(401, gin.H{"error": "unauthorized"})
		ctx.Abort()
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		ctx.IndentedJSON(401, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
