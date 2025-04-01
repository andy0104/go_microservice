package repository

import (
	"database/sql"
	"errors"
	"go_microservice/user_service/dto"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func (up *UserRepository) Create(payload dto.UserSignupRequest) (int64, error) {
	query := `INSERT INTO "Users"(first_name, last_name, email_id, user_password) VALUES($1,$2,$3,$4)`
	result, err := up.db.Exec(query, payload.FirstName, payload.LastName, payload.Email, payload.Password)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (up *UserRepository) GetUserByEmailID(email string) (*User, error) {
	var user User

	if err := up.db.Get(
		&user,
		`SELECT user_id, first_name, last_name, email_id, user_password FROM "Users" WHERE email_id = $1`,
		email,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
