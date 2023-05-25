package routes

import (
	ctrl "fiber-test/internal/controllers"
	svc "fiber-test/internal/service"
	"github.com/gofiber/fiber/v2"
)

func SetupUsersRoutes(app *fiber.App) {
	service := svc.NewService()
	controller := ctrl.NewController(
		service,
	)
	users := app.Group("/users")
	// Auth routes
	users.Get("/", controller.GetUser)
	users.Get("/:id", controller.GetUserByID)
	users.Post("/", controller.CreateUser)
	users.Delete("/:id", controller.RemoveUser)
}
