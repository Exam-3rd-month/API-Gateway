package authhandler

import (
	"gateway/genprotos/auth_pb"
	"log/slog"
)

type AuthHandler struct {
	logger *slog.Logger
	auth   auth_pb.AuthServiceClient
}

func NewAuthHandler(logger *slog.Logger, auth auth_pb.AuthServiceClient) *AuthHandler {
	return &AuthHandler{
		logger: logger,
		auth:   auth,
	}
}
