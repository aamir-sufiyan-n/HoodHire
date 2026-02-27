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
}

func InitApp()*APP{
	db:=database.DB
	redis:=config.InitRedis()

	authRepo:=&repositories.AuthRepo{DB: db}
	seekerRepo:=&repositories.SeekerRepo{DB: db}
	hirerRepo:=&repositories.HirerRepo{DB: db}

	authServ:=&services.AuthServices{Repo: authRepo,Redis:redis}
	seekerServ:=&services.SeekerServices{Repo: seekerRepo}
	hirerServ:=&services.HirerServices{Repo: hirerRepo}

	authHandler:=&controllers.AuthController{Serv: authServ}
	seekerHandler:=&controllers.SeekerController{Service: seekerServ}
	hirerHandler:=&controllers.HirerController{Serv: hirerServ}
	return &APP{
		AuthHandler:authHandler ,
		SeekerHandler: seekerHandler,
		HirerHandler: hirerHandler,
	}
}