package handlers

import "github.com/gofiber/fiber/v2"

type HealthHandler struct{}

func (hh *HealthHandler) Ping(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}
