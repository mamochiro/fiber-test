package controllers

import "fiber-test/internal/service"

type Controller struct {
	service service.Interface
}

func NewController(s service.Interface) *Controller {
	return &Controller{
		service: s,
	}
}
