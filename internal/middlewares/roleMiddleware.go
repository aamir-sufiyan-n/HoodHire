package middlewares

import "github.com/gofiber/fiber/v3"

func RoleMiddleware(roles ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "k3ejhbc3ejfc3e",
			})
		}

		for _, r := range roles {
			if r == role {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "forbidden - you do not have permission",
		})
	}
}