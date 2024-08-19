package authhandler

import (
	"gateway/genprotos/auth_pb"
	"gateway/internal/config"
	"log/slog"
)

type AuthHandler struct {
	logger *slog.Logger
	auth   auth_pb.AuthServiceClient
	config *config.Config
}

func NewAuthHandler(config *config.Config, logger *slog.Logger, auth auth_pb.AuthServiceClient) *AuthHandler {
	return &AuthHandler{
		logger: logger,
		auth:   auth,
		config: config,
	}
}
