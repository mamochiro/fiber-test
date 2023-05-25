package controllers

import (
	"fiber-test/internal/models"
	"fiber-test/internal/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (r *Controller) GetUser(c *fiber.Ctx) error {
	return c.JSON(models.UserLists)
}

func (r *Controller) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.HandleAPIError(c, err)
	}
	user := models.User{}
	data, err := user.GetUserByID(id)
	if err != nil {
		return utils.HandleAPIError(c, err)
	}
	return c.JSON(data)
}

func (r *Controller) CreateUser(c *fiber.Ctx) error {
	var (
		request models.CreateUserRequest
		user    models.User
	)
	if err := c.BodyParser(&request); err != nil {
		return utils.HandleAPIError(c, err)
	}
	user.AddList(request.Name, request.Email)
	return nil
}

func (r *Controller) RemoveUser(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return utils.HandleAPIError(c, err)
	}
	user := models.User{}
	err = user.RemoveList(id)
	if err != nil {
		return utils.HandleAPIError(c, err)
	}
	return nil
}
