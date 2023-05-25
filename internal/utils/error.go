package utils

import (
	"fiber-test/internal/models"
	"github.com/gofiber/fiber/v2"
)

func HandleAPIError(c *fiber.Ctx, err error) error {
	apiErr := models.APIError{
		Message: err.Error(),
	}
	return c.Status(fiber.StatusInternalServerError).JSON(apiErr)
}
