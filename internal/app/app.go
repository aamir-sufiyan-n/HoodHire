package app

import (
	config "hoodhire/config"
	"hoodhire/database"
	controllers "hoodhire/internal/Controllers"
	repositories "hoodhire/internal/Repositories"
	services "hoodhire/internal/Services"
)

type APP struct { 
	AuthHandler *controllers.AuthController
	SeekerHandler *controllers.SeekerController
}

func InitApp()*APP{
	db:=database.DB
	redis:=config.InitRedis()

	authRepo:=&repositories.AuthRepo{DB: db}
	seekerRepo:=&repositories.SeekerRepo{DB: db}

	authServ:=&services.AuthServices{Repo: authRepo,Redis:redis}
	seekerServ:=&services.SeekerServices{Repo: seekerRepo}

	authHandler:=&controllers.AuthController{Serv: authServ}
	seekerHandler:=&controllers.SeekerController{Service: seekerServ}
	return &APP{
		AuthHandler:authHandler ,
		SeekerHandler: seekerHandler,
	}
}