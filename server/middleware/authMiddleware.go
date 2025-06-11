package middleware

import (
	"strings"

	"github.com/BoomTHDev/wear-pos-server/config"
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	"github.com/BoomTHDev/wear-pos-server/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return custom.ErrUnauthorized("UNAUTHORIZED", "Authorization header is required", nil)
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return custom.ErrUnauthorized("UNAUTHORIZED", "Invalid authorization header format", nil)
		}

		token := tokenParts[1]
		claims, err := util.VerifyToken(token, config.ConfigGetting().Server.JWTSecret)
		if err != nil {
			return custom.ErrUnauthorized("UNAUTHORIZED", "Invalid token", err)
		}

		c.Locals("user_id", claims.UserID)

		return c.Next()
	}
}
