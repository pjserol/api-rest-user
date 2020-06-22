// +build integration

package db

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/pjserol/api-rest-user/common/errors"
	"github.com/pjserol/api-rest-user/model"
)

func Test_DatabaseGetUser(t *testing.T) {
	db, err := NewServiceDB(context.Background(), "local", DB)
	if err != nil {
		t.Fatal(err)
	}

	testUsers := []model.User{
		model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test@test.com", FirstName: "John", LastName: "Doe"},
	}

	tests := []struct {
		name         string
		id           string
		expectedUser model.User
		expectedErr  error
	}{
		{
			name:         "test get user happy path",
			id:           "11111111-1111-1111-1111-111111111111",
			expectedUser: testUsers[0],
			expectedErr:  nil,
		},
		{
			name:         "test get user not found",
			id:           "11111111-1111-1111-1111-111111111112",
			expectedUser: model.User{},
			expectedErr:  errors.NotFoundErr{Err: fmt.Errorf("userId not found")},
		},
	}

	for _, test := range tests {
		cleanup(t, db)
		insertUsers(t, db, testUsers)

		usr, err := db.GetUser(context.Background(), test.id)
		if !model.UserEqual(test.expectedUser, usr) {
			t.Errorf("for %s, \nexpected user %+v, \nbut got %+v", test.name, test.expectedUser, usr)
		}

		if !reflect.DeepEqual(err, test.expectedErr) {
			t.Errorf("for %s, \nexpected error, \nbut got error %v", test.name, err)
		}
	}
}

func Test_DatabaseCreateUser(t *testing.T) {
	db, err := NewServiceDB(context.Background(), "local", DB)
	if err != nil {
		t.Fatal(err)
	}

	testUsers := []model.User{
		model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test@test.com", FirstName: "John", LastName: "Doe"},
	}

	tests := []struct {
		name         string
		userToInsert model.CreateUserInput
		expectedUser model.User
		expectedErr  error
	}{
		{
			name:         "test create user happy path",
			userToInsert: model.CreateUserInput{Email: "mike@test.com", FirstName: "Mike", LastName: "Simmons"},
			expectedUser: model.User{Email: "mike@test.com", FirstName: "Mike", LastName: "Simmons"},
			expectedErr:  nil,
		},
		{
			name:         "test create user conflict",
			userToInsert: model.CreateUserInput{Email: "test@test.com", FirstName: "John", LastName: "Doe"},
			expectedUser: model.User{},
			expectedErr:  errors.ConflictErr{Err: fmt.Errorf("user already exists with that email address")},
		},
	}

	for _, test := range tests {
		cleanup(t, db)
		insertUsers(t, db, testUsers)

		newUser, err := db.CreateUser(context.Background(), test.userToInsert)
		usr, _ := db.GetUser(context.Background(), newUser.ID)
		if !model.UserEqual(test.expectedUser, usr) {
			t.Errorf("for %s, \nexpected user %+v, \nbut got %+v", test.name, test.expectedUser, usr)
		}

		if !reflect.DeepEqual(err, test.expectedErr) {
			t.Errorf("for %s, \nexpected error, \nbut got error %v", test.name, err)
		}
	}
}

func Test_DatabaseUpdateUser(t *testing.T) {
	db, err := NewServiceDB(context.Background(), "local", DB)
	if err != nil {
		t.Fatal(err)
	}

	testUsers := []model.User{
		model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test@test.com", FirstName: "John", LastName: "Doe"},
	}

	tests := []struct {
		name         string
		userToUpdate model.User
		expectedUser model.User
		expectedErr  error
	}{
		{
			name:         "test update user happy path",
			userToUpdate: model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test2@test.com", FirstName: "John 2", LastName: "Doe 2"},
			expectedUser: model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test2@test.com", FirstName: "John 2", LastName: "Doe 2"},
			expectedErr:  nil,
		},
		{
			name:         "test update user not found",
			userToUpdate: mmodel.User{ID: "11111111-1111-1111-1111-111111111112", Email: "testnotfound@test.com"},
			expectedErr:  errors.NotFoundErr{Err: fmt.Errorf("userId not found")},
		},
	}

	for _, test := range tests {
		cleanup(t, db)
		insertUsers(t, db, testUsers)

		_, err := db.UpdateUser(context.Background(), test.userToUpdate)
		usr, _ := db.GetUser(context.Background(), test.userToUpdate.ID)

		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("for %s, \nexpected user %+v, \nbut got %+v", test.name, test.expectedUser, usr)
		}

		if !model.UserEqual(test.expectedUser, usr) {
			t.Errorf("for %s, expected user %v, but got user %v", test.name, test.expectedUser, usr)
		}

		if !reflect.DeepEqual(err, test.expectedErr) {
			t.Errorf("for %s, \nexpected error, \nbut got error %v", test.name, err)
		}

	}
}

///////////////////
// insert & cleanup
///////////////////

func insertUsers(t *testing.T, db ServiceDB, users []model.User) {
	for _, user := range users {
		if _, err := db.DB.ExecContext(context.Background(),
			"INSERT INTO user_account (id, created_at, email, first_name, last_name) values ($1, 1234567890, $2, $3, $4)",
			user.ID,
			user.Email,
			user.FirstName,
			user.LastName,
		); err != nil {
			t.Fatal(err)
		}
	}
}

func getUser(t *testing.T, db ServiceDB, id string) model.User {
	usr := model.User{}
	query := fmt.Sprintf(`SELECT id, created_at, email, first_name, last_name
		FROM user_account 
		WHERE id=$1`)

	_ = db.DB.QueryRowContext(context.Background(), query, id).Scan(
		&usr.ID,
		&usr.CreatedAt,
		&usr.Email,
		&usr.FirstName,
		&usr.LastName,
	)

	return usr
}

func cleanup(t *testing.T, db ServiceDB) {
	if _, err := db.DB.ExecContext(context.Background(), "TRUNCATE TABLE user_account CASCADE"); err != nil {
		t.Fatal(err)
	}
}
