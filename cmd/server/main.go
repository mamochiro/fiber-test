package main

import (
	"fiber-test/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	//app.Use(recover)

	// Setup route
	routes.SetupUsersRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
