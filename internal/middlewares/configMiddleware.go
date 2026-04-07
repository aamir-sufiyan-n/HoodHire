package middlewares

import (
    "hoodhire/internal/repositories"
    "github.com/gofiber/fiber/v3"
)

func FeatureMiddleware(configRepo *repositories.WebRepo, key string) fiber.Handler {
    return func(c fiber.Ctx) error {
        config, err := configRepo.GetConfigByKey(key)
        if err != nil {
            // if config key doesn't exist, allow by default
            return c.Next()
        }

        if !config.IsActive {
            return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
                "error": "this feature is currently unavailable",
            })
        }

        return c.Next()
    }
}