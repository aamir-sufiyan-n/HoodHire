package middlewares

import (
	"hoodhire/database"
	"hoodhire/structures/models"
	"hoodhire/utils"

	"github.com/gofiber/fiber/v3"
)

func BlockMiddleware(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}
	if user.IsBlocked {
		utils.ClearCookie(c)
		return c.Status(403).JSON(fiber.Map{"error": "your account has been blocked"})
	}
	return c.Next()
}