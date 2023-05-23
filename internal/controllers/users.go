package controllers

import (
	"fiber-test/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (r *Controller) GetUser(c *fiber.Ctx) error {
	users := []models.User{
		{
			ID:    1,
			Name:  "mark",
			Email: "mark@email.com",
		},
	}
	return c.JSON(users)
}
