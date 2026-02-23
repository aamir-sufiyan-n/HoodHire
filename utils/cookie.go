package utils

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func SetCookie(c fiber.Ctx,access, refresh string) {
	c.Cookie(&fiber.Cookie{
		Name: "access-token",
		Value: access,
		Expires: time.Now().Add(time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})
	c.Cookie(&fiber.Cookie{
		Name: "refresh-token",
		Value: refresh,
		Expires: time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})
}
func ClearCookie(c fiber.Ctx){
		c.Cookie(&fiber.Cookie{
		Name: "access-token",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})
	c.Cookie(&fiber.Cookie{
		Name: "refresh-token",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})
}