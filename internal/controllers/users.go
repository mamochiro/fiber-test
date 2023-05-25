package controllers

import (
	"fiber-test/internal/models"
	"fiber-test/internal/utils"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (r *Controller) GetUser(c *fiber.Ctx) error {
	users := r.service.GetUsers()
	return c.JSON(users)
}

func (r *Controller) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.HandleAPIError(c, err)
	}
	data, err := r.service.GetUserByID(id)
	if err != nil {
		return err
	}
	return c.JSON(data)
}

func (r *Controller) CreateUser(c *fiber.Ctx) error {
	var request models.CreateUserRequest

	log.Println(string(c.Body()))
	if err := c.BodyParser(&request); err != nil {
		return utils.HandleAPIError(c, err)
	}
	err := r.service.CreateUser(request.Name, request.Email)
	if err != nil {
		return utils.HandleAPIError(c, err)
	}
	return nil
}

func (r *Controller) RemoveUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.HandleAPIError(c, err)
	}
	if err := r.service.DeleteUser(id); err != nil {
		return utils.HandleAPIError(c, err)
	}
	return nil
}
