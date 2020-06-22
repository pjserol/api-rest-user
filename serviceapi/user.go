package serviceapi

import (
	"context"

	"github.com/pjserol/api-rest-user/model"
)

//GetUser return a user
func (s Service) GetUser(ctx context.Context, id string) (model.User, error) {
	usr, err := s.configSrv.db.GetUser(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return usr, nil
}

// CreateUser create a user
func (s Service) CreateUser(ctx context.Context, input model.CreateUserInput) (model.User, error) {
	createdUser, err := s.configSrv.db.CreateUser(ctx, input)
	if err != nil {
		return model.User{}, err
	}

	return createdUser, nil
}

// UpdateUser update of the user
func (s Service) UpdateUser(ctx context.Context, input model.UpdateUserInput) (model.User, error) {
	usr, err := s.configSrv.db.GetUser(ctx, input.ID)
	if err != nil {
		return model.User{}, err
	}

	if input.Email != "" {
		usr.Email = input.Email
	}
	if input.FirstName != "" {
		usr.FirstName = input.FirstName
	}
	if input.LastName != "" {
		usr.LastName = input.LastName
	}

	if _, err := s.configSrv.db.UpdateUser(ctx, usr); err != nil {
		return model.User{}, err
	}

	return usr, nil
}
