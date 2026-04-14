package app

import (
	config "hoodhire/config"
	"hoodhire/database"
	controllers "hoodhire/internal/controllers"
	repositories "hoodhire/internal/repositories"
	services "hoodhire/internal/services"
)

type APP struct {
	AuthHandler   *controllers.AuthController
	SeekerHandler *controllers.SeekerController
	HirerHandler  *controllers.HirerController
	JobHandlers   *controllers.JobController
	FollowHandler *controllers.FollowController
	TicketHandler *controllers.TicketController
	BondHandler   *controllers.BondController
	RoleHandler  *controllers.RoleController
	WebConfigHandler *controllers.WebConfigController
	Cathandler *controllers.CategoryController
	SubHandler *controllers.SubscriptionController
}

func InitApp() *APP {
	db := database.DB
	redis := config.InitRedis()

	authRepo := &repositories.AuthRepo{DB: db}
	seekerRepo := &repositories.SeekerRepo{DB: db}
	hirerRepo := &repositories.HirerRepo{DB: db}
	jobRepo := &repositories.JobRepo{DB: db}
	follorepo := &repositories.FollowRepo{DB: db}
	ticketRepo := &repositories.TicketRepo{DB: db}
	bondRepo := &repositories.BondRepo{DB: db}
	RoleRepo:=&repositories.RoleRepo{DB: db}
	webRepo:=&repositories.WebRepo{DB: db}
	categoryRepo:=&repositories.CategoryRepo{DB: db}
	SubRepo:=&repositories.SubscriptionRepo{DB: db}

	authServ := &services.AuthServices{Repo: authRepo, Redis: redis}
	seekerServ := &services.SeekerServices{Repo: seekerRepo}
	hirerServ := &services.HirerServices{Repo: hirerRepo}
	jobServ := &services.JobServices{Repo: jobRepo, HirerRepo: hirerRepo, BondRepo: bondRepo}
	folloserv := &services.FollowServices{Repo: follorepo}
	ticketServ := &services.TicketServices{Repo: ticketRepo}
	// bondServ := services.NewBondServices(bondRepo, hirerRepo, jobRepo)
	bondServ := &services.BondServices{Repo: bondRepo,HirerRepo: hirerRepo,JobRepo: jobRepo}
	roleServ:= &services.RoleServices{Repo: *RoleRepo}
	webServ:=services.NewWebConfigService(webRepo)
	catServ:=services.NewCategoryService(categoryRepo)
	subServ:=services.NewSubscriptionService(SubRepo,hirerRepo)

	authHandler := &controllers.AuthController{Serv: authServ}
	seekerHandler := &controllers.SeekerController{Service: seekerServ}
	hirerHandler := &controllers.HirerController{Service: hirerServ}
	jobHandler := &controllers.JobController{Service: jobServ}
	followHandler := &controllers.FollowController{Service: folloserv}
	ticketHanler := &controllers.TicketController{Service: ticketServ}
	bondHandler := &controllers.BondController{Service: bondServ}
	roleHandler := &controllers.RoleController{Serv: *roleServ}
	webhandler:= controllers.NewWebConfigController(webServ)
	catHandler:=controllers.NewCategoryController(catServ)
	sunHandler:=controllers.NewSubscriptionController(subServ)
	return &APP{
		AuthHandler:   authHandler,
		SeekerHandler: seekerHandler,
		HirerHandler:  hirerHandler,
		JobHandlers:   jobHandler,
		FollowHandler: followHandler,
		TicketHandler: ticketHanler,
		BondHandler:   bondHandler,
		RoleHandler: roleHandler,
		WebConfigHandler: webhandler,
		Cathandler: catHandler,
		SubHandler: sunHandler,
	}
}
