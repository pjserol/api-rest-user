package serviceapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pjserol/api-rest-user/model"
)

func now() time.Time {
	return time.Unix(123456789, 0)
}

func Test_GetUser(t *testing.T) {

	testErr := fmt.Errorf("testerror")

	tests := []struct {
		name         string
		db           mockDB
		expecteduser model.User
		expectedErr  error
	}{
		{
			name: "test happy path",
			db: mockDB{
				user: model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test@mail.com", FirstName: "John", LastName: "Doe"},
			},
			expecteduser: model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test@mail.com", FirstName: "John", LastName: "Doe"},
			expectedErr:  nil,
		},
		{
			name: "test with error",
			db: mockDB{
				user: model.User{},
				err:  testErr,
			},
			expecteduser: model.User{},
			expectedErr:  testErr,
		},
	}

	for _, test := range tests {
		configSrv := configSrv{
			db: test.db,
		}

		s := newService(configSrv, configVar{}, now)

		usr, err := s.GetUser(context.Background(), "11111111-1111-1111-1111-111111111111")

		if !model.UserEqual(test.expecteduser, usr) {
			t.Errorf("for %s, \nexpected user %+v, \nbut got %+v", test.name, test.expecteduser, usr)
		}

		if err != test.expectedErr {
			t.Errorf("for %s, \nexpected error %+v, \nbut got error %+v", test.name, test.expectedErr, err)
		}
	}
}

func Test_UserCreate(t *testing.T) {

	testErr := fmt.Errorf("testerror")

	tests := []struct {
		name         string
		db           mockDB
		input        model.CreateUserInput
		expecteduser model.User
		expectedErr  error
	}{
		{
			name: "test user create happy path",
			db: mockDB{
				user: model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test@mail.com", FirstName: "John", LastName: "Doe"},
			},
			input:        model.CreateUserInput{Email: "test@mail.com", FirstName: "John", LastName: "Doe"},
			expecteduser: model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test@mail.com", FirstName: "John", LastName: "Doe"},
			expectedErr:  nil,
		},
		{
			name: "test user create with error",
			db: mockDB{
				user: model.User{},
				err:  testErr,
			},
			input:        model.CreateUserInput{Email: "test@mail.com", FirstName: "John", LastName: "Doe"},
			expecteduser: model.User{},
			expectedErr:  testErr,
		},
	}

	for _, test := range tests {
		configSrv := configSrv{
			db: test.db,
		}

		s := newService(configSrv, configVar{}, now)

		usr, err := s.CreateUser(context.Background(), test.input)

		if !model.UserEqual(test.expecteduser, usr) {
			t.Errorf("for %s, \nexpected user %+v, \nbut got %+v", test.name, test.expecteduser, usr)
		}

		if err != nil && test.expectedErr == nil {
			t.Errorf("for %s, \nexpected error, \nbut got error %v", test.name, err)
		}
	}
}

func Test_UserUpdate(t *testing.T) {

	testErr := fmt.Errorf("testerror")

	tests := []struct {
		name         string
		db           mockDB
		input        model.UpdateUserInput
		expecteduser model.User
		expectedErr  error
	}{
		{
			name: "test user update happy path",
			db: mockDB{
				user: model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test@test.com", FirstName: "John", LastName: "Doe"},
			},
			input:        model.UpdateUserInput{Email: "test@test.com", FirstName: "John", LastName: "Doe"},
			expecteduser: model.User{ID: "11111111-1111-1111-1111-111111111111", Email: "test@test.com", FirstName: "John", LastName: "Doe"},
			expectedErr:  nil,
		},
		{
			name: "test user update with error",
			db: mockDB{
				user: model.User{},
				err:  testErr,
			},
			input:        model.UpdateUserInput{Email: "test@test.com", FirstName: "John", LastName: "Doe"},
			expecteduser: model.User{},
			expectedErr:  testErr,
		},
	}

	for _, test := range tests {
		configSrv := configSrv{
			db: test.db,
		}

		s := newService(configSrv, configVar{}, now)

		usr, err := s.UpdateUser(context.Background(), test.input)

		if !model.UserEqual(test.expecteduser, usr) {
			t.Errorf("for %s, \nexpected user %+v, \nbut got %+v", test.name, test.expecteduser, usr)
		}

		if err != nil && test.expectedErr == nil {
			t.Errorf("for %s, \nexpected error, \nbut got error %v", test.name, err)
		}
	}
}
