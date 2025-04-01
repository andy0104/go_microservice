package responses

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Details any    `json:"details"`
}

var (
	ErrNoRecords        = errors.New("resource not found")
	ErrValidationFields = errors.New("validation error")
	ErrInvalidLogin     = errors.New("invalid login credentials")
	ErrInternalServer   = errors.New("internal server error")
	ErrEmailInUseServer = errors.New("email is already in use")
	ErrInvalidJwtToken  = errors.New("invalid token")
	ErrAuthTokenMissing = errors.New("missing authorization token")
	ErrAuthTokenInvalid = errors.New("invalid authorization token")
)

func WriteJson(c *fiber.Ctx, status int, data any, isError bool) error {
	resp := new(Response)
	if isError {
		resp.Data = nil
		resp.Error = data
	} else {
		resp.Data = data
		resp.Error = nil
	}
	return c.Status(status).JSON(resp)
}

func ErrorJson(c *fiber.Ctx, status int, err error, data any) error {
	errors := ErrorResponse{
		Message: err.Error(),
		Details: data,
	}
	return WriteJson(c, status, errors, true)
}
