package dto

type UserSignupRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email_id" validate:"required,email"`
	Password  string `json:"user_password" validate:"required,min=8,max=20"`
}
