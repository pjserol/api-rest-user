package serviceapi

import (
	"context"

	"github.com/pjserol/api-rest-user/model"
)

type mockDB struct {
	user model.User
	err  error
}

func (m mockDB) GetUser(ctx context.Context, id string) (model.User, error) {
	return m.user, m.err
}

func (m mockDB) CreateUser(ctx context.Context, user model.CreateUserInput) (model.User, error) {
	return m.user, m.err
}

func (m mockDB) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	return user, m.err
}
