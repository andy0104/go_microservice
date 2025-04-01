package config

import (
	"go_microservice/user_service/handlers"
	"go_microservice/user_service/routes"

	"github.com/gofiber/fiber/v2"
)

func InitServer(app *fiber.App, hndlrs handlers.IndexHandler) {
	// setup app routes
	rr := routes.NewRoute(&hndlrs)
	rr.SetupAppRoutes(app)
}
