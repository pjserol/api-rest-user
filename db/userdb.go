package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/pjserol/api-rest-user/common/errors"
	"github.com/pjserol/api-rest-user/model"
)

// GetUser return a user from the DB
func (d ServiceDB) GetUser(ctx context.Context, id string) (model.User, error) {
	usr := model.User{}

	if err := d.DB.QueryRowContext(ctx, `
		SELECT id, created_at, email, first_name, last_name 
		FROM user_account 
		WHERE id = $1`, id).Scan(
		&usr.ID,
		&usr.CreatedAt,
		&usr.Email,
		&usr.FirstName,
		&usr.LastName,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, errors.NotFoundErr{Err: fmt.Errorf("userId not found")}
		}
		return model.User{}, err
	}
	return usr, nil
}

// CreateUser insert a user into the DB
func (d ServiceDB) CreateUser(ctx context.Context, input model.CreateUserInput) (model.User, error) {
	usr := model.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	if err := d.DB.QueryRowContext(ctx, `
		INSERT INTO user_account (
			email,
			first_name,
			last_name
			)
		VALUES ($1, $2, $3) RETURNING id`,
		usr.Email,
		usr.FirstName,
		usr.LastName,
	).Scan(
		&usr.ID,
	); err != nil {
		if strings.Contains(err.Error(), "user_account_email_key") {
			return model.User{}, errors.ConflictErr{Err: fmt.Errorf("user already exists with that email address")}
		}
		return model.User{}, err
	}

	return usr, nil
}

// UpdateUser update a user into the DB
func (d ServiceDB) UpdateUser(ctx context.Context, usr model.User) (model.User, error) {
	if err := d.DB.QueryRowContext(ctx, `
		UPDATE user_account SET
			email = $2,
			first_name = $3,
			last_name = $4
		WHERE id = $1 returning created_at`,
		usr.ID,
		usr.Email,
		usr.FirstName,
		usr.LastName,
	).Scan(&usr.CreatedAt); err != nil {
		return model.User{}, err
	}

	return usr, nil
}
