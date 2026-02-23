package main
import (
	"hoodhire/config"
	"hoodhire/database"
	"hoodhire/internal/app"
	"hoodhire/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.LoadConfig()
	database.Connect()
	database.MigrateDB()

	app := app.InitApp()
	r := fiber.New(fiber.Config{})

	routes.SetupRoutes(r, app)

	if err := r.Listen(":8080"); err != nil {
		panic(err)
	}
}