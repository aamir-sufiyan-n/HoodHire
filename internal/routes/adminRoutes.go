package routes

import (
	"hoodhire/internal/app"
	"hoodhire/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func SetupAdminRoutes(app *fiber.App, handler *app.APP) {
	adminApi := app.Group("/admin", middlewares.AuthMiddleware, middlewares.RoleMiddleware("admin"))
	{

		adminApi.Get("/users", handler.AuthHandler.GetUsers)
		adminApi.Get("/users/:id", handler.AuthHandler.GetUserByID)
		adminApi.Delete("/users/:id", handler.AuthHandler.DeleteUser)
		adminApi.Patch("/users/:id/block", handler.AuthHandler.BlockUser)
		adminApi.Patch("/users/:id/unblock", handler.AuthHandler.UnblockUser)

		adminApi.Get("/businesses", handler.HirerHandler.GetBusinesses)
		adminApi.Patch("/businesses/:id/block", handler.HirerHandler.BlockBusiness)
		adminApi.Patch("/businesses/:id/unblock", handler.HirerHandler.UnblockBusiness)
		adminApi.Delete("/businesses/:id", handler.HirerHandler.DeleteBusiness)

		adminApi.Put("/hirer/:userID/status", handler.HirerHandler.UpdateBusinessStatus)
		adminApi.Get("/hirers", handler.HirerHandler.GetAllHirers)

		adminApi.Get("/tickets", handler.TicketHandler.GetTickets)
		adminApi.Get("/tickets/business/:businessID", handler.TicketHandler.GetTicketsByBusiness)
		adminApi.Patch("/tickets/:id/resolve", handler.TicketHandler.ResolveTicket)
		adminApi.Patch("/tickets/:id/review", handler.TicketHandler.ReviewTicket)
		adminApi.Patch("/tickets/:id/dismiss", handler.TicketHandler.DismissTicket)

		adminApi.Patch("/businesses/:userID/approve", handler.HirerHandler.AprroveBusiness)
		adminApi.Patch("/businesses/:userID/reject", handler.HirerHandler.RejectBusiness)
	}
}
