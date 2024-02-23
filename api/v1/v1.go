package v1

import (
	"github.com/database64128/proxy-sharing-go/ent"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Register registers the v1 routes.
func Register(logger *zap.Logger, client *ent.Client, v1 fiber.Router) {
	// TODO
}

// StandardError is the standard error response.
type StandardError struct {
	Message string `json:"error"`
}
