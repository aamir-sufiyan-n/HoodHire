package middlewares

import (
	"fmt"
	"hoodhire/config"
	"hoodhire/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func ServiceAuthMiddleware(c fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	var claims = &utils.Claims{}
	t, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(config.AppConfig.JwtKey), nil
	})

	if err != nil || !t.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "invalid or expired token"})
	}

	c.Locals("userID", claims.UserID)
	c.Locals("role", claims.Role)
	return c.Next()
}