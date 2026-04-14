package routes

import (
	// "hoodhire/database"
	"hoodhire/database"
	"hoodhire/internal/app"
	"hoodhire/internal/middlewares"
	"hoodhire/internal/repositories"

	// "hoodhire/internal/repositories"

	"github.com/gofiber/fiber/v3"
)

func SetupAdminRoutes(app *fiber.App, handler *app.APP) {
	roleRepo := &repositories.RoleRepo{DB: database.DB}
	// app.Get("/me/permissions", middlewares.AuthMiddleware, handler.RoleHandler.GetMyPermissions)

	adminApi := app.Group("/admin", middlewares.AuthMiddleware)
	adminApi.Get("/roles", handler.RoleHandler.GetAllRoles)
	adminApi.Get("/me", handler.AuthHandler.GetMe)
	adminApi.Patch("/me/password", handler.AuthHandler.ChangePassword)
	adminApi.Post("/permissions", middlewares.PermissionMiddleware(roleRepo, "rbac_control"), handler.RoleHandler.CreatePermission)
	adminApi.Put("/permissions/:permissionId", middlewares.PermissionMiddleware(roleRepo, "rbac_control"), handler.RoleHandler.UpdatePermission)
	adminApi.Delete("/permissions/:permissionId", middlewares.PermissionMiddleware(roleRepo, "rbac_control"), handler.RoleHandler.DeletePermission)

	// user management
	userRoutes := adminApi.Group("/users", middlewares.PermissionMiddleware(roleRepo, "user_management"))
	userRoutes.Get("/export", handler.AuthHandler.ExportUsers)
	userRoutes.Post("/", handler.AuthHandler.AdminCreateUser)
	userRoutes.Get("/", handler.AuthHandler.GetUsers)
	userRoutes.Put("/:id", handler.AuthHandler.EditUser)
	userRoutes.Get("/:id", handler.AuthHandler.GetUserByID)
	userRoutes.Delete("/:id", handler.AuthHandler.DeleteUser)
	userRoutes.Patch("/:id/block", handler.AuthHandler.BlockUser)
	userRoutes.Patch("/:id/unblock", handler.AuthHandler.UnblockUser)

	// business management
	businessRoutes := adminApi.Group("/businesses", middlewares.PermissionMiddleware(roleRepo, "business_management"))
	businessRoutes.Get("/export", handler.HirerHandler.ExportBusinesses)
	businessRoutes.Get("/", handler.HirerHandler.GetBusinesses)
	businessRoutes.Patch("/:id/block", handler.HirerHandler.BlockBusiness)
	businessRoutes.Patch("/:id/unblock", handler.HirerHandler.UnblockBusiness)
	businessRoutes.Delete("/:id", handler.HirerHandler.DeleteBusiness)
	businessRoutes.Patch("/:userID/approve", handler.HirerHandler.AprroveBusiness)
	businessRoutes.Patch("/:userID/reject", handler.HirerHandler.RejectBusiness)

	// hirer management 
	hirerRoutes := adminApi.Group("/hirers", middlewares.PermissionMiddleware(roleRepo, "business_management"))
	hirerRoutes.Put("/:userID/status", handler.HirerHandler.UpdateBusinessStatus)
	hirerRoutes.Get("/", handler.HirerHandler.GetAllHirers)

	// ticket management
	ticketRoutes := adminApi.Group("/tickets", middlewares.PermissionMiddleware(roleRepo, "ticket_management"))
	ticketRoutes.Get("/", handler.TicketHandler.GetTickets)
	ticketRoutes.Get("/business/:businessID", handler.TicketHandler.GetTicketsByBusiness)
	ticketRoutes.Patch("/:id/resolve", handler.TicketHandler.ResolveTicket)
	ticketRoutes.Patch("/:id/review", handler.TicketHandler.ReviewTicket)
	ticketRoutes.Patch("/:id/dismiss", handler.TicketHandler.DismissTicket)

	// rbac control — only admin should have this permission enabled
	roleRoutes := adminApi.Group("/roles", middlewares.PermissionMiddleware(roleRepo, "rbac_control"))
	roleRoutes.Post("/", handler.RoleHandler.CreateRole)
	roleRoutes.Get("/:roleId/permissions", handler.RoleHandler.GetRolePermissions)
	roleRoutes.Put("/:roleId/permissions", handler.RoleHandler.UpdateRolePermissions)
	roleRoutes.Put("/:roleId", handler.RoleHandler.UpdateRole)
	roleRoutes.Delete("/:roleId", handler.RoleHandler.DeleteRole)

	//web config control
	configRoutes := adminApi.Group("/config", middlewares.PermissionMiddleware(roleRepo, "web_config_control"))
	configRoutes.Get("/", handler.WebConfigHandler.GetAllConfigs)
	configRoutes.Patch("/toggle", handler.WebConfigHandler.ToggleConfig)

	//jobs managemenet
	jobRoutes := adminApi.Group("/jobs", middlewares.PermissionMiddleware(roleRepo, "jobs_management"))
	jobRoutes.Get("/export", handler.JobHandlers.ExportJobs)
	jobRoutes.Get("/", handler.JobHandlers.AdminGetAllJobs)
	jobRoutes.Delete("/:id", handler.JobHandlers.AdminDeleteJob)
	jobRoutes.Patch("/:id/status", handler.JobHandlers.AdminUpdateJobStatus)
	
	//category management
	categoryRoutes := adminApi.Group("/categories", middlewares.PermissionMiddleware(roleRepo, "jobs_management"))
	categoryRoutes.Get("/", handler.Cathandler.GetAllCategories)
	categoryRoutes.Get("/stats", handler.Cathandler.GetCategoryStats)
	categoryRoutes.Post("/", handler.Cathandler.CreateCategory)
	categoryRoutes.Put("/:id", handler.Cathandler.UpdateCategory)
	categoryRoutes.Delete("/:id", handler.Cathandler.DeleteCategory)

	//plans managemenet
	SubscriptionRoutes:=adminApi.Group("/subscriptions",middlewares.PermissionMiddleware(roleRepo,"subscription_management"))
	SubscriptionRoutes.Patch("/:id/toggle",handler.SubHandler.SetPlanActive)
	SubscriptionRoutes.Get("/",handler.SubHandler.GetPlans)
	SubscriptionRoutes.Patch("/:id",handler.SubHandler.UpdatePlan)
	SubscriptionRoutes.Delete("/:id",handler.SubHandler.DeletePlan)
	SubscriptionRoutes.Post("/",handler.SubHandler.CreatePlan)



	adminApi.Get("/permissions", middlewares.PermissionMiddleware(roleRepo, "rbac_control"), handler.RoleHandler.GetAllPermissions)
}
