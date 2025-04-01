package routes

import (
	"go_microservice/user_service/handlers"

	"github.com/gofiber/fiber/v2"
)

type AppRoutes struct {
	handlers *handlers.IndexHandler
}

func NewRoute(handlers *handlers.IndexHandler) *AppRoutes {
	return &AppRoutes{
		handlers: handlers,
	}
}

func (r *AppRoutes) SetupAppRoutes(router *fiber.App) {
	api := router.Group("/api")

	// define v1
	v1 := api.Group("/v1")

	// define health route
	v1.Route("/health", r.setUpHealthRoute)

	// define user route
	v1.Route("/user", r.setUpUserRoute)
}
