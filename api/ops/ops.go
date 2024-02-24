// Package ops implements the Open Proxy Sharing API v1.
package ops

import (
	"encoding/base64"

	"github.com/database64128/proxy-sharing-go/ent"
	"github.com/database64128/proxy-sharing-go/ent/account"
	"github.com/database64128/proxy-sharing-go/ent/registrationtoken"
	"github.com/database64128/proxy-sharing-go/httphelper"
	"github.com/database64128/proxy-sharing-go/tokenhelper"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Config is the configuration for the Open Proxy Sharing API.
type Config struct{}

// RegisterRoutes registers the Open Proxy Sharing API routes.
func (c Config) RegisterRoutes(router fiber.Router, client *ent.Client, logger *zap.Logger) {
	router.Post("/register", newRegisterHandler(client, logger))
	router.Post("/refresh", newRefreshHandler(client, logger))

	router.Use(newAuthMiddleware(client))

	router.Get("/account", newGetAccountHandler())
}

type registerRefreshResponse struct {
	Username     string `json:"username"`
	AccessToken  []byte `json:"access_token"`
	RefreshToken []byte `json:"refresh_token"`
}

func newRegisterHandler(client *ent.Client, logger *zap.Logger) fiber.Handler {
	type request struct {
		RegistrationToken []byte `json:"registration_token"`
		Username          string `json:"username"`
	}

	return func(c *fiber.Ctx) error {
		var req request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(httphelper.StandardError{Message: err.Error()})
		}

		registrationToken, err := client.RegistrationToken.Query().Where(registrationtoken.Token(req.RegistrationToken)).Only(c.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusUnauthorized).JSON(httphelper.StandardError{Message: "invalid registration token"})
			}
			return err
		}

		accessToken, refreshToken, err := tokenhelper.NewAccessTokenAndRefreshTokenBytes()
		if err != nil {
			return err
		}

		_, err = client.Account.Create().
			SetUsername(req.Username).
			SetAccessToken(accessToken).
			SetRefreshToken(refreshToken).
			SetRegistrationToken(registrationToken).
			Save(c.Context())
		if err != nil {
			if ent.IsValidationError(err) {
				return c.Status(fiber.StatusBadRequest).JSON(httphelper.StandardError{Message: err.Error()})
			}
			if ent.IsConstraintError(err) {
				return c.Status(fiber.StatusConflict).JSON(httphelper.StandardError{Message: err.Error()})
			}
			return err
		}

		logger.Info("Registered new account",
			zap.String("username", req.Username),
			zap.String("registrationToken", registrationToken.Name),
		)

		return c.JSON(registerRefreshResponse{
			Username:     req.Username,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
	}
}

func newRefreshHandler(client *ent.Client, logger *zap.Logger) fiber.Handler {
	type request struct {
		RefreshToken []byte `json:"refresh_token"`
	}

	return func(c *fiber.Ctx) error {
		var req request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(httphelper.StandardError{Message: err.Error()})
		}

		account, err := client.Account.Query().Where(account.RefreshToken(req.RefreshToken)).Only(c.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusUnauthorized).JSON(httphelper.StandardError{Message: "invalid refresh token"})
			}
			return err
		}

		accessToken, err := tokenhelper.NewTokenBytes()
		if err != nil {
			return err
		}

		account, err = account.Update().
			SetAccessToken(accessToken).
			Save(c.Context())
		if err != nil {
			return err
		}

		logger.Info("Refreshed account access token", zap.String("username", account.Username))

		return c.JSON(registerRefreshResponse{
			Username:     account.Username,
			AccessToken:  accessToken,
			RefreshToken: account.RefreshToken,
		})
	}
}

func newAuthMiddleware(client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		const (
			bearerPrefix      = "Bearer "
			tokenFormat       = "jTGA1dfSmqtyNsLIrM9zkPIdjvw76I1z7LqJAAg13TU="
			bearerTokenFormat = bearerPrefix + tokenFormat
		)

		auth := c.Get(fiber.HeaderAuthorization)
		if len(auth) != len(bearerTokenFormat) || auth[:len(bearerPrefix)] != bearerPrefix {
			return c.Status(fiber.StatusUnauthorized).JSON(httphelper.StandardError{Message: "missing or malformed access token"})
		}

		token, err := base64.StdEncoding.DecodeString(auth[len(bearerPrefix):])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(httphelper.StandardError{Message: "malformed access token: " + err.Error()})
		}

		account, err := client.Account.Query().Where(account.AccessToken(token)).Only(c.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusUnauthorized).JSON(httphelper.StandardError{Message: "invalid access token"})
			}
			return err
		}

		c.Locals("account", account)
		return c.Next()
	}
}

func accountFromCtx(c *fiber.Ctx) *ent.Account {
	return c.Locals("account").(*ent.Account)
}

func newGetAccountHandler() fiber.Handler {
	type response struct {
		Username string `json:"username"`
	}

	return func(c *fiber.Ctx) error {
		account := accountFromCtx(c)
		return c.JSON(response{Username: account.Username})
	}
}
