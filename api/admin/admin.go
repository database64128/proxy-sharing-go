package admin

import (
	"time"

	"github.com/database64128/proxy-sharing-go/ent"
	"github.com/database64128/proxy-sharing-go/ent/account"
	"github.com/database64128/proxy-sharing-go/httphelper"
	"github.com/database64128/proxy-sharing-go/tokenhelper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"go.uber.org/zap"
)

// Config stores the configuration for the admin API.
type Config struct {
	// AccessTokens is the list of access tokens for the admin API.
	// These tokens cannot be used with non-admin endpoints.
	AccessTokens []string `json:"accessTokens"`
}

// RegisterRoutes registers the admin API routes.
func (c Config) RegisterRoutes(router fiber.Router, client *ent.Client, logger *zap.Logger) {
	router.Use(c.newAuthMiddleware())

	rtg := router.Group("/registration-tokens")
	rtg.Get("/", newListRegistrationTokensHandler(client))
	rtg.Get("/:id", newGetRegistrationTokenHandler(client))
	rtg.Post("/", newCreateRegistrationTokenHandler(client))
	rtg.Patch("/:id", newRenameRegistrationTokenHandler(client))
	rtg.Delete("/:id", newDeleteRegistrationTokenHandler(client, logger))
}

func (c Config) newAuthMiddleware() fiber.Handler {
	tokenSet := make(map[string]struct{}, len(c.AccessTokens))
	for _, token := range c.AccessTokens {
		tokenSet[token] = struct{}{}
	}
	return keyauth.New(keyauth.Config{
		Validator: func(c *fiber.Ctx, s string) (bool, error) {
			_, ok := tokenSet[s]
			return ok, nil
		},
	})
}

type registrationToken struct {
	ID         int       `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Name       string    `json:"name"`
	Token      []byte    `json:"token"`
}

func registrationTokenFromEnt(token *ent.RegistrationToken) registrationToken {
	return registrationToken{
		ID:         token.ID,
		CreateTime: token.CreateTime,
		UpdateTime: token.UpdateTime,
		Name:       token.Name,
		Token:      token.Token,
	}
}

func newListRegistrationTokensHandler(client *ent.Client) fiber.Handler {
	type response struct {
		Tokens []registrationToken `json:"tokens"`
	}

	return func(c *fiber.Ctx) error {
		tokens, err := client.RegistrationToken.Query().All(c.Context())
		if err != nil {
			return err
		}

		resp := response{Tokens: make([]registrationToken, len(tokens))}
		for i, token := range tokens {
			resp.Tokens[i] = registrationTokenFromEnt(token)
		}
		return c.JSON(resp)
	}
}

func newGetRegistrationTokenHandler(client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(httphelper.StandardError{Message: err.Error()})
		}

		token, err := client.RegistrationToken.Get(c.Context(), id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(httphelper.StandardError{Message: "token not found"})
			}
			return err
		}

		return c.JSON(registrationTokenFromEnt(token))
	}
}

type createUpdateRegistrationTokenRequest struct {
	Name string `json:"name"`
}

func newCreateRegistrationTokenHandler(client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req createUpdateRegistrationTokenRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(httphelper.StandardError{Message: err.Error()})
		}

		b := tokenhelper.NewTokenBytes()

		token, err := client.RegistrationToken.Create().
			SetName(req.Name).
			SetToken(b).
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

		return c.Status(fiber.StatusCreated).JSON(registrationTokenFromEnt(token))
	}
}

func newRenameRegistrationTokenHandler(client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(httphelper.StandardError{Message: err.Error()})
		}

		var req createUpdateRegistrationTokenRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(httphelper.StandardError{Message: err.Error()})
		}

		token, err := client.RegistrationToken.UpdateOneID(id).
			SetName(req.Name).
			Save(c.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(httphelper.StandardError{Message: "token not found"})
			}
			if ent.IsValidationError(err) {
				return c.Status(fiber.StatusBadRequest).JSON(httphelper.StandardError{Message: err.Error()})
			}
			if ent.IsConstraintError(err) {
				return c.Status(fiber.StatusConflict).JSON(httphelper.StandardError{Message: err.Error()})
			}
			return err
		}

		return c.JSON(registrationTokenFromEnt(token))
	}
}

func newDeleteRegistrationTokenHandler(client *ent.Client, logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(httphelper.StandardError{Message: err.Error()})
		}

		if c.QueryBool("purge") {
			// The nuclear option: delete all accounts registered with this token.
			n, err := client.Account.Delete().Where(account.RegistrationTokenID(id)).Exec(c.Context())
			if err != nil {
				return err
			}
			logger.Info("Deleted accounts registered with token", zap.Int("tokenID", id), zap.Int("count", n))
		}

		err = client.RegistrationToken.DeleteOneID(id).Exec(c.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(httphelper.StandardError{Message: "token not found"})
			}
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
