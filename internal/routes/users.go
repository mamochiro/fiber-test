package routes

import (
	usersCtrl "fiber-test/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupUsersRoutes(app *fiber.App) {
	controller := usersCtrl.NewController()
	users := app.Group("/users")
	// Auth routes
	users.Get("/", controller.GetUser)
}
