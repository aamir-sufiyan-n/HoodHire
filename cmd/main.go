package main
import (
	"hoodhire/config"
	"hoodhire/database"
	"hoodhire/internal/app"
	"hoodhire/internal/routes"
	"github.com/gofiber/fiber/v3/middleware/cors"


	"github.com/gofiber/fiber/v3"
)
func main() {
	config.LoadConfig()
	database.Connect()
	database.MigrateDB()
	app := app.InitApp()
	r := fiber.New(fiber.Config{})
	// 1. ADD CORS MIDDLEWARE FIRST (This must happen before SetupRoutes)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"}, // React/Vite ports
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
	}))
	// 2. ADD ROUTES SECOND
	routes.SetupRoutes(r, app)
	if err := r.Listen(":8080"); err != nil {
		panic(err)
	}
}