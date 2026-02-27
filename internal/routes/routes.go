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

	seekerApi := app.Group("/seeker", middlewares.AuthMiddleware, middlewares.RoleMiddleware("seeker"))
	{
		seekerApi.Post("/profile", handler.SeekerHandler.SetupSeekerProfile)
		seekerApi.Get("/profile", handler.SeekerHandler.GetProfile)
		seekerApi.Put("/profile", handler.SeekerHandler.UpdateSeeker)
		seekerApi.Delete("/profile", handler.SeekerHandler.DeleteSeeker)

		seekerApi.Put("/education", handler.SeekerHandler.UpsertEducation)

		
		seekerApi.Post("/experience", handler.SeekerHandler.AddWorkExperience)
		seekerApi.Get("/experience", handler.SeekerHandler.GetWorkExperiences)
		seekerApi.Delete("/experience/:id", handler.SeekerHandler.DeleteWorkExperience)

		seekerApi.Put("/preference", handler.SeekerHandler.UpsertWorkPreference)
		seekerApi.Get("/preference", handler.SeekerHandler.GetWorkPreference)
	}
	hirerApi := app.Group("/hirer", middlewares.AuthMiddleware, middlewares.RoleMiddleware("hirer"))
	{
		hirerApi.Post("/profile", handler.HirerHandler.CreateProfile)
		hirerApi.Get("/profile", handler.HirerHandler.GetHirerProfile)
		hirerApi.Put("/profile", handler.HirerHandler.UpdateProfile)
		hirerApi.Delete("/profile", handler.HirerHandler.DeleteProfile)
	}

}
