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

	authServ := &services.AuthServices{Repo: authRepo, Redis: redis}
	seekerServ := &services.SeekerServices{Repo: seekerRepo}
	hirerServ := &services.HirerServices{Repo: hirerRepo}
	jobServ := &services.JobServices{Repo: jobRepo, HirerRepo: hirerRepo, BondRepo: bondRepo}
	folloserv := &services.FollowServices{Repo: follorepo}
	ticketServ := &services.TicketServices{Repo: ticketRepo}
	bondServ := services.NewBondServices(bondRepo, hirerRepo, jobRepo)
	authHandler := &controllers.AuthController{Serv: authServ}
	seekerHandler := &controllers.SeekerController{Service: seekerServ}
	hirerHandler := &controllers.HirerController{Service: hirerServ}
	jobHandler := &controllers.JobController{Service: jobServ}
	followHandler := &controllers.FollowController{Service: folloserv}
	ticketHanler := &controllers.TicketController{Service: ticketServ}
	bondHandler := &controllers.BondController{Service: bondServ}
	return &APP{
		AuthHandler:   authHandler,
		SeekerHandler: seekerHandler,
		HirerHandler:  hirerHandler,
		JobHandlers:   jobHandler,
		FollowHandler: followHandler,
		TicketHandler: ticketHanler,
		BondHandler:   bondHandler,
	}
}
