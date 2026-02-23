package middlewares

import (
	"fmt"
	"hoodhire/config"
	"hoodhire/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c fiber.Ctx) error{
	tokenString:=c.Cookies("access-token")
	if tokenString == ""{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":"you are unautharized, please Log in",
		})
	}
	var claims = &utils.Claims{}
	t,err:=jwt.ParseWithClaims(tokenString,claims,func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(config.AppConfig.JwtKey), nil
	 })

	 if err != nil || !t.Valid{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"invalid or expired token"})
	 }
	 c.Locals("userID",claims.UserID)
	 c.Locals("username",claims.Username)
	 c.Locals("email",claims.Email)
	 c.Locals("role",claims.Role)

	 return c.Next()
}