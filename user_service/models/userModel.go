package models

import "time"

type UserModel struct {
	ID        int64  `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email_id"`
	Password  string `db:"user_password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
