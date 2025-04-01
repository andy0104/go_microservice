package handlers

import (
	"errors"
	"go_microservice/user_service/dto"
	"go_microservice/user_service/services"
	"go_microservice/user_service/utility/responses"
	"go_microservice/user_service/utility/validation"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc services.IndexServices
}

func (uh *UserHandler) Signup(ctx *fiber.Ctx) error {
	var req dto.UserSignupRequest
	if err := ctx.BodyParser(&req); err != nil {
		return responses.ErrorJson(ctx, fiber.StatusBadRequest, err, nil)
	}

	if err := validation.Validate.Struct(req); err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		return responses.ErrorJson(ctx, fiber.StatusBadRequest, responses.ErrValidationFields, validationErrors)
	}

	msg, err := uh.svc.UserService.UserSignup(req)
	if err != nil {
		if errors.Is(err, responses.ErrEmailInUseServer) {
			return responses.ErrorJson(ctx, fiber.StatusConflict, err, nil)
		} else {
			return responses.ErrorJson(ctx, fiber.StatusInternalServerError, responses.ErrValidationFields, nil)
		}
	}

	return responses.WriteJson(ctx, fiber.StatusCreated, fiber.Map{"message": msg}, false)
}

func (uh *UserHandler) Signin(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "user signin",
	})
}
