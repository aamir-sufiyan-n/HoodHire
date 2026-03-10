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

	app.Get("/jobs", handler.JobHandlers.GetActiveJobs)
	app.Get("/jobs/category/:categoryID", handler.JobHandlers.GetJobsByCategory)
	app.Get("/jobs/locality/:locality", handler.JobHandlers.GetJobsByLocality)
	app.Get("/jobs/:id", handler.JobHandlers.GetJobByID)
	app.Get("/categories", handler.SeekerHandler.GetJobCategories)
	app.Get("/businesses", handler.HirerHandler.GetAllBusinesses)
	app.Get("/businesses/:id", handler.HirerHandler.GetBusinessByID)
	

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

		seekerApi.Put("/categories", handler.SeekerHandler.UpdateJobInterests)

		seekerApi.Post("/jobs/:id/apply", handler.JobHandlers.ApplyToJob)
		seekerApi.Get("/applications", handler.JobHandlers.GetMyApplications)
		seekerApi.Delete("/applications/:applicationID", handler.JobHandlers.WithdrawApplication)
	}
	app.Get("/seeker/:id", handler.SeekerHandler.GetSeekerByID)
	hirerApi := app.Group("/hirer", middlewares.AuthMiddleware, middlewares.RoleMiddleware("hirer"))
	{
		hirerApi.Post("/profile", handler.HirerHandler.CreateProfile)
		hirerApi.Get("/profile", handler.HirerHandler.GetHirerProfile)
		hirerApi.Put("/profile", handler.HirerHandler.UpdateProfile)
		hirerApi.Delete("/profile", handler.HirerHandler.DeleteProfile)

		hirerApi.Get("/seeker/:id",handler.SeekerHandler.GetSeekerByID)
		hirerApi.Post("/jobs", handler.JobHandlers.CreateJob)
		hirerApi.Get("/jobs", handler.JobHandlers.GetMyJobs)
		hirerApi.Patch("/jobs/applications/:applicationID/status", handler.JobHandlers.UpdateApplicationStatus)
		hirerApi.Put("/jobs/:id", handler.JobHandlers.UpdateJob)
		hirerApi.Patch("/jobs/:id/status", handler.JobHandlers.UpdateJobStatus)
		hirerApi.Delete("/jobs/:id", handler.JobHandlers.DeleteJob)
		hirerApi.Get("/jobs/:id/applications", handler.JobHandlers.GetApplicationsForJob)
	}
}
