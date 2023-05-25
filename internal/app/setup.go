package app

import (
	"fiber-test/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupApp() *fiber.App {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	//app.Use(recover)

	// Setup route
	routes.SetupUsersRoutes(app)

	return app
}
