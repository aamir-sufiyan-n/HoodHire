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
	app.Get("/businesses/:businessID/reviews", handler.FollowHandler.GetReviewsByBusiness)
	app.Get("/auth/verify", middlewares.ServiceAuthMiddleware, func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"user_id": c.Locals("userID"),
			"role":    c.Locals("role"),
		})
	})
	app.Get("/bonds/check", handler.BondHandler.CheckActiveBond)

	seekerApi := app.Group("/seeker", middlewares.AuthMiddleware, middlewares.RoleMiddleware("seeker"))
	{
		seekerApi.Post("/profile", handler.SeekerHandler.SetupSeekerProfile)
		seekerApi.Get("/profile", handler.SeekerHandler.GetProfile)
		seekerApi.Put("/profile", handler.SeekerHandler.UpdateSeeker)
		seekerApi.Delete("/profile", handler.SeekerHandler.DeleteSeeker)
		seekerApi.Patch("/profile/picture", handler.SeekerHandler.UploadProfilePicture)
		seekerApi.Delete("/profile/picture", handler.SeekerHandler.RemoveProfilePicture)

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

		seekerApi.Post("/follow/:businessID", handler.FollowHandler.FollowBusiness)
		seekerApi.Delete("/follow/:businessID", handler.FollowHandler.UnfollowBusiness)
		seekerApi.Get("/following", handler.FollowHandler.GetFollowedBusinesses)
		seekerApi.Get("/follow/:businessID", handler.FollowHandler.IsFollowing)

		seekerApi.Post("/businesses/:businessID/review", handler.FollowHandler.CreateReview)
		seekerApi.Get("/businesses/:businessID/review", handler.FollowHandler.GetReviewsByBusiness)
		seekerApi.Put("/businesses/:businessID/review", handler.FollowHandler.UpdateReview)
		seekerApi.Delete("/businesses/:businessID/review", handler.FollowHandler.DeleteReview)
		seekerApi.Get("/businesses/:businessID/my-review", handler.FollowHandler.GetMyReview)

		seekerApi.Post("/tickets", handler.TicketHandler.CreateTicket)
		seekerApi.Get("/tickets", handler.TicketHandler.GetMyTickets)
		seekerApi.Delete("/tickets/:ticketID", handler.TicketHandler.DeleteTicket)

		seekerApi.Post("/favorite/:businessID", handler.SeekerHandler.FavoriteBusiness)
		seekerApi.Delete("/favorite/:businessID", handler.SeekerHandler.UnFavoriteBusiness)
		seekerApi.Get("/favorite", handler.SeekerHandler.GetFavoriteBusiness)
		seekerApi.Get("/favorite/:businessID", handler.SeekerHandler.IsFavorited)

		seekerApi.Post("/saved/jobs/:jobID", handler.SeekerHandler.SaveJob)
		seekerApi.Delete("/saved/jobs/:jobID", handler.SeekerHandler.UnsaveJob)
		seekerApi.Get("/saved/jobs", handler.SeekerHandler.GetSavedJobs)
		seekerApi.Get("/saved/jobs/:jobID", handler.SeekerHandler.IsJobSaved)

		seekerApi.Get("/bonds", handler.BondHandler.GetMyBonds)
	}
	app.Get("/seeker/:id", handler.SeekerHandler.GetSeekerByID)
	hirerApi := app.Group("/hirer", middlewares.AuthMiddleware, middlewares.RoleMiddleware("hirer"))
	{
		hirerApi.Post("/profile", handler.HirerHandler.CreateProfile)
		hirerApi.Get("/profile", handler.HirerHandler.GetHirerProfile)
		hirerApi.Put("/profile", handler.HirerHandler.UpdateProfile)
		hirerApi.Delete("/profile", handler.HirerHandler.DeleteProfile)
		hirerApi.Patch("/profile/picture", handler.HirerHandler.UploadProfilePicture)
		hirerApi.Delete("/profile/picture", handler.HirerHandler.RemoveProfilePicture)

		hirerApi.Get("/seeker/:id", handler.SeekerHandler.GetSeekerByID)
		hirerApi.Post("/jobs", handler.JobHandlers.CreateJob)
		hirerApi.Get("/jobs", handler.JobHandlers.GetMyJobs)
		hirerApi.Patch("/jobs/applications/:applicationID/status", handler.JobHandlers.UpdateApplicationStatus)
		hirerApi.Put("/jobs/:id", handler.JobHandlers.UpdateJob)
		hirerApi.Patch("/jobs/:id/status", handler.JobHandlers.UpdateJobStatus)
		hirerApi.Delete("/jobs/:id", handler.JobHandlers.DeleteJob)
		hirerApi.Get("/jobs/:id/applications", handler.JobHandlers.GetApplicationsForJob)

		hirerApi.Post("/tickets", handler.TicketHandler.CreateTicket)
		hirerApi.Get("/tickets", handler.TicketHandler.GetMyTickets)
		hirerApi.Delete("/tickets/:ticketID", handler.TicketHandler.DeleteTicket)

		hirerApi.Get("/bonds", handler.BondHandler.GetHirerBonds)
		hirerApi.Patch("/bonds/:applicationID/deactivate", handler.BondHandler.DeactivateBond)

		hirerApi.Get("/staff", handler.HirerHandler.GetStaff)
		hirerApi.Delete("/staff/:bondID", handler.HirerHandler.RemoveStaff)

	}
}

// adminApi.Get("/tickets", handler.TicketHandler.GetAllTickets)
// adminApi.Get("/tickets/type/:type", handler.TicketHandler.GetTicketsByType)
// adminApi.Get("/tickets/status/:status", handler.TicketHandler.GetTicketsByStatus)
// adminApi.Patch("/tickets/:ticketID/status", handler.TicketHandler.UpdateTicketStatus)
// adminApi.Get("/tickets/business/:businessID", handler.TicketHandler.GetTicketsByBusiness)
