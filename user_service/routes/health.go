package routes

import "github.com/gofiber/fiber/v2"

func (r *AppRoutes) setUpHealthRoute(router fiber.Router) {
	router.Get("/", r.handlers.Health.Ping)
}
