package routes

import "github.com/gofiber/fiber/v2"

func (rr *AppRoutes) setUpUserRoute(router fiber.Router) {
	router.Post("/signup", rr.handlers.User.Signup)
	router.Post("/signin", rr.handlers.User.Signin)
}
