package middlewares

import (
	"hoodhire/internal/repositories"

	"github.com/gofiber/fiber/v3"
)

func PermissionMiddleware(roleRepo *repositories.RoleRepo, permission string) fiber.Handler {
    return func(c fiber.Ctx) error {
        role, ok := c.Locals("role").(string)
        if !ok || role == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "unauthorized",
            })
		}

        // fetch the role and its permissions from DB
        roleData, err := roleRepo.GetRoleByName(role)
        if err != nil {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                "error": "role not found",
            })
        }

        // check if the required permission is allowed
        for _, rp := range roleData.RolePermissions {
            if rp.PermissionID != 0 && rp.IsAllowed {
                // fetch permission name to match
                var permName string
                roleRepo.DB.Table("permissions").
                    Select("name").
                    Where("id = ?", rp.PermissionID).
                    Scan(&permName)

                if permName == permission {
                    return c.Next()
                }
            }
        }

        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": "forbidden - you do not have access to this resource",
        })
    }
}