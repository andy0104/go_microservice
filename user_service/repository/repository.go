package repository

import (
	"go_microservice/user_service/dto"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	UserId       int64     `db:"user_id"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	EmailId      string    `db:"email_id"`
	UserPassword string    `db:"user_password"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type Repository struct {
	User interface {
		Create(dto.UserSignupRequest) (int64, error)
		GetUserByEmailID(string) (*User, error)
	}
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		User: &UserRepository{db: db},
	}
}
