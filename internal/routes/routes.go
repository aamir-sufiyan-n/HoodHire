package routes

import (
	"hoodhire/internal/app"
	"hoodhire/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, handler *app.APP) {

	auth := app.Group("/auth")
	{
		auth.Post("/login", handler.AuthHandler.Login)
		auth.Post("/logout", handler.AuthHandler.Logout)
		seeker := auth.Group("/seeker")
		{
			seeker.Post("/send-otp", func(c fiber.Ctx) error {
				c.Locals("role", "seeker")
				return handler.AuthHandler.SendOTP(c)
			})
			seeker.Post("/verify", handler.AuthHandler.Signup)
			seeker.Post("/resend-otp", handler.AuthHandler.ResendOTP)
		}
		hirer := auth.Group("/hirer")
		{
			hirer.Post("/send-otp", func(c fiber.Ctx) error {
				c.Locals("role", "hirer") // Set role in context
				return handler.AuthHandler.SendOTP(c)
			})
			hirer.Post("/verify", handler.AuthHandler.Signup)
			hirer.Post("/resend-otp", handler.AuthHandler.ResendOTP)
		}
	}

	seekerApi := app.Group("/seeker",middlewares.AuthMiddleware,middlewares.RoleMiddleware("seeker"))
	{
		seekerApi.Post("/profile", handler.SeekerHandler.SetupSeekerProfile,middlewares.AuthMiddleware)
		seekerApi.Get("/profile", handler.SeekerHandler.GetProfile,middlewares.AuthMiddleware)
		seekerApi.Put("/profile", handler.SeekerHandler.UpdateSeeker,middlewares.AuthMiddleware)
		seekerApi.Delete("/profile", handler.SeekerHandler.DeleteSeeker,middlewares.AuthMiddleware)
	}

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "healthy",
			"message": "HoodHire API is running",
		})
	}) 	
}
