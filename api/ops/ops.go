// Package ops implements the Open Proxy Sharing API v1.
package ops

import (
	"encoding/base64"

	"github.com/database64128/proxy-sharing-go/ent"
	"github.com/database64128/proxy-sharing-go/ent/account"
	"github.com/database64128/proxy-sharing-go/ent/registrationtoken"
	"github.com/database64128/proxy-sharing-go/httpx"
	"github.com/database64128/proxy-sharing-go/tokens"
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

	router.Get("/account", newGetCurrentAccountHandler())
	router.Get("/accounts", newListAccountsHandler(client))
	router.Get("/accounts/:id", newGetAccountHandler(client))
}

type registerRefreshResponse struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	AccessToken  []byte `json:"access_token"`
	RefreshToken []byte `json:"refresh_token"`
}

func registerRefreshResponseFromEntAccount(account *ent.Account) registerRefreshResponse {
	return registerRefreshResponse{
		ID:           account.ID,
		Username:     account.Username,
		AccessToken:  account.AccessToken,
		RefreshToken: account.RefreshToken,
	}
}

func newRegisterHandler(client *ent.Client, logger *zap.Logger) fiber.Handler {
	type request struct {
		RegistrationToken []byte `json:"registration_token"`
		Username          string `json:"username"`
	}

	return func(c *fiber.Ctx) error {
		var req request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(httpx.StandardError{Message: err.Error()})
		}

		registrationToken, err := client.RegistrationToken.Query().Where(registrationtoken.Token(req.RegistrationToken)).Only(c.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusUnauthorized).JSON(httpx.StandardError{Message: "invalid registration token"})
			}
			return err
		}

		accessToken, refreshToken := tokens.NewAccessTokenAndRefreshTokenBytes()

		account, err := client.Account.Create().
			SetUsername(req.Username).
			SetAccessToken(accessToken).
			SetRefreshToken(refreshToken).
			SetRegistrationToken(registrationToken).
			Save(c.Context())
		if err != nil {
			if ent.IsValidationError(err) {
				return c.Status(fiber.StatusBadRequest).JSON(httpx.StandardError{Message: err.Error()})
			}
			if ent.IsConstraintError(err) {
				return c.Status(fiber.StatusConflict).JSON(httpx.StandardError{Message: err.Error()})
			}
			return err
		}

		logger.Info("Registered new account",
			zap.String("username", req.Username),
			zap.String("registrationToken", registrationToken.Name),
		)

		return c.JSON(registerRefreshResponseFromEntAccount(account))
	}
}

func newRefreshHandler(client *ent.Client, logger *zap.Logger) fiber.Handler {
	type request struct {
		RefreshToken []byte `json:"refresh_token"`
	}

	return func(c *fiber.Ctx) error {
		var req request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(httpx.StandardError{Message: err.Error()})
		}

		account, err := client.Account.Query().Where(account.RefreshToken(req.RefreshToken)).Only(c.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusUnauthorized).JSON(httpx.StandardError{Message: "invalid refresh token"})
			}
			return err
		}

		accessToken := tokens.NewTokenBytes()

		account, err = account.Update().
			SetAccessToken(accessToken).
			Save(c.Context())
		if err != nil {
			return err
		}

		logger.Info("Refreshed account access token", zap.String("username", account.Username))

		return c.JSON(registerRefreshResponseFromEntAccount(account))
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
			return c.Status(fiber.StatusUnauthorized).JSON(httpx.StandardError{Message: "missing or malformed access token"})
		}

		token, err := base64.StdEncoding.DecodeString(auth[len(bearerPrefix):])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(httpx.StandardError{Message: "malformed access token: " + err.Error()})
		}

		account, err := client.Account.Query().Where(account.AccessToken(token)).Only(c.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusUnauthorized).JSON(httpx.StandardError{Message: "invalid access token"})
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

type accountResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func accountResponseFromEntAccount(account *ent.Account) accountResponse {
	return accountResponse{
		ID:       account.ID,
		Username: account.Username,
	}
}

func newGetCurrentAccountHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		account := accountFromCtx(c)
		return c.JSON(accountResponseFromEntAccount(account))
	}
}

func newListAccountsHandler(client *ent.Client) fiber.Handler {
	type response struct {
		Accounts []accountResponse `json:"accounts"`
	}

	return func(c *fiber.Ctx) error {
		aq := client.Account.Query()
		if username := c.Query("username"); username != "" {
			aq = aq.Where(account.Username(username))
		}

		accounts, err := aq.All(c.Context())
		if err != nil {
			return err
		}

		resp := response{Accounts: make([]accountResponse, len(accounts))}
		for i, account := range accounts {
			resp.Accounts[i] = accountResponseFromEntAccount(account)
		}
		return c.JSON(resp)
	}
}

func newGetAccountHandler(client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(httpx.StandardError{Message: err.Error()})
		}

		account, err := client.Account.Get(c.Context(), id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(httpx.StandardError{Message: "account not found"})
			}
			return err
		}

		return c.JSON(accountResponseFromEntAccount(account))
	}
}
