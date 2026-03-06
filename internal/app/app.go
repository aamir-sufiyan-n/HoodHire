package app

import (
	config "hoodhire/config"
	"hoodhire/database"
	controllers "hoodhire/internal/controllers"
	repositories "hoodhire/internal/repositories"
	services "hoodhire/internal/services"
)

type APP struct { 
	AuthHandler *controllers.AuthController
	SeekerHandler *controllers.SeekerController
	HirerHandler *controllers.HirerController
	JobHandlers *controllers.JobController
}

func InitApp()*APP{
	db:=database.DB
	redis:=config.InitRedis()

	authRepo:=&repositories.AuthRepo{DB: db}
	seekerRepo:=&repositories.SeekerRepo{DB: db}
	hirerRepo:=&repositories.HirerRepo{DB: db}
	jobRepo:=&repositories.JobRepo{DB: db}

	authServ:=&services.AuthServices{Repo: authRepo,Redis:redis}
	seekerServ:=&services.SeekerServices{Repo: seekerRepo}
	hirerServ:=&services.HirerServices{Repo: hirerRepo}
	jobServ:=&services.JobServices{Repo: jobRepo,HirerRepo: hirerRepo}

	authHandler:=&controllers.AuthController{Serv: authServ}
	seekerHandler:=&controllers.SeekerController{Service: seekerServ}
	hirerHandler:=&controllers.HirerController{Service: hirerServ}
	jobHandler:=&controllers.JobController{Service: jobServ}
	return &APP{
		AuthHandler:authHandler ,
		SeekerHandler: seekerHandler,
		HirerHandler: hirerHandler,
		JobHandlers: jobHandler,
	}
}