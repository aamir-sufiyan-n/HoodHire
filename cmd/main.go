package main

import (
	"hoodhire/config"
	"hoodhire/database"
	"hoodhire/internal/app"
	"hoodhire/internal/repositories"
	"hoodhire/internal/routes"
	"hoodhire/utils"
	"os"

	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.LoadConfig()
	utils.InitCloudinary()
	database.Connect()
	database.MigrateDB()

	database.SeedWebConfig(database.DB)
	database.AdminSeeder(database.DB)
	database.SeedPermissions(database.DB)
	database.SeedAdminRole(database.DB)

	app := app.InitApp()
	r := fiber.New(fiber.Config{})
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:5173",
			"https://hood-hire-frontend.vercel.app/",
		},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
	}))
	routes.SetupRoutes(r, app)
	utils.StartSubscriptionCron(database.DB, &repositories.SubscriptionRepo{DB: database.DB})
	routes.SetupAdminRoutes(r, app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Listen(":" + port); err != nil {
		panic(err)
	}
}
