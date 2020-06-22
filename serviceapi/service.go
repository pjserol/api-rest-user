package serviceapi

import (
	"context"
	"os"
	"time"

	"github.com/pjserol/api-rest-user/config"
	"github.com/pjserol/api-rest-user/db"
	"github.com/pjserol/api-rest-user/model"
)

type userDB interface {
	GetUser(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, input model.CreateUserInput) (model.User, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
}

///////////////////
// service
///////////////////

// Service struct
type Service struct {
	configSrv configSrv
	configVar configVar
	now       func() time.Time
}

type configSrv struct {
	db userDB
}

type configVar struct {
}

///////////////////
// Service
///////////////////

func newService(configSrv configSrv, configVar configVar, now func() time.Time) Service {
	return Service{
		configSrv: configSrv,
		configVar: configVar,
		now:       now,
	}
}

// InitService init all the services of the API
func InitService(ctx context.Context) (Service, error) {
	now := time.Now
	env := config.InitEnvironment()

	if os.Getenv("ENVIRONMENT") == "local" {
		now = func() time.Time {
			return time.Unix(123456789, 0)
		}
	}

	db, err := db.NewServiceDB(ctx, env.AppEnvironment, db.DB)
	if err != nil {
		return Service{}, err
	}

	confSrv := configSrv{
		db: db,
	}

	confVar := configVar{}

	return newService(confSrv, confVar, now), nil
}
