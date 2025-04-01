package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func FormatValidationErrors(err error) map[string]string {
	errors := map[string]string{}

	if _, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range err.(validator.ValidationErrors) {
			switch fieldError.Tag() {
			case "required":
				errors[fieldError.Field()] = fmt.Sprintf("%s is required", fieldError.Field())
			case "email":
				errors[fieldError.Field()] = "Invalid email format"
			case "min":
				errors[fieldError.Field()] = fmt.Sprintf("%s must be atleast of %s characters", fieldError.Field(), fieldError.Param())
			case "max":
				errors[fieldError.Field()] = fmt.Sprintf("%s must not exceed %s characters", fieldError.Field(), fieldError.Param())
			default:
				errors[fieldError.Field()] = "Invalid value"
			}
		}
	}

	return errors
}
