package handlers

import (
	"go_microservice/user_service/services"

	"github.com/gofiber/fiber/v2"
)

type IndexHandler struct {
	Health interface {
		Ping(*fiber.Ctx) error
	}
	User interface {
		Signup(*fiber.Ctx) error
		Signin(*fiber.Ctx) error
	}
}

func NewIndexHandler(svc services.IndexServices) IndexHandler {
	return IndexHandler{
		Health: &HealthHandler{},
		User:   &UserHandler{svc: svc},
	}
}
